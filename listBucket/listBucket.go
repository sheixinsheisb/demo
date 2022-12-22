package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	strageapi "google.golang.org/api/storage/v1"
)

func main() {
	buckets, err := listBuckets()
	if err != nil {
		panic(err)
	}
	fmt.Println("Got buckets: ", buckets)
}

func listBuckets() ([]string, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	credentials, err := google.FindDefaultCredentials(ctx, strageapi.CloudPlatformScope)
	if err != nil {
		return nil, fmt.Errorf("failed to load default credentials: %v", err)
	}
	fmt.Println("credentials: ", string(credentials.JSON))

	var buckets []string
	it := client.Buckets(ctx, credentials.ProjectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to iterate buckets: %v", err)
		}
		buckets = append(buckets, battrs.Name)
	}
	return buckets, nil
}
