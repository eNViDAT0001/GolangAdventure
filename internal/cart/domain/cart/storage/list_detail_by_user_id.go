package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/paging"
	"github.com/eNViDAT0001/GolangAdventure/external/paging/paging_query"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/cart/entities"
	"gorm.io/gorm"
)

func (c cartStorage) ListCartByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) ([]entities.CartDetail, error) {
	result := make([]entities.CartDetail, 0)
	db := wrap_gorm.GetDB()
	query := db.Table(entities.CartDetail{}.TableName()).Where("user_id = ?", userID)

	paging_query.SetPagingQuery(&filter, entities.CartDetail{}.TableName(), query)

	query = query.Scan(&result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	if query.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return result, err
}
