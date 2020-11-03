package services

import (
	"context"
	"github.com/segmentio/ksuid"
	"idraw/models"
	"mime/multipart"
	"time"
)

func CreateNote(ctx context.Context, noteID string) (*models.Note, error) {
	note := models.Note{
		ID:  noteID,
		Version: 0,
		DateCreated: time.Now().UTC(),
	}

	err := note.CreateNote(ctx)

	if err != nil {
		return nil, err
	}

	return &note, nil
}

func AddImageToNote(ctx context.Context, noteID string, file multipart.File, fileName string) error {
	err := UploadToCloudStorage(ctx, fileName, file)
	if err != nil {
		return err
	}

	note := models.Note{
		ID: noteID,
	}

	version := 1
	err = note.GetNote(ctx)

	if err != nil {
		note.Version = version
		err = note.CreateNote(ctx)
		if err != nil {
			return err
		}
	} else {
		note.Version = note.Version + 1
		err = note.UpdateNote(ctx)
		if err != nil {
			return err
		}
		version = note.Version
	}

	noteImage := models.NoteImage{
		ID:          ksuid.New().String(),
		NoteID:      noteID,
		FileName:    fileName,
		DateCreated: time.Now().UTC(),
		Version: version,
	}

	err = noteImage.CreateNoteImage(ctx)

	if err != nil {
		return err
	}

	return nil
}

func GetNoteImages(ctx context.Context, noteID string) ([]string, error) {
	noteImage := &models.NoteImage{
		NoteID:      noteID,
	}

	noteImages, err := noteImage.FindAllByNoteID(ctx)
	if err != nil {
		return nil, err
	}

	imageURLs := make([]string, 0, len(noteImages))

	for _, noteImage := range noteImages {
		url, err := GetFileURLFromCloudStorage(ctx, noteImage.FileName)

		if err != nil {
			return nil, err
		}

		imageURLs = append(imageURLs, url)
	}

	return imageURLs, nil
}
