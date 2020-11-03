package models

import (
	"cloud.google.com/go/datastore"
	"context"
	"idraw/config"
	"time"
)

type NoteImage struct {
	ID          string    `json:"id"`
	NoteID 		string 	  `json:"note_id"`
	FileName    string    `json:"file_name"`
	DateCreated time.Time `json:"date_created"`
}

func (noteImage *NoteImage) CreateNoteImage(ctx context.Context) error {
	client, err := datastore.NewClient(ctx, config.ProjectId)
	if err != nil {
		panic(err)
	}

	key := datastore.NameKey("NoteImage", noteImage.ID, nil)
	_, err = client.Put(ctx, key, noteImage)
	if err != nil {
		return err
	}

	return nil
}

func (noteImage *NoteImage) FindAllByNoteID(ctx context.Context) ([]NoteImage, error) {
	client, err := datastore.NewClient(ctx, config.ProjectId)
	if err != nil {
		return nil, err
	}

	noteImages := []NoteImage{}

	query := datastore.NewQuery("NoteImage").
		Filter("NoteID =", noteImage.NoteID)

	_, err = client.GetAll(ctx, query, noteImages)
	if err != nil {
		return nil, err
	}

	return noteImages, nil
}
