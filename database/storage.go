package database

import (
	"cloud.google.com/go/storage"
	"context"
	"log"
)

var storageBucket *storage.BucketHandle

func InitStorage() *storage.BucketHandle {
	client, err := firebaseApp.Storage(context.Background())
	if err != nil {
		log.Fatalf("firebaseApp.Storage: %v", err)
	}

	bucket, err := client.Bucket("gs://getjob-ef46d.appspot.com/")
	if err != nil {
		log.Fatalf("firebaseApp.Storage: %v", err)
	}

	storageBucket = bucket

	return bucket
}
