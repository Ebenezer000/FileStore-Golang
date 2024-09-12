package models

import "time"

type FileChunk struct {
	ID        string    `json:"id"`
	FileID    string    `json:"file_id"`
	ChunkData []byte    `json:"chunk_data"`
	CreatedAt time.Time `json:"created_at"`
}

type FileMeta struct {
	ID        string    `json:"id"`
	FileName  string    `json:"file_name"`
	FileSize  int64     `json:"file_size"`
	CreatedAt time.Time `json:"created_at"`
}
