package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/profiles/2017-03-09/resources/mgmt/subscriptions"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
	// Use Azure CLI authentication
	authorizer, err := auth.NewAuthorizerFromCLI()
	handleError(err)

	// Get the subscription ID
	subscriptionID, err := getSubscriptionId(authorizer)
	handleError(err)

	// Create a new instance of the resources client
	resourcesClient := resources.NewClient(subscriptionID)
	resourcesClient.Authorizer = authorizer

	// List resources
	resourcesList, err := resourcesClient.ListComplete(context.Background(), "", "", nil)
	handleError(err)

	// Print the resources
	i := 0
	for resourcesList.NotDone() {
		if i > 1 {
			fmt.Println("")
		}
		r := resourcesList.Value()
		fmt.Println("Name:", *r.Name)
		fmt.Println("Type:", *r.Type)
		err = resourcesList.NextWithContext(context.Background())
		handleError(err)
		i++
	}
}

func getSubscriptionId(authorizer autorest.Authorizer) (string, error) {
	subscriptionsClient := subscriptions.NewClient()
	subscriptionsClient.Authorizer = authorizer

	subList, err := subscriptionsClient.List(context.Background())
	if err != nil {
		return "", err
	}

	for _, sub := range subList.Values() {
		if sub.State == subscriptions.Enabled {
			return *sub.SubscriptionID, nil
		}
	}

	return "", fmt.Errorf("no enabled subscriptions found in tenant")
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
