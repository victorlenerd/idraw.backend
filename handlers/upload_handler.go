package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
	"idraw/services"
	"idraw/utils"
	"net/http"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)

	if noteID, ok := params["noteID"]; !ok {
		fmt.Fprint(w, "noteID is required parameter")
		w.WriteHeader(400)
		return
	} else {
		file, err := utils.GetFileFromReq("image", r)

		if err != nil {
			fmt.Fprint(w, err.Error())
			w.WriteHeader(500)
			return
		}

		ctx, cancel := context.WithTimeout(ctx, time.Second*50)
		defer cancel()

		fileName := ksuid.New().String()

		services.AddImageToNote(ctx, noteID, file, fileName)

		fmt.Fprintf(w, "Blob %v uploaded.\n", fileName)

		w.WriteHeader(200)
		return
	}

}