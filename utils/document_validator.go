package utils

import (
	"errors"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func ValidateDocument(fileHeader *multipart.FileHeader, w http.ResponseWriter, r *http.Request) (bool, error) {

	// Validate Type of file
	if !IsImageFile(fileHeader) {
		return false, errors.New("uploaded document is not an image")
	}

	// Validate Size of file
	if !IsImageSizeInLimit(fileHeader) {
		return false, errors.New("image size is greater than allowed size")
	}

	return true, nil
}

// isImageFile checks if the file is an image based on its content type
func IsImageFile(fileHeader *multipart.FileHeader) bool {
	switch fileHeader.Header.Get("Content-Type") {
	case "image/jpeg", "image/png", "image/gif", "image/webp":
		return true
	default:
		return false
	}
}

// IsImageSizeInLimit checks if the file size is within limit
func IsImageSizeInLimit(fileHeader *multipart.FileHeader) bool {

	maxImageSize, _ := strconv.ParseInt(os.Getenv("MAX_IMAGE_SIZE"), 16, 0)

	maxSize := int64(maxImageSize * 1024 * 1024)
	return fileHeader.Size <= maxSize
}

// Restrict File Size
// r.Body = http.MaxBytesReader(w, r.Body, 8*1024*1024)
// _, err := r.MultipartReader()
// if err != nil {
// 	return false, errors.New("image size is greater than allowed size")
// }

// r.Body = http.MaxBytesReader(w, r.Body, 5000000)
// err := r.ParseMultipartForm(5000000) // 8 MB limit
// if err != nil {
// 	return false, errors.New("image size is greater than allowed size")
// }

// Check file size based on ContentLength of request
// if r.ContentLength > 5*1024*1024 {
// 	return false, errors.New("image size is greater than allowed size 1")
// }
