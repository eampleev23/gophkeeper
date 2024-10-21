package services

import (
	"context"
	"github.com/eampleev23/gophkeeper/internal/models"
)

func (s *Services) GetBankCardByID(ctx context.Context, userID, inputID int) (bankCardOutput models.BankCard, err error) {
	s.l.ZL.Info("GetBankCardByID service has started..")
	bankCardOutput, err = s.s.GetBankCardByID(ctx, userID, inputID)
	return bankCardOutput, err
}
