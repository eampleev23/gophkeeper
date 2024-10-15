package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// ErrConflict ошибка, которую используем для сигнала о нарушении целостности данных.
var ErrConflict = errors.New("data conflict")

// InsertUser занимается непосредственно запросом вставки строки в бд.
func (d DBStore) InsertUser(ctx context.Context, userRegReq models.UserRegReq) (newUser models.User, err error) {
	newUser = models.User{}
	bytesPassword := []byte(userRegReq.Password)
	cost := 14
	hashedPassword, err := bcrypt.GenerateFromPassword(bytesPassword, cost)
	if err != nil {
		return newUser, err
	}

	err = d.dbConn.QueryRow( // нужен скан
		`INSERT INTO
    users (login, password)
	VALUES($1, $2)
	RETURNING
	    id, login, password`,
		userRegReq.Login,
		hashedPassword).Scan(
		&newUser.ID,
		&newUser.Login,
		&newUser.Password)
	// Проверяем, что ошибка сигнализирует о потенциальном нарушении целостности данных
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
		err = ErrConflict
	}
	return newUser, err
}

func (d DBStore) InsertLoginPassword(ctx context.Context, inputModel models.LoginPassword) (outputModel models.LoginPassword, err error) {
	d.l.ZL.Info("Зашли в InsertLoginPassword в DBStore..")
	d.l.ZL.Info("", zap.Any("inputModel", inputModel))
	tx, err := d.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return models.LoginPassword{}, fmt.Errorf("Transaction fail.. %w", err)
	}
	defer tx.Rollback()
	tx.QueryRow( // нужен скан
		`INSERT INTO
    data_items (data_type, owner_id, meta_name, meta_value)
	VALUES($1, $2, $3, $4)
	RETURNING
	    id, owner_id, meta_name, meta_value`,
		1, inputModel.OwnerID, inputModel.MetaName, inputModel.MetaValue).Scan(
		&outputModel.ID,
		&outputModel.OwnerID,
		&outputModel.MetaName,
		&outputModel.MetaValue)
	// Хэшируем логин и пароль
	bytesLogin := []byte(inputModel.Login)
	cost := 14
	hashedLogin, err := bcrypt.GenerateFromPassword(bytesLogin, cost)
	bytesPassword := []byte(inputModel.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytesPassword, cost)

	tx.QueryRow( // нужен скан
		`INSERT INTO
    login_password_items (item_id, hash_login, hash_password)
	VALUES($1, $2, $3)
	RETURNING
	    hash_login, hash_password`,
		outputModel.ID, hashedLogin, hashedPassword).Scan(
		&outputModel.Login,
		&outputModel.Password)
	tx.Commit()
	return outputModel, err
}
