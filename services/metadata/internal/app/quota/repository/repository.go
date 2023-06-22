package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/pkg/db"
	gormext "github.com/h-varmazyar/gopack/gorm"
	log "github.com/sirupsen/logrus"
)

type Quota struct {
	gormext.UniversalModel
	UserID              uuid.UUID
	MaxRequestPerMinute int64
	MaxVolumeInMonth    int64
	TotalUsedVolume     int64
}

type Repository interface {
	ReturnByUserID(ctx context.Context, userID uuid.UUID) (*Quota, error)
	UpdateUsedVolume(ctx context.Context, quotaID uuid.UUID, volume int64) error
	ResetAllQuota(ctx context.Context) error
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (Repository, error) {
	if err := migration(ctx, db); err != nil {
		return nil, err
	}
	return NewQuotaPostgresRepository(ctx, logger, db.PostgresDB)
}

func migration(_ context.Context, dbInstance *db.DB) error {
	err := dbInstance.PostgresDB.AutoMigrate(new(Quota))
	if err != nil {
		return err
	}

	return nil
}
