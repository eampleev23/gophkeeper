package store

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"go.uber.org/zap"
)

func (d DBStore) InsertBankCard(ctx context.Context, inputModel models.BankCard) (outputModel models.BankCard, err error) {
	d.l.ZL.Debug("Зашли в InsertBankCard в DBStore..")
	d.l.ZL.Debug("", zap.Any("inputModel", inputModel))
	tx, err := d.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return models.BankCard{}, fmt.Errorf("Transaction fail.. %w", err)
	}
	defer tx.Rollback()
	tx.QueryRow( // нужен скан
		`INSERT INTO
    data_items (data_type, owner_id, meta_value)
	VALUES($1, $2, $3)
	RETURNING
	    id, owner_id, meta_value`,
		2, inputModel.OwnerID, inputModel.MetaValue).Scan(
		&outputModel.ID,
		&outputModel.OwnerID,
		&outputModel.MetaValue)

	// Шифруем данные карты
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
	// создаём вектор инициализации для номера карты
	nonceCardNumber, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	// создаём вектор инициализации для valid thru
	nonceValidThru, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	// создаём вектор инициализации для имени владельца карты
	nonceOwnerName, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	// создаём вектор инициализации для CVC
	nonceCVC, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// переводим в байты данные карты
	notEncryptedCardNumber := []byte(inputModel.CardNumber)
	notEncryptedValidThru := []byte(inputModel.ValidThru)
	notEncryptedOwnerName := []byte(inputModel.OwnerName)
	notEncryptedCVC := []byte(inputModel.CVC)

	// шифруем данные карты
	encryptedCardNumber := aesgcm.Seal(nil, nonceCardNumber, notEncryptedCardNumber, nil)
	encryptedValidThru := aesgcm.Seal(nil, nonceValidThru, notEncryptedValidThru, nil)
	encryptedOwnerName := aesgcm.Seal(nil, nonceOwnerName, notEncryptedOwnerName, nil)
	encryptedCVC := aesgcm.Seal(nil, nonceCVC, notEncryptedCVC, nil)

	// шифруем зашифрованные данные карты в строки для хранения в бд
	encryptedCardNumberStr := byteToString(encryptedCardNumber)
	encryptedValidThruStr := byteToString(encryptedValidThru)
	encryptedOwnerNameStr := byteToString(encryptedOwnerName)
	encryptedCVCStr := byteToString(encryptedCVC)

	// также переводим в строки вектора инициализации
	nonceCardNumberStr := byteToString(nonceCardNumber)
	nonceValidThruStr := byteToString(nonceValidThru)
	nonceOwnerNameStr := byteToString(nonceOwnerName)
	nonceCVCStr := byteToString(nonceCVC)

	tx.QueryRow( // нужен скан
		`INSERT INTO
	bank_card_items (item_id, card_number, valid_thru, owner_name, cvc, nonce_card_number, nonce_valid_thru, nonce_owner_name, nonce_cvc)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING
	   item_id, card_number, valid_thru, owner_name, cvc,
	    nonce_card_number, nonce_valid_thru, nonce_owner_name, nonce_cvc`,
		outputModel.ID, encryptedCardNumberStr, encryptedValidThruStr, encryptedOwnerNameStr, encryptedCVCStr,
		nonceCardNumberStr, nonceValidThruStr, nonceOwnerNameStr, nonceCVCStr).Scan(
		&outputModel.ID,
		&outputModel.CardNumber,
		&outputModel.ValidThru,
		&outputModel.OwnerName,
		&outputModel.CVC,
		&outputModel.NonceCardNumber,
		&outputModel.NonceValidThru,
		&outputModel.NonceOwnerName,
		&outputModel.NonceCVC,
	)
	tx.Commit()
	return outputModel, err
}
