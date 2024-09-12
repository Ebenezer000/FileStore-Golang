package storage

import (
	"filestore-golang/api/models"
	"filestore-golang/config"
)

// SaveFileChunk saves a file chunk to the database
func SaveFileChunk(fileChunk models.FileChunk) error {
	query := `INSERT INTO file_chunks (file_id, chunk_data, created_at) VALUES ($1, $2, $3)`
	_, err := config.DB.Exec(query, fileChunk.FileID, fileChunk.ChunkData, fileChunk.CreatedAt)
	return err
}

// GetFileChunks retrieves file chunks from the database by file ID
func GetFileChunks(fileID string) ([]models.FileChunk, error) {
	chunks := []models.FileChunk{}
	query := `SELECT * FROM file_chunks WHERE file_id = $1`
	rows, err := config.DB.Query(query, fileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var chunk models.FileChunk
		err := rows.Scan(&chunk.ID, &chunk.FileID, &chunk.ChunkData, &chunk.CreatedAt)
		if err != nil {
			return nil, err
		}
		chunks = append(chunks, chunk)
	}

	return chunks, nil
}
