package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/arvanStorage/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type quotaPostgresRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewQuotaPostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (Repository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &quotaPostgresRepository{
		db:     db,
		logger: logger,
	}, nil
}

func (r *quotaPostgresRepository) ReturnByUserID(_ context.Context, userID uuid.UUID) (*Quota, error) {
	quota := new(Quota)
	if err := r.db.Model(new(Quota)).Where("user_id = ?", userID).First(quota).Error; err != nil {
		return nil, err
	}
	return quota, nil
}

func (r *quotaPostgresRepository) UpdateUsedVolume(_ context.Context, quotaID uuid.UUID, volume int64) error {
	if err := r.db.Model(new(Quota)).Where("id = ?", quotaID).Update("total_used_volume", volume).Error; err != nil {
		return err
	}
	return nil
}
