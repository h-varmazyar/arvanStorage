package fileCache

import (
	"context"
	"github.com/google/uuid"
)

type FileInfo struct {
	UploadID            uuid.UUID `json:"upload_id"`
	Key                 string    `json:"key"`
	TotalUploadedVolume int64     `json:"total_uploaded_volume"`
	content             []byte    `json:"-"`
}

type Caching interface {
	ReturnFileInfo(ctx context.Context, id uuid.UUID) (*FileInfo, error)
	FileChecksum(ctx context.Context, id uuid.UUID) string
	NewFile(ctx context.Context, key string) *FileInfo
	AddPart(ctx context.Context, id uuid.UUID, bytes []byte) error
	AsyncPersistToStorage(ctx context.Context, id uuid.UUID)
	RemoveFile(ctx context.Context, id uuid.UUID) error
}
