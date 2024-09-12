package utils

import (
	"filestore-golang/api/models"
	"filestore-golang/storage"
	"io"
	"net/http"
	"sort"
	"sync"
	"time"
)

// ProcessFileInParallel now accepts saveChunkFunc as a parameter, which will be passed in by the service.
func ProcessFileInParallel(file io.Reader, fileMeta models.FileMeta, saveChunkFunc func(models.FileChunk) error) error {
	var wg sync.WaitGroup
	chunkSize := 1 * 1024 * 1024 // 1MB

	for {
		buffer := make([]byte, chunkSize)
		n, err := file.Read(buffer)
		if n == 0 || err == io.EOF {
			break
		}

		fileChunk := models.FileChunk{
			FileID:    fileMeta.ID,
			ChunkData: buffer[:n],
			CreatedAt: time.Now(),
		}

		wg.Add(1)
		go func(chunk models.FileChunk) {
			defer wg.Done()
			err := saveChunkFunc(chunk)
			if err != nil {
				// Log or handle the error
				return
			}
		}(fileChunk)
	}

	wg.Wait()
	return nil
}

// MergeFileChunks retrieves chunks from the database, merges them, and sends them to the response writer
func MergeFileChunks(fileID string, w http.ResponseWriter) error {
	fileChunks, err := storage.GetFileChunks(fileID)
	if err != nil {
		return err
	}

	sort.Slice(fileChunks, func(i, j int) bool {
		return fileChunks[i].ID < fileChunks[j].ID
	})

	for _, chunk := range fileChunks {
		_, err := w.Write(chunk.ChunkData)
		if err != nil {
			return err
		}
	}

	return nil
}
