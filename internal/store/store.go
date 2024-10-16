package store

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/cnf"
	"github.com/eampleev23/gophkeeper/internal/mlg"
	"github.com/eampleev23/gophkeeper/internal/models"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

type Store interface {
	// DBConnClose закрывает соединение с базой данных
	DBConnClose() (err error)
	// InsertUser добавляет нового пользователя или возвращает ошибку о конфликте данных
	InsertUser(ctx context.Context, userRegReq models.UserRegReq) (userBack models.User, err error)
	// GetUserByLoginAndPassword проверяет по логину и паролю зарегистрирован ли такой пользователь и если да,
	// то возвращает модель пользователя
	GetUserByLoginAndPassword(ctx context.Context, userLoginReq models.UserLoginReq) (userBack models.User, err error)
	// InsertLoginPassword - метод для добавления пары логин-пароль.
	InsertLoginPassword(ctx context.Context, inputModel models.LoginPassword) (outputModel models.LoginPassword, err error)
	// GetDataItemsByUserID возвращает все сохраненные пароли пользователя
	GetDataItemsByUserID(ctx context.Context, userID int) (dataItems []models.DataItem, err error)
}

func NewStorage(c *cnf.Config, l *mlg.ZapLog) (Store, error) {
	s, err := NewDBStore(c, l)
	if err != nil {
		return nil, fmt.Errorf("error creating new db store: %w", err)
	}
	l.ZL.Debug("DB store created success..")
	return s, nil
}
