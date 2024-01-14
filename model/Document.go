package model

// Document represents the structure of the data to be stored in DB
type Document struct {
	ID       int    `json:"id"`
	Filename string `json:"filename"`
	FileType string `json:"file_type"`
	Size     int64  `json:"size"`
	Uri      string `json:"uri"`
}
