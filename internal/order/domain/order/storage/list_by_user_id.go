package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Backend/internal/order/entities"
)

func (s *orderStorage) ListByUserID(ctx context.Context, userID uint, input paging.GetListInput) ([]entities.Order, error) {
	result := make([]entities.Order, 0)
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Order{}).Where("user_id = ?", userID)
	paging.SetPagingQuery(&input, entities.Order{}.TableName(), query)
	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
