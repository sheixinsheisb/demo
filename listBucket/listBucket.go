package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func main() {
	var projectID string
	flag.StringVar(&projectID, "project_id", "", "GCP project ID")
	flag.Parse()

	for {
		buckets, err := listBuckets(projectID)
		if err != nil {
			panic(err)
		}
		log.Println("Got buckets: ", buckets, " at ", time.Now())
		time.Sleep(time.Minute)
	}
}

func listBuckets(projectID string) ([]string, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	var buckets []string
	it := client.Buckets(ctx, projectID)
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
