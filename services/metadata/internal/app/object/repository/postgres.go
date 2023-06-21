package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/arvanStorage/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type objectPostgresRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewObjectPostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (Repository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &objectPostgresRepository{
		db:     db,
		logger: logger,
	}, nil
}

func (r *objectPostgresRepository) Save(_ context.Context, object *Object) error {
	if err := r.db.Model(new(Object)).Save(object).Error; err != nil {
		return err
	}
	return nil
}

func (r *objectPostgresRepository) ReturnByKey(_ context.Context, key string) (*Object, error) {
	object := new(Object)
	if err := r.db.Model(new(Object)).Where("key = ?", key).First(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

func (r *objectPostgresRepository) TotalMonthlyVolume(_ context.Context, userID uuid.UUID) (int64, error) {
	var sum int64
	if err := r.db.Table("object").Select("sum(size)").Where("user_id = ?", userID).Row().Scan(&sum); err != nil {
		return 0, nil
	}
	return sum, nil
}
