package utils

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
)

func UploadDocument(file multipart.File) (string, error) {

	// Create a temporary file to save the uploaded file
	tempFile, err := os.CreateTemp(os.Getenv("UPLOAD_PATH"), "upload-*.png") // This can be read from env
	if err != nil {
		return "", errors.New("unable to create temporary file")
	}
	defer tempFile.Close()

	// Write the file data to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		return "", errors.New("unable to write file")
	}

	return tempFile.Name(), nil
}
