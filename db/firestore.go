package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var Client *firestore.Client

func InitFirestore() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	client, err := firestore.NewClient(ctx, "your-project-id", sa)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	Client = client
}
