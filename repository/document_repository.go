package repository

import (
	db "brankasv1/config/db"
	m "brankasv1/model"
	"context"
	"log/slog"
)

func AddDocument(doc m.Document) (m.Document, error) {

	// Acquiring a Connection from DB Connection Pool
	dbconn, err := db.AcquireDBConnection()
	if err != nil {
		return m.Document{}, err
	}
	defer db.ReleaseDBConnection(dbconn)

	// Inserting the document inside the DB
	slog.Info("Inserting the document meta into database :", doc)
	var insertedDoc m.Document
	err = dbconn.QueryRow(context.Background(), `
        INSERT INTO documents (filename, file_type, size, uri)
        VALUES ($1, $2, $3, $4)
        RETURNING *`,
		doc.Filename, doc.FileType, doc.Size, doc.Uri).Scan(
		&insertedDoc.ID,
		&insertedDoc.Filename,
		&insertedDoc.FileType,
		&insertedDoc.Size,
		&insertedDoc.Uri,
	)
	if err != nil {
		return m.Document{}, err
	}

	return insertedDoc, nil
}
