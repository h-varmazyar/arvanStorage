package worker

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/app/quota/repository"
	log "github.com/sirupsen/logrus"
	"time"
)

type ResetQuotaWorker struct {
	cron *gocron.Scheduler
	db   repository.Repository
	log  *log.Logger
}

func NewResetQuotaWorker(log *log.Logger, db repository.Repository) *ResetQuotaWorker {
	return &ResetQuotaWorker{
		cron: gocron.NewScheduler(time.Local),
		log:  log,
		db:   db,
	}
}

func (w *ResetQuotaWorker) Start() error {
	_, err := w.cron.Cron("0 0 1 * *").Do(w.do)
	if err != nil {
		return err
	}
	w.cron.StartAsync()
	return nil
}

func (w *ResetQuotaWorker) do() {
	err := w.db.ResetAllQuota(context.Background())
	if err != nil {
		w.log.WithError(err).Error("failed to reset quota")
	}
}
