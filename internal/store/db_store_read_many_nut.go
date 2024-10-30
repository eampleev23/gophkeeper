package store

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"go.uber.org/zap"
)

func (d DBStore) GetDataItemsByUserID(ctx context.Context, userID int) (dataItems []models.DataItem, err error) {
	d.l.ZL.Info("DB method GetDataItemsByUserID is called ..")
	rows, err := d.dbConn.QueryContext(
		ctx,
		`SELECT 
    				id,data_type,owner_id,meta_value, created_at, updated_at
					FROM
					    data_items
					WHERE
					owner_id = $1`, userID)
	if err != nil {
		return nil, fmt.Errorf("not get data items for user by user_id %w", err)
	}
	for rows.Next() {
		var v models.DataItem
		err = rows.Scan(&v.ID, &v.DataType, &v.OwnerID, &v.MetaValue, &v.CreatedAt, &v.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf(" rows scan fail: %w", err)
		}
		d.l.ZL.Debug("got data item",
			zap.Int("id", v.ID),
			zap.Int("data type", v.DataType),
			zap.Int("owner id", v.OwnerID),
			zap.String("meta value", v.MetaValue),
			zap.Time("meta value", v.CreatedAt),
			zap.Time("meta value", v.UpdatedAt),
			zap.Int("id", v.ID),
		)
		dataItems = append(dataItems, v)
	}
	// проверяем на ошибки
	if err = rows.Err(); err != nil {

		return nil, fmt.Errorf("rows.Err in db store GetDataItemsByUserID: %w", err)
	}
	return dataItems, nil
}
