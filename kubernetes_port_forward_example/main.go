package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

func main() {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting kubeconfig: %v\n", err)
		os.Exit(1)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating clientset: %v\n", err)
		os.Exit(1)
	}

	namespace := "hello-world"
	serviceName := "hello-world"
	servicePort := int32(80)
	localPort := int32(8000)

	// Get the service to find the pod selector
	svc, err := clientset.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting service: %v\n", err)
		os.Exit(1)
	}

	// Build label selector from service selector
	var labelSelector string
	for k, v := range svc.Spec.Selector {
		if labelSelector != "" {
			labelSelector += ","
		}
		labelSelector += fmt.Sprintf("%s=%s", k, v)
	}

	// Find a pod matching the service selector
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error listing pods: %v\n", err)
		os.Exit(1)
	}
	if len(pods.Items) == 0 {
		fmt.Fprintf(os.Stderr, "no pods found for service %s\n", serviceName)
		os.Exit(1)
	}
	podName := pods.Items[0].Name

	// Resolve the targetPort for the service port
	var targetPort int32
	for _, p := range svc.Spec.Ports {
		if p.Port == servicePort {
			if p.TargetPort.IntVal != 0 {
				targetPort = p.TargetPort.IntVal
			} else {
				// Named targetPort: look it up in the pod's container ports
				portName := p.TargetPort.StrVal
				for _, container := range pods.Items[0].Spec.Containers {
					for _, cp := range container.Ports {
						if cp.Name == portName {
							targetPort = cp.ContainerPort
						}
					}
				}
			}
			break
		}
	}
	if targetPort == 0 {
		fmt.Fprintf(os.Stderr, "could not resolve targetPort for service port %d\n", servicePort)
		os.Exit(1)
	}

	// Set up the SPDY connection for port-forwarding
	roundTripper, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating round tripper: %v\n", err)
		os.Exit(1)
	}

	serverURL, err := url.Parse(config.Host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing host URL: %v\n", err)
		os.Exit(1)
	}
	serverURL.Path = fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward", namespace, podName)

	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: roundTripper}, http.MethodPost, serverURL)

	stopChan := make(chan struct{}, 1)
	readyChan := make(chan struct{})

	fw, err := portforward.New(dialer, []string{fmt.Sprintf("%d:%d", localPort, targetPort)}, stopChan, readyChan, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating port-forwarder: %v\n", err)
		os.Exit(1)
	}

	// Stop on Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		close(stopChan)
	}()

	fmt.Printf("Forwarding %s.%s.svc %d -> %d (pod port)\n", serviceName, namespace, localPort, targetPort)
	if err := fw.ForwardPorts(); err != nil {
		fmt.Fprintf(os.Stderr, "error forwarding ports: %v\n", err)
		os.Exit(1)
	}
}
