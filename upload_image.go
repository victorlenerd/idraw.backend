package main

import (
	"cloud.google.com/go/storage"
	"context"
	"net/http"
)

func UploadImageToCloudStorage(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}

}