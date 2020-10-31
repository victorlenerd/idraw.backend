package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"idraw/services"
	"net/http"
)

func GetNoteImages(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)

	if noteID, ok := params["noteID"]; !ok {
		fmt.Fprint(w, "noteID is required parameter")
		w.WriteHeader(400)
		return
	} else {
		urls, err := services.GetNoteImages(ctx, noteID)

		if err != nil {
			fmt.Fprint(w, err.Error())
			w.WriteHeader(500)
			return
		}

		data, _ := json.Marshal(urls)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}

}