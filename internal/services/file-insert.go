package services

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
)

func (s *DBServices) InsertFileDataItem(
	ctx context.Context,
	inputModel models.FileDataItem) (outputModel models.FileDataItem, err error) {
	s.l.ZL.Info("Insert FileDataItem service has started..")
	outputModel, err = s.s.InsertFileDataItem(ctx, inputModel)
	if err != nil {
		return outputModel, fmt.Errorf("%w", err)
	}
	return outputModel, nil
}