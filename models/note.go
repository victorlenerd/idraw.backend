package models

import (
	"cloud.google.com/go/datastore"
	"context"
	"idraw/config"
	"time"
)

type Note struct {
	ID          string    `json:"id"`
	Version     int       `json:"version"`
	DateCreated time.Time `json:"date_created"`
}

func (note *Note) CreateNote(ctx context.Context) error {
	client, err := datastore.NewClient(ctx, config.ProjectId)
	if err != nil {
		return err
	}

	key := datastore.NameKey("Note", note.ID, nil)
	_, err = client.Put(ctx, key, note)
	if err != nil {
		return err
	}

	return nil
}

func (note *Note) GetNote(ctx context.Context) error {
	client, err := datastore.NewClient(ctx, config.ProjectId)
	if err != nil {
		return err
	}

	key := datastore.NameKey("Note", note.ID, nil)
	err = client.Get(ctx, key, note)
	if err != nil {
		return err
	}

	return nil
}

func (note *Note) UpdateNote(ctx context.Context) error {
	client, err := datastore.NewClient(ctx, config.ProjectId)
	if err != nil {
		return err
	}

	key := datastore.NameKey("Note", note.ID, nil)
	_, err = client.Put(ctx, key, note)
	if err != nil {
		return err
	}

	return nil
}