package services

import (
	"context"
	"github.com/segmentio/ksuid"
	"idraw/models"
	"mime/multipart"
	"time"
)

func CreateNote(ctx context.Context) (*models.Note, error) {

	note := &models.Note{
		ID:  ksuid.New().String(),
		DateCreated: time.Now(),
	}

	err := note.Create(ctx)

	if err != nil {
		return nil, err
	}

	return note, nil
}

func AddImageToNote(ctx context.Context, noteID string, file multipart.File, fileName string) error {
	err := UploadToCloudStorage(ctx, fileName, file)
	if err != nil {
		return err
	}

	noteImage := models.NoteImage{
		ID:          "",
		NoteID:      noteID,
		FileName:    fileName,
		Time:        0,
		DateCreated: time.Time{},
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
