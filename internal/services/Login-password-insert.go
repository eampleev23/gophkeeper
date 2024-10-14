package services

import (
	"context"
	"github.com/eampleev23/gophkeeper/internal/models"
)

func (s *Services) InsertLoginPassword(
	ctx context.Context,
	model models.LoginPassword) (result models.LoginPassword, err error) {
	s.l.ZL.Info("Insert login password service has started..")
	return model, nil
}
