package services

import (
	"context"
	"github.com/eampleev23/gophkeeper/internal/models"
)

type Services interface {
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
	// GetTextDataItemByID возвращает зашифрованные текстовые данные.
	GetTextDataItemByID(ctx context.Context, userID, inputID int) (textDataItemOutput models.TextDataItem, err error)
}
