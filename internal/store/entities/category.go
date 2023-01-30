package entities

import (
	"github.com/eNViDAT0001/Backend/external/wrap_gorm"
)

type Category struct {
	wrap_gorm.SoftDeleteModel
	CategoryParentID *uint  `gorm:"column:category_parent_id" json:"category_parent_id"`
	Name             string `gorm:"column:name" json:"name"`
	ImagePath        string `gorm:"column:image_path" json:"image_path"`
}

func (Category) TableName() string {
	return "Category"
}
