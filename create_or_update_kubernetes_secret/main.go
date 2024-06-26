package main

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	// Create the Kubernetes client.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %s", err.Error())
	}

	secretName := "my-secret"
	secretNamespace := "default"

	// Check if the secret already exists.
	secret, err := clientset.CoreV1().Secrets(secretNamespace).Get(context.TODO(), secretName, metav1.GetOptions{})
	if err == nil && secret != nil {
		// Update secret if it exists.
		secret.StringData = map[string]string{
			"username": "updated_user",
			"password": "updated_password",
		}
		_, updateErr := clientset.CoreV1().Secrets(secretNamespace).Update(context.TODO(), secret, metav1.UpdateOptions{})
		if updateErr != nil {
			log.Fatalf("Failed to update secret: %s", updateErr.Error())
		}
		fmt.Println("Secret updated")
	} else {
		// Create secret if it does not exist.
		secret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      secretName,
				Namespace: secretNamespace,
			},
			StringData: map[string]string{
				"username": "created_user",
				"password": "created_password",
			},
		}
		_, createErr := clientset.CoreV1().Secrets(secretNamespace).Create(context.TODO(), secret, metav1.CreateOptions{})
		if createErr != nil {
			log.Fatalf("Failed to create secret: %s", createErr.Error())
		}
		fmt.Println("Secret created")
	}
}
