package models

import (
	"cloud.google.com/go/datastore"
	"context"
	"idraw/config"
	"time"
)

type Note struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	DateCreated time.Time `json:"date_created"`
}

func (note *Note) Create(ctx context.Context) error {
	client, err := datastore.NewClient(ctx, config.ProjectId)
	if err != nil {
		return err
	}

	key := datastore.NameKey("Note", note.ID, nil)
	_, err = client.Put(ctx, key, &note)
	if err != nil {
		return err
	}

	return nil
}
