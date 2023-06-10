package database

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strings"
)

const (
	storageBucketName = "getjob-ef46d.appspot.com"
	ArchitectPhotos   = "architect-photos/"
	ClientPhotos      = "client-photos/"
)

var storageBucket *storage.BucketHandle

func InitStorage() *storage.BucketHandle {
	client, err := firebaseApp.Storage(context.Background())
	if err != nil {
		log.Fatalf("firebaseApp.Storage: %v", err)
	}

	bucket, err := client.Bucket(storageBucketName)
	if err != nil {
		log.Fatalf("firebaseApp.Storage: %v", err)
	}

	storageBucket = bucket

	return bucket
}

func CreateClientPhoto(file *multipart.FileHeader) (string, error) {
	object := storageBucket.Object(ClientPhotos + file.Filename)

	writer := object.NewWriter(context.TODO())
	defer func() {
		err := writer.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	reader, _ := file.Open()
	defer func() {
		err := reader.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	_, err := io.Copy(writer, reader)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/getjob-ef46d.appspot.com/o/client-photos%%2F%s?alt=media", file.Filename), nil
}

func DeleteClientImage(oldPhotoUrl string) error {
	objectName := strings.Split(strings.Split(oldPhotoUrl, "%2F")[1], "?alt")[0]
	obj := storageBucket.Object(ClientPhotos + objectName)

	// Удаляем файл из хранилища
	err := obj.Delete(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func CreateArchitectPhoto(file *multipart.FileHeader) (string, error) {
	object := storageBucket.Object(ArchitectPhotos + file.Filename)

	writer := object.NewWriter(context.TODO())
	defer func() {
		err := writer.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	reader, _ := file.Open()
	defer func() {
		err := reader.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	_, err := io.Copy(writer, reader)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/getjob-ef46d.appspot.com/o/architect-photos%%2F%s?alt=media", file.Filename), nil
}

func DeleteArchitectImage(oldPhotoUrl string) error {
	objectName := strings.Split(strings.Split(oldPhotoUrl, "%2F")[1], "?alt")[0]
	obj := storageBucket.Object(ArchitectPhotos + objectName)

	// Удаляем файл из хранилища
	err := obj.Delete(context.TODO())
	if err != nil {
		return err
	}

	return nil
}
