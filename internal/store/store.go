package store

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/mlg"
	"github.com/eampleev23/gophkeeper/internal/models"
	"github.com/eampleev23/gophkeeper/internal/server_config"
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
	// InsertBankCard - метод для добавления банковской карты.
	InsertBankCard(ctx context.Context, inputModel models.BankCard) (outputModel models.BankCard, err error)
	// InsertTextDataItem - метод для добавления произвольных текстовых данных.
	InsertTextDataItem(ctx context.Context, inputModel models.TextDataItem) (outputModel models.TextDataItem, err error)
	// GetDataItemsByUserID возвращает все сохраненные пароли пользователя
	GetDataItemsByUserID(ctx context.Context, userID int) (dataItems []models.DataItem, err error)
	// GetLoginPassItemByID возвращает конкретную пару логин-пароль
	GetLoginPassItemByID(ctx context.Context, userID, inputID int) (loginPassOutput models.LoginPassword, err error)
	// GetBankCardByID возвращает данные о конкретной банковской карте
	GetBankCardByID(ctx context.Context, userID, inputID int) (bankCardOutput models.BankCard, err error)
}

func NewStorage(c *server_config.Config, l *mlg.ZapLog) (Store, error) {
	s, err := NewDBStore(c, l)
	if err != nil {
		return nil, fmt.Errorf("error creating new db store: %w", err)
	}
	l.ZL.Debug("DB store created success..")
	return s, nil
}
