package gcp

import (
	"net/http"
	// "google.golang.org/api/googleapi"
)

func Init() {
	// scopes := strings.Join([]string{
	// 	compute.DevstorageFullControlScope,
	// 	compute.ComputeScope,
	// }, " ")
}

func computeMain(client *http.Client, argv []string) {
	// service, err := compute.New(client)
	// if err != nil {
	// 	log.Fatalf("Unable to create Compute service: %v", err)
	// }

	// projectId := argv[0]
	// instanceName := argv[1]
	//
	// prefix := "https://www.googleapis.com/compute/v1/projects/" + projectId
	// imageURL := "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-7-wheezy-v20140606"
	// zone := "us-central1-a"

	// res, err := service.Images.List(projectId).Do()
	// log.Printf("Got compute.Images.List, err: %#v, %v", res, err)
}
