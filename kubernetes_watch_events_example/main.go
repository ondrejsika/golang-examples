package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, _ := kubeConfig.ClientConfig()
	namespace, _, _ := kubeConfig.Namespace()
	clientset, _ := kubernetes.NewForConfig(config)
	// you can use namespace="" to watch events in all namespaces
	watcher, _ := clientset.CoreV1().Events(namespace).Watch(context.TODO(), metav1.ListOptions{})
	for event := range watcher.ResultChan() {
		kubeEvent := event.Object.(*corev1.Event)
		fmt.Printf(
			"id=%s timestamp=%d namespace=%s kind=%s name=%s message=%s\n",
			kubeEvent.ObjectMeta.UID, kubeEvent.ObjectMeta.CreationTimestamp.Unix(),
			kubeEvent.ObjectMeta.Namespace, kubeEvent.InvolvedObject.Kind,
			kubeEvent.InvolvedObject.Name, kubeEvent.Message,
		)
	}
}
