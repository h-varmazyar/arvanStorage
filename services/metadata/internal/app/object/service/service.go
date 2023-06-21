package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/arvanStorage/api"
	"github.com/h-varmazyar/arvanStorage/pkg/errors"
	metadataApi "github.com/h-varmazyar/arvanStorage/services/metadata/api"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/app/object/repository"
	quota "github.com/h-varmazyar/arvanStorage/services/metadata/internal/app/quota/service"
	"github.com/h-varmazyar/arvanStorage/services/metadata/internal/pkg/fileCache"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type Service struct {
	logger       *log.Logger
	db           repository.Repository
	quotaService *quota.Service
	fileCache    fileCache.Caching
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
	metadataApi.RegisterObjectServiceServer(server, s)
}

func (s *Service) List(ctx context.Context, req *metadataApi.ObjectListReq) (*metadataApi.Objects, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}

func (s *Service) Return(ctx context.Context, req *metadataApi.ObjectReturnReq) (*metadataApi.Object, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}

func (s *Service) Delete(ctx context.Context, req *metadataApi.ObjectDeleteReq) (*api.Void, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}

func (s *Service) NewUpload(ctx context.Context, req *metadataApi.ObjectNewUploadReq) (*metadataApi.ObjectNewUploadResp, error) {
	var (
		err            error
		userID         uuid.UUID
		remainingQuota *metadataApi.QuotaRemainingResp
	)
	userID, err = uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}
	remainingQuota, err = s.quotaService.Remaining(ctx, &metadataApi.QuotaRemainingReq{UserID: userID.String()})
	if err != nil {
		return nil, err
	}
	if remainingQuota.TotalRemainingVolume == 0 {
		return nil, errors.New(ctx, codes.ResourceExhausted).AddDetails("insufficient volume")
	}

	uploadID := uuid.New()
	//todo: use in cache

	return &metadataApi.ObjectNewUploadResp{UploadID: uploadID.String()}, nil
}

func (s *Service) UploadPart(ctx context.Context, req *metadataApi.ObjectUploadPartReq) (*api.Void, error) {
	var (
		err            error
		userID         uuid.UUID
		uploadID       uuid.UUID
		uploadedInfo   *fileCache.FileInfo
		remainingQuota *metadataApi.QuotaRemainingResp
	)
	userID, err = uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}

	uploadID, err = uuid.Parse(req.UploadID)
	if err != nil {
		return nil, err
	}

	remainingQuota, err = s.quotaService.Remaining(ctx, &metadataApi.QuotaRemainingReq{UserID: userID.String()})
	if err != nil {
		return nil, err
	}

	uploadedInfo, err = s.fileCache.ReturnFileInfo(uploadID)
	if err == nil {
		return nil, err
	}

	if remainingQuota.TotalRemainingVolume < int64(len(req.Body))+uploadedInfo.TotalUploadedVolume {
		return nil, errors.New(ctx, codes.ResourceExhausted).AddDetails("insufficient volume")
	}

	if err = s.fileCache.AddPart(uploadID, []byte(req.Body)); err != nil {
		return nil, err
	}

	return new(api.Void), nil
}

func (s *Service) CompleteUpload(ctx context.Context, req *metadataApi.ObjectCompleteUploadReq) (*api.Void, error) {
	var (
		err          error
		userID       uuid.UUID
		uploadID     uuid.UUID
		uploadedInfo *fileCache.FileInfo
		object       *repository.Object
	)
	userID, err = uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}

	uploadID, err = uuid.Parse(req.UploadID)
	if err != nil {
		return nil, err
	}

	uploadedInfo, err = s.fileCache.ReturnFileInfo(uploadID)
	if err == nil {
		return nil, err
	}

	object, err = s.db.ReturnByKey(ctx, uploadedInfo.Key)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			object = &repository.Object{
				Key:    uploadedInfo.Key,
				UserID: userID,
			}

		} else {
			return nil, err
		}
	}

	if checksum := s.fileCache.FileChecksum(uploadID); checksum != object.Checksum {
		object.Checksum = checksum
		object.Size = uploadedInfo.TotalUploadedVolume
		s.fileCache.AsyncPersistToStorage(uploadID)
	}

	if err = s.updateVolumeQuota(ctx, userID); err != nil {
		return nil, err
	}

	if err = s.db.Save(ctx, object); err != nil {
		return nil, err
	}

	return new(api.Void), nil
}

func (s *Service) AbortUpload(_ context.Context, req *metadataApi.ObjectAbortUploadReq) (*api.Void, error) {
	var (
		err      error
		uploadID uuid.UUID
	)

	uploadID, err = uuid.Parse(req.UploadID)
	if err != nil {
		return nil, err
	}

	err = s.fileCache.RemoveFile(uploadID)
	if err == nil {
		return nil, err
	}

	return new(api.Void), nil
}
