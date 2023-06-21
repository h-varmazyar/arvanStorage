package fileCache

import (
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	redisClient *redis.Client
}

type FileInfo struct {
	UploadID            uuid.UUID
	Key                 string
	TotalUploadedVolume int64
}

type Caching interface {
	ReturnFileInfo(id uuid.UUID) (*FileInfo, error)
	FileChecksum(id uuid.UUID) string
	NewFile(id uuid.UUID)
	AddPart(id uuid.UUID, bytes []byte) error
	AsyncPersistToStorage(id uuid.UUID)
	RemoveFile(id uuid.UUID) error
}
