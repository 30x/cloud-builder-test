package main

import (
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudbuild/v1"
	"context"
)


func main() {

	ctx := context.Background()

	//Get our client
	client, err := google.DefaultClient(ctx, cloudbuild.CloudPlatformScope)
	if err != nil {
		fmt.Printf("Error getting client: %v\n", err)
	}

	//Get our service
	buildService, err := cloudbuild.New(client)
	if err != nil {
		fmt.Printf("Error getting builder: %v\n", err)
	}


	listOp := buildService.Projects.Builds.List("cloud-container-demo")
	resp, err := listOp.Do()
	if err != nil {
		fmt.Printf("Error running list operation: %v\n", err)
	}

	fmt.Printf("Got Response:\n%v", resp)

}