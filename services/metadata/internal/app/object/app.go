package object

import (
	"context"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/app/object/repository"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/app/object/service"
	quotaService "github.com/h-varmazyar/arvanStorage/services/metadata/internal/app/quota/service"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/pkg/db"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/pkg/fileCache"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB, configs *Configs, quotaService *quotaService.Service) (*App, error) {
	dbInstance, err := repository.NewRepository(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	fileCacheImp := fileCache.NewRedisFileCache(configs.RedisAddress, configs.RedisPassword)
	return &App{
		Service: service.NewService(ctx, logger, dbInstance, quotaService, fileCacheImp),
	}, nil
}
