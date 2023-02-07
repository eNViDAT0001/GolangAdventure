package order_item

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/internal/order/entities"
)

type UseCase interface {
	ListByOrderID(ctx context.Context, orderID uint) ([]entities.OrderItem, error)
}
