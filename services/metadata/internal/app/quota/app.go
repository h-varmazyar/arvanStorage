package quota

import (
	"context"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/app/quota/repository"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/app/quota/service"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/pkg/db"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB) (*App, error) {
	dbInstance, err := repository.NewRepository(ctx, logger, db)
	if err != nil {
		return nil, err
	}
	return &App{
		Service: service.NewService(ctx, logger, dbInstance),
	}, nil
}
