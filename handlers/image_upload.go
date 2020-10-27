package handlers

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/segmentio/ksuid"
	"io"
	"net/http"
	"time"
)

func ImageUploadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		panic(err)
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	bucketName := "idraw-app-images"

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	objectName := ksuid.New().String()

	wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)

	if _, err = io.Copy(wc, file); err != nil {
		fmt.Errorf("io.Copy: %v", err)
	}

	if err := wc.Close(); err != nil {
		fmt.Errorf("Writer.Close: %v", err)
	}

	fmt.Fprintf(w, "Blob %v uploaded.\n", objectName)

	w.WriteHeader(200)
}