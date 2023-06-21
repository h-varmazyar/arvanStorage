package service

import (
	"context"
	"github.com/google/uuid"
	metadataApi "github.com/h-varmazyar/arvanStorage/services/metadata/api"
)

func (s *Service) updateVolumeQuota(ctx context.Context, userID uuid.UUID) error {
	totalVolume, err := s.db.TotalMonthlyVolume(ctx, userID)
	if err != nil {
		return err
	}
	if _, err = s.quotaService.UpdateUsedVolume(ctx, &metadataApi.Volume{
		UserID: userID.String(),
		Volume: totalVolume,
	}); err != nil {
		return err
	}
	return nil
}
