package database

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"path/filepath"
)

var Firestore *firestore.Client

func CreateDatabase() *firestore.Client {
	conf := &firebase.Config{
		DatabaseURL: "https://gofirebase-d9916.firebaseio.com/",
	}
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}
	Firestore, err = app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}

	return Firestore
}
