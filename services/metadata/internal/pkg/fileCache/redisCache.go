package fileCache

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisFileCache struct {
	redisClient *redis.Client
}

func NewRedisFileCache(redisAddress, redisPassword string) *RedisFileCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0,
	})
	return &RedisFileCache{redisClient: rdb}
}

func (c *RedisFileCache) ReturnFileInfo(ctx context.Context, id uuid.UUID) (*FileInfo, error) {
	result, err := c.redisClient.Get(ctx, id.String()).Result()
	if err != nil {
		return nil, err
	}
	file := new(FileInfo)
	err = json.Unmarshal([]byte(result), file)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (c *RedisFileCache) FileChecksum(ctx context.Context, id uuid.UUID) string {
	result, err := c.redisClient.Get(ctx, id.String()).Result()
	if err != nil {
		return ""
	}
	file := new(FileInfo)
	err = json.Unmarshal([]byte(result), file)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", md5.Sum(file.content))
}

func (c *RedisFileCache) NewFile(ctx context.Context, key string) *FileInfo {
	uploadID := uuid.New()
	file := &FileInfo{
		UploadID: uploadID,
		Key:      key,
		content:  make([]byte, 0),
	}
	fileBytes, _ := json.Marshal(file)
	c.redisClient.Set(ctx, uploadID.String(), string(fileBytes), 0)
	return file
}

func (c *RedisFileCache) AddPart(ctx context.Context, id uuid.UUID, bytes []byte) error {
	err := c.redisClient.Exists(ctx, id.String()).Err()
	if err != nil {
		return err
	}
	//todo: add new part to previous parts - this section is not related to this task
	return nil
}

func (c *RedisFileCache) AsyncPersistToStorage(ctx context.Context, id uuid.UUID) {
	//todo: persist all file contents into persistent storage - this section is not related to this task
}

func (c *RedisFileCache) RemoveFile(ctx context.Context, id uuid.UUID) error {
	err := c.redisClient.Del(ctx, id.String()).Err()
	if err != nil {
		return err
	}
	return nil
}
