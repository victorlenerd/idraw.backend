package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
	"idraw/services"
	"idraw/utils"
	"log"
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
			log.Fatalln(err)
			w.WriteHeader(500)
			return
		}

		ctx, cancel := context.WithTimeout(ctx, time.Second*50)
		defer cancel()

		fileName := ksuid.New().String()

		err = services.AddImageToNote(ctx, noteID, file, fileName)
		if err != nil {
			fmt.Fprint(w, err.Error())
			log.Fatalln(err)
			w.WriteHeader(500)
			return
		}

		fmt.Fprintf(w, "Blob %v uploaded.\n", fileName)

		w.WriteHeader(200)
		return
	}

}