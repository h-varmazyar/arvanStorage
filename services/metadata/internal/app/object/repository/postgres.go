package repository

import (
	"context"
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
