package services

import (
	"filestore-golang/api/models"
	"filestore-golang/config"
	"filestore-golang/utils"
	"net/http"
)

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
	// Add logic to retrieve file chunks from the database based on fileID
	// Example: db.Where("file_id = ?", fileID).Order("chunk_id").Find(&fileChunks)
}

func UploadFileService(r *http.Request) (string, error) {
	// Read file from request
	file, header, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create file metadata
	fileID := utils.GenerateID()
	fileMeta := models.FileMeta{
		ID:       fileID,
		FileName: header.Filename,
		FileSize: header.Size,
	}

	// Split and upload in parallel, passing SaveFileChunk as the function to save the chunks
	err = utils.ProcessFileInParallel(file, fileMeta, SaveFileChunk)
	if err != nil {
		return "", err
	}

	return fileID, nil
}

func GetFilesService() ([]models.FileMeta, error) {
	// Query the database for all file metadata
	files := []models.FileMeta{}
	return files, nil
}

func DownloadFileService(fileID string, w http.ResponseWriter) error {
	// Retrieve file chunks from the database and merge them
	err := utils.MergeFileChunks(fileID, w)
	return err
}
