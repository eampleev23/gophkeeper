package store

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"go.uber.org/zap"
)

const typeTextData = 3

func (d DBStore) InsertTextDataItem(ctx context.Context, inputModel models.TextDataItem) (outputModel models.TextDataItem, err error) {
	d.l.ZL.Info("InsertTextDataItem dbstore method has started")
	d.l.ZL.Debug("", zap.Any("inputModel", inputModel))
	tx, err := d.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return models.TextDataItem{}, fmt.Errorf("Transaction fail.. %w", err)
	}
	defer tx.Rollback()
	tx.QueryRow( // нужен скан
		`INSERT INTO
    data_items (data_type, owner_id, meta_value)
	VALUES($1, $2, $3)
	RETURNING
	    id, owner_id, meta_value`,
		typeTextData, inputModel.OwnerID, inputModel.MetaValue).Scan(
		&outputModel.ID,
		&outputModel.OwnerID,
		&outputModel.MetaValue)

	// Шифруем данные карты
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
	// создаём вектор инициализации для номера карты
	nonceTextContent, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	// переводим в байты текст
	notEncryptedTextContent := []byte(inputModel.TextContent)
	// шифруем текст
	encryptedTextContent := aesgcm.Seal(nil, nonceTextContent, notEncryptedTextContent, nil)
	// шифруем зашифрованные данные карты в строки для хранения в бд
	encryptedTextContentStr := packBytesToString(encryptedTextContent)
	// также переводим в строку вектор инициализации
	nonceTextContentStr := packBytesToString(nonceTextContent)

	tx.QueryRow( // нужен скан
		`INSERT INTO
	text_items (item_id, text_content, nonce_text_content)
	VALUES($1, $2, $3)
	RETURNING
	   item_id, text_content, nonce_text_content`,
		outputModel.ID, encryptedTextContentStr, nonceTextContentStr).Scan(
		&outputModel.ID,
		&outputModel.TextContent,
		&outputModel.NonceTextContent,
	)
	tx.Commit()
	return outputModel, err
}
