package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/order/entities"
)

func (s orderStorage) UpdateOrderStatus(ctx context.Context, orderID uint, status entities.OrderStatus) error {
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Order{}).Where("id = ?", orderID).Update("status", string(status)).Error
	if err != nil {
		return err
	}
	return nil
}
