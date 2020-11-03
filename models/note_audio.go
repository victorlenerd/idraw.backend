package models

import (
	"cloud.google.com/go/datastore"
	"context"
	"idraw/config"
	"time"
)

type NoteAudio struct {
	ID          string    `json:"id"`
	FileName    string    `json:"file_name"`
	NoteID      string    `json:"note_id"`
	Version     int       `json:"version"`
	DateCreated time.Time `json:"date_created"`
}

func (noteAudio *NoteAudio) CreateNoteAudio(ctx context.Context) error {
	client, err := datastore.NewClient(ctx, config.ProjectId)
	if err != nil {
		return err
	}

	key := datastore.NameKey("NoteAudio", noteAudio.ID, nil)
	_, err = client.Put(ctx, key, &noteAudio)
	if err != nil {
		return err
	}

	return nil
}
