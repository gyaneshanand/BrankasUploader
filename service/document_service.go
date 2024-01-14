package service

import (
	"brankasv1/model"
	"brankasv1/repository"
	"brankasv1/utils"
	"errors"
	"log"
	"net/http"
)

func ProcessDocument(w http.ResponseWriter, r *http.Request) (model.Document, error) {

	// Get th file data from the request
	file, handler, err := r.FormFile("data")
	if err != nil {
		return model.Document{}, errors.New("unable to fetch document from request")
	}
	defer file.Close()

	// Validate the Document
	_, err = utils.ValidateDocument(handler, w, r)
	if err != nil {
		return model.Document{}, err
	}

	// Create a temporary file to save the uploaded file
	tempFileName, err := utils.UploadDocument(file)
	if err != nil {
		return model.Document{}, err
	}

	// Save Data to Database using Document Repository
	newDoc := model.Document{
		Filename: handler.Filename,
		FileType: handler.Header.Get("Content-Type"),
		Size:     handler.Size,
		Uri:      tempFileName,
	}
	doc, err := repository.AddDocument(newDoc)
	if err != nil {
		return model.Document{}, err
	}

	log.Println(doc)

	return doc, nil
}
