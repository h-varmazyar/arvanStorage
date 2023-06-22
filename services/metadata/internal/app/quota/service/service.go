package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/arvanStorage/api"
	"github.com/h-varmazyar/arvanStorage/pkg/errors"
	metadataApi "github.com/h-varmazyar/arvanStorage/services/metadata/api"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/app/quota/repository"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

type Service struct {
	logger *log.Logger
	db     repository.Repository
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, db repository.Repository) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		GrpcService.logger = logger
		GrpcService.db = db
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	metadataApi.RegisterQuotaServiceServer(server, s)
}

func (s *Service) Remaining(ctx context.Context, req *metadataApi.QuotaRemainingReq) (*metadataApi.QuotaRemainingResp, error) {
	var (
		err    error
		userID uuid.UUID
		quota  *repository.Quota
		resp   *metadataApi.QuotaRemainingResp
	)
	if userID, err = uuid.Parse(req.UserID); err != nil {
		return nil, err
	}

	if quota, err = s.db.ReturnByUserID(ctx, userID); err != nil {
		return nil, err
	}
	resp = &metadataApi.QuotaRemainingResp{
		TotalRemainingVolume: quota.MaxVolumeInMonth - quota.TotalUsedVolume,
	}
	return resp, nil
}

func (s *Service) RequestLimit(ctx context.Context, req *metadataApi.QuotaRequestLimitReq) (*metadataApi.QuotaRequestLimitResp, error) {
	var (
		err    error
		userID uuid.UUID
		quota  *repository.Quota
		resp   *metadataApi.QuotaRequestLimitResp
	)
	if userID, err = uuid.Parse(req.UserID); err != nil {
		return nil, err
	}

	if quota, err = s.db.ReturnByUserID(ctx, userID); err != nil {
		return nil, err
	}
	resp = &metadataApi.QuotaRequestLimitResp{
		Count:    quota.MaxRequestPerMinute,
		Duration: int64(time.Minute),
	}
	return resp, nil
}

func (s *Service) UpdateUsedVolume(ctx context.Context, req *metadataApi.Volume) (*api.Void, error) {
	var (
		err    error
		userID uuid.UUID
		quota  *repository.Quota
	)
	if userID, err = uuid.Parse(req.UserID); err != nil {
		return nil, err
	}

	if req.Volume < 0 {
		return nil, errors.New(ctx, codes.FailedPrecondition).AddDetails("invalid volume")
	}

	if quota, err = s.db.ReturnByUserID(ctx, userID); err != nil {
		return nil, err
	}

	if quota.MaxVolumeInMonth < req.Volume {
		return nil, errors.New(ctx, codes.FailedPrecondition).AddDetails("insufficient volume")
	}
	if err = s.db.UpdateUsedVolume(ctx, quota.ID, req.Volume); err != nil {
		return nil, err
	}
	return new(api.Void), nil
}
