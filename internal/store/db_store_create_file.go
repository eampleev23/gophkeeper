package store

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"go.uber.org/zap"
)

const typeFileData = 4

func (d DBStore) InsertFileDataItem(ctx context.Context, inputModel models.FileDataItem) (outputModel models.FileDataItem, err error) {
	d.l.ZL.Info("InsertFileDataItem dbstore method has started")
	d.l.ZL.Debug("", zap.Any("inputModel", inputModel))
	tx, err := d.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return models.FileDataItem{}, fmt.Errorf("Transaction fail.. %w", err)
	}
	defer tx.Rollback()
	tx.QueryRow( // нужен скан
		`INSERT INTO
    data_items (data_type, owner_id, meta_value)
	VALUES($1, $2, $3)
	RETURNING
	    id, owner_id, meta_value`,
		typeFileData, inputModel.OwnerID, inputModel.MetaValue).Scan(
		&outputModel.ID,
		&outputModel.OwnerID,
		&outputModel.MetaValue)

	tx.QueryRow( // нужен скан
		`INSERT INTO
	file_items (item_id, server_path)
	VALUES($1, $2)
	RETURNING
	   item_id, server_path`,
		outputModel.ID, inputModel.ServerPath).Scan(
		&outputModel.ID,
		&outputModel.ServerPath,
	)
	tx.Commit()
	return outputModel, err
}
