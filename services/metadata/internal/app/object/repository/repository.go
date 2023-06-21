package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/pkg/db"
	gormext "github.com/h-varmazyar/gopack/gorm"
	log "github.com/sirupsen/logrus"
)

type Object struct {
	gormext.UniversalModel
	UserID   uuid.UUID
	Key      string
	Size     int64
	Checksum string
}

type Repository interface {
	Save(ctx context.Context, object *Object) error
	ReturnByKey(ctx context.Context, key string) (*Object, error)
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (Repository, error) {
	if err := migration(ctx, db); err != nil {
		return nil, err
	}
	return NewObjectPostgresRepository(ctx, logger, db.PostgresDB)
}

func migration(_ context.Context, dbInstance *db.DB) error {
	err := dbInstance.PostgresDB.AutoMigrate(new(Object))
	if err != nil {
		return err
	}

	return nil
}
