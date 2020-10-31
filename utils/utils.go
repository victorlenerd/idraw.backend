package utils

import (
	"fmt"
	"mime/multipart"
	"net/http"
)

func GetFileFromReq(fileName string, r *http.Request) (multipart.File, error) {
	file, _, err := r.FormFile(fileName)

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)

		return nil, err
	}
	defer file.Close()


	return file, nil
}
