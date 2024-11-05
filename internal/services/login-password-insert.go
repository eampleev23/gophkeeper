package services

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
)

func (s *DBServices) InsertLoginPassword(
	ctx context.Context,
	inputModel models.LoginPassword) (outputModel models.LoginPassword, err error) {
	s.l.ZL.Info("Insert login password service has started..")
	outputModel, err = s.s.InsertLoginPassword(ctx, inputModel)
	if err != nil {
		return outputModel, fmt.Errorf("%w", err)
	}
	return outputModel, nil
}
