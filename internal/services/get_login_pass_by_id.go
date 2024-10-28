package services

import (
	"context"
	"github.com/eampleev23/gophkeeper/internal/models"
)

func (s *DBServices) GetLoginPassItemByID(ctx context.Context, userID, inputID int) (loginPassOutput models.LoginPassword, err error) {
	s.l.ZL.Info("GetLoginPassItemByID is called..")
	loginPassOutput, err = s.s.GetLoginPassItemByID(ctx, userID, inputID)
	return loginPassOutput, nil
}
