package services

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
)

func (s *DBServices) GetFileItemByID(ctx context.Context, userID, inputID int) (fileItemOutput models.FileDataItem, err error) {
	s.l.ZL.Info("GetTextDataItemByID service has started..")
	fileItemOutput, err = s.s.GetFileItemByID(ctx, userID, inputID)
	if err != nil {
		return models.FileDataItem{}, fmt.Errorf("GetTextDataItemByID db method failed: %w", err)
	}
	return fileItemOutput, nil
}
