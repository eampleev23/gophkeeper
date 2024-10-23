package store

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"go.uber.org/zap"
)

const typeLoginPass = 1

func (d DBStore) InsertLoginPassword(ctx context.Context, inputModel models.LoginPassword) (outputModel models.LoginPassword, err error) {
	d.l.ZL.Debug("Зашли в InsertLoginPassword в DBStore..")
	d.l.ZL.Debug("", zap.Any("inputModel", inputModel))
	tx, err := d.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return models.LoginPassword{}, fmt.Errorf("Transaction fail.. %w", err)
	}
	defer tx.Rollback()
	tx.QueryRow( // нужен скан
		`INSERT INTO
    data_items (data_type, owner_id, meta_value)
	VALUES($1, $2, $3)
	RETURNING
	    id, owner_id, meta_value`,
		typeLoginPass, inputModel.OwnerID, inputModel.MetaValue).Scan(
		&outputModel.ID,
		&outputModel.OwnerID,
		&outputModel.MetaValue)

	// Шифруем логин и пароль
	// NewCipher создает и возвращает новый cipher.Block.
	// Ключевым аргументом должен быть ключ AES, 16, 24 или 32 байта
	// для выбора AES-128, AES-192 или AES-256.
	aesblock, err := aes.NewCipher([]byte(d.c.SecretKeyForData))
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
	fmt.Println("encryptedLogin=", encryptedLogin)
	encryptedPassword := aesgcm.Seal(nil, noncePassword, notEncryptedPassword, nil)

	// Видимо нужно парсить каждый байт в строку и эту строку сохранять в бд, как вам идея? ))
	// таким образом encryptedLoginStr будет в виде строки с разделителем например !

	// шифруем зашифрованные логин и пароль в строку
	//encryptedLoginStr := byteToString(encryptedLogin)
	encryptedPasswordStr := byteToString(encryptedPassword)
	nonceLoginStr := byteToString(nonceLogin)
	noncePasswordStr := byteToString(noncePassword)

	var loginBytes []byte

	tx.QueryRow( // нужен скан
		`INSERT INTO
	login_password_items (item_id, login, hash_password, nonce_login, nonce_password)
	VALUES($1, $2, $3, $4, $5)
	RETURNING
	   item_id, login, hash_password, nonce_login, nonce_password`,
		outputModel.ID, encryptedLogin, encryptedPasswordStr, nonceLoginStr, noncePasswordStr).Scan(
		&outputModel.ID,
		&loginBytes,
		&outputModel.Password,
		&outputModel.NonceLogin,
		&outputModel.NoncePassword)
	tx.Commit()
	fmt.Println("loginBytes=", loginBytes)
	return outputModel, err
}
