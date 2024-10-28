package services

import (
	"context"
	"github.com/eampleev23/gophkeeper/internal/models"
	"go.uber.org/zap"
)

func (s *DBServices) GetDataItemsByUserID(ctx context.Context, userID int) (dataItems []models.DataItem, err error) {
	s.l.ZL.Info("Service GetDataItemsByUserID called with userId", zap.Int("user_id", userID))
	dataItems, err = s.s.GetDataItemsByUserID(ctx, userID)
	return dataItems, nil
}
