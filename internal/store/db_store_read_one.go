package store

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (d DBStore) GetBankCardByID(ctx context.Context, userID, inputID int) (outputBankCard models.BankCard, err error) {
	d.l.ZL.Info("DBStore method GetBankCardByID has called")
	row := d.dbConn.QueryRowContext(ctx,
		`SELECT item_id,
       card_number, valid_thru, owner_name, cvc, nonce_card_number,
       nonce_valid_thru, nonce_owner_name, nonce_cvc
       FROM bank_card_items WHERE item_id = $1 LIMIT 1`,
		inputID,
	)
	err = row.Scan(&outputBankCard.ID, &outputBankCard.CardNumber,
		&outputBankCard.ValidThru, &outputBankCard.OwnerName, &outputBankCard.CVC,
		&outputBankCard.NonceCardNumber,
		&outputBankCard.NonceValidThru, &outputBankCard.NonceOwnerName,
		&outputBankCard.NonceCVC) // Разбираем результат
	if err != nil {
		return outputBankCard, fmt.Errorf("faild to get login-pass couple by this id %w", err)
	}
	return outputBankCard, nil
}

func (d DBStore) GetUserByLoginAndPassword(
	ctx context.Context,
	userLoginReq models.UserLoginReq,
) (
	userBack models.User,
	err error) {
	userBack = models.User{}

	// получаем данные по логину
	row := d.dbConn.QueryRowContext(ctx,
		`SELECT id, login, password FROM users WHERE login = $1 LIMIT 1`,
		userLoginReq.Login,
	)
	err = row.Scan(&userBack.ID, &userBack.Login, &userBack.Password) // Разбираем результат
	if err != nil {
		return userBack, fmt.Errorf("faild to get user by login and password like this %w", err)
	}

	// расшифровываем пароль
	hashedPassword := []byte(userBack.Password)
	inputPasswordBytes := []byte(userLoginReq.Password)
	err = bcrypt.CompareHashAndPassword(hashedPassword, inputPasswordBytes)
	if err != nil {
		return userBack, fmt.Errorf("faild to get user by login and password like this %w", err)
	}
	return userBack, nil
}

func (d DBStore) GetLoginPassItemByID(
	ctx context.Context,
	userID, inputID int) (
	loginPassOutput models.LoginPassword,
	err error) {
	d.l.ZL.Info("GetLoginPassItemByID db method is called..")
	// получаем данные по логину
	row := d.dbConn.QueryRowContext(ctx,
		`SELECT item_id, hash_login, hash_password, nonce_login, nonce_password FROM login_password_items WHERE item_id = $1 LIMIT 1`,
		inputID,
	)
	err = row.Scan(&loginPassOutput.ID, &loginPassOutput.Login,
		&loginPassOutput.Password, &loginPassOutput.NonceLogin,
		&loginPassOutput.NoncePassword) // Разбираем результат
	if err != nil {
		return loginPassOutput, fmt.Errorf("faild to get login-pass couple by this id %w", err)
	}
	return loginPassOutput, nil
}
