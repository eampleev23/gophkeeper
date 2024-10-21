package services

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
)

func (s *Services) InsertBankCard(
	ctx context.Context,
	inputModel models.BankCard) (outputModel models.BankCard, err error) {
	s.l.ZL.Info("Insert bank card service has started..")
	outputModel, err = s.s.InsertBankCard(ctx, inputModel)
	if err != nil {
		return outputModel, fmt.Errorf("%w", err)
	}
	return outputModel, nil
}
