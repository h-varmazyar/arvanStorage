package db

import (
	"context"
	gormext "github.com/h-varmazyar/gopack/gorm"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DB struct {
	PostgresDB *gorm.DB
}

func NewDatabase(ctx context.Context, configs gormext.Configs) (*DB, error) {
	db := new(DB)
	if configs.DbType == gormext.PostgreSQL {
		postgres, err := newPostgres(ctx, configs)
		if err != nil {
			log.WithError(err).Error("failed to create new postgres")
			return nil, err
		}
		db.PostgresDB = postgres
	}
	return db, nil
}
