package services

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"idraw/config"
	"io"
	"mime/multipart"
	"strconv"
)

func UploadToCloudStorage(ctx context.Context, fileName string, file multipart.File) error {
	client, err := storage.NewClient(ctx)

	if err != nil {
		return err
	}

	wc := client.Bucket(config.ImageBucket).Object(fileName).NewWriter(ctx)

	if _, err = io.Copy(wc, file); err != nil {
		fmt.Errorf("io.Copy: %v", err)
		return  err
	}

	if err := wc.Close(); err != nil {
		fmt.Errorf("Writer.Close: %v", err)
		return err
	}

	return nil
}

func GetFileURLFromCloudStorage(ctx context.Context, fileName string) (string, error) {
	client, err := storage.NewClient(ctx)

	if err != nil {
		return "", err
	}

	reader, err := client.Bucket(config.ImageBucket).Object(fileName).NewReader(ctx)
	if err != nil {
		return "0", nil
	}

	image := make([]byte, reader.Attrs.Size)

	byteSize, err := reader.Read(image)

	return strconv.Itoa(byteSize), err
}