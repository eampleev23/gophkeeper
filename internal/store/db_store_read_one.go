package store

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"golang.org/x/crypto/bcrypt"
)

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
		`SELECT item_id, hash_login, hash_password, nonce_login FROM login_password_items WHERE item_id = $1 LIMIT 1`,
		inputID,
	)
	err = row.Scan(&loginPassOutput.ID, &loginPassOutput.Login, &loginPassOutput.Password, &loginPassOutput.NonceLogin) // Разбираем результат
	if err != nil {
		return loginPassOutput, fmt.Errorf("faild to get login-pass couple by this id %w", err)
	}
	return loginPassOutput, nil
}
