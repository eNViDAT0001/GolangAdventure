package entities

import (
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
)

type Order struct {
	wrap_gorm.SoftDeleteModel
	UserID             uint        `gorm:"column:user_id" json:"user_id"`
	ProviderID         uint        `gorm:"column:provider_id" json:"provider_id"`
	Name               string      `gorm:"column:name" json:"name"`
	Gender             *bool       `gorm:"column:gender" json:"gender"`
	Phone              string      `gorm:"column:phone" json:"phone"`
	Province           string      `gorm:"column:province" json:"province"`
	District           string      `gorm:"column:district" json:"district"`
	Ward               string      `gorm:"column:ward" json:"ward"`
	Street             string      `gorm:"column:street" json:"street"`
	Quantity           int         `gorm:"column:quantity" json:"quantity"`
	Total              int         `gorm:"column:total" json:"total"`
	Discount           int         `gorm:"column:discount" json:"discount"`
	Status             OrderStatus `gorm:"column:status" json:"status"`
	StatusDescriptions string      `gorm:"column:status_descriptions" json:"status_descriptions"`
}

func (Order) WithFields() []string {
	return []string{"provider_id", "user_id", "name", "gender", "status", "id"}
}
func (Order) SearchFields() []string {
	return []string{"name", "gender", "status", "phone", "province", "district", "ward", "street", "quantity"}
}
func (Order) SortFields() []string {
	return []string{"provider_id", "user_id", "name", "gender", "status", "total", "discount", "id", "province", "district", "ward"}
}
func (Order) TableName() string {
	return "Order"
}

type OrderStatus string

const (
	WaitingOrder    OrderStatus = "WAITING"
	ConfirmedOrder  OrderStatus = "CONFIRMED"
	DeliveringOrder OrderStatus = "DELIVERING"
	DeliveredOrder  OrderStatus = "DELIVERED"
	CancelOrder     OrderStatus = "CANCEL"
)
