package database

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"path/filepath"
)

var firebaseApp *firebase.App

var Firestore *firestore.Client

var conf = &firebase.Config{
	DatabaseURL:   "https://gofirebase-d9916.firebaseio.com/",
	StorageBucket: "gs://getjob-ef46d.appspot.com/",
}

func InitFirebaseApp() {
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

	firebaseApp = app
}

func InitFirestore() *firestore.Client {
	var err error
	Firestore, err = firebaseApp.Firestore(context.Background())
	if err != nil {
		log.Fatalf("firebaseApp.Firestore: %v", err)
	}

	return Firestore
}
