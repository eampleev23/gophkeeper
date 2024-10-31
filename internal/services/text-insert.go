package services

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
)

func (s *DBServices) InsertTextDataItem(
	ctx context.Context,
	inputModel models.TextDataItem) (
	outputModel models.TextDataItem,
	err error) {
	s.l.ZL.Info("Insert Text Data Item service has started..")
	outputModel, err = s.s.InsertTextDataItem(ctx, inputModel)
	if err != nil {
		return outputModel, fmt.Errorf("%w", err)
	}
	return outputModel, nil
}
