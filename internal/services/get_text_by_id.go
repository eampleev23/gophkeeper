package services

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
)

func (s *Services) GetTextDataItemByID(ctx context.Context, userID, inputID int) (textDataItemOutput models.TextDataItem, err error) {
	s.l.ZL.Info("GetTextDataItemByID service has started..")
	textDataItemOutput, err = s.s.GetTextDataItemByID(ctx, userID, inputID)
	if err != nil {
		return models.TextDataItem{}, fmt.Errorf("GetTextDataItemByID db method failed: %w", err)
	}
	return textDataItemOutput, nil
}
