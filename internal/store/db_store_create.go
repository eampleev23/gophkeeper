package store

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"strconv"
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

	// Шифруем логин и пароль
	key := []byte("TuUdlQmYyD1DTaiGVV31ipyWnbKa0jUD")
	// NewCipher создает и возвращает новый cipher.Block.
	// Ключевым аргументом должен быть ключ AES, 16, 24 или 32 байта
	// для выбора AES-128, AES-192 или AES-256.
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	// NewGCM возвращает заданный 128-битный блочный шифр
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	// создаём вектор инициализации для логина
	nonceLogin, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	// создаём вектор инициализации для пароля
	noncePassword, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// переводим в байты логин и пароль
	notEncryptedLogin := []byte(inputModel.Login)
	notEncryptedPassword := []byte(inputModel.Password)

	// шифруем логин и пароль
	encryptedLogin := aesgcm.Seal(nil, nonceLogin, notEncryptedLogin, nil)
	encryptedPassword := aesgcm.Seal(nil, noncePassword, notEncryptedPassword, nil)

	// Видимо нужно парсить каждый байт в строку и эту строку сохранять в бд, как вам идея? ))
	// таким образом encryptedLoginStr будет в виде строки с разделителем например !

	// шифруем зашифрованные логин и пароль в строку
	encryptedLoginStr := byteToString(encryptedLogin)
	encryptedPasswordStr := byteToString(encryptedPassword)
	nonceLoginStr := byteToString(nonceLogin)
	noncePasswordStr := byteToString(noncePassword)

	tx.QueryRow( // нужен скан
		`INSERT INTO
	login_password_items (item_id, hash_login, hash_password, nonce_login, nonce_password)
	VALUES($1, $2, $3, $4, $5)
	RETURNING
	   item_id, hash_login, hash_password, nonce_login, nonce_password`,
		outputModel.ID, encryptedLoginStr, encryptedPasswordStr, nonceLoginStr, noncePasswordStr).Scan(
		&outputModel.ID,
		&outputModel.Login,
		&outputModel.Password,
		&outputModel.NonceLogin,
		&outputModel.NoncePassword)
	tx.Commit()
	return outputModel, err
}

func generateRandom(size int) ([]byte, error) {
	// генерируем криптостойкие случайные байты в b
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func byteToString(b []byte) string {
	var result string
	for i := 0; i < len(b); i++ {
		intVal := int(b[i])
		strVal := strconv.Itoa(intVal)
		result += strVal + "!"
	}
	return result
}
