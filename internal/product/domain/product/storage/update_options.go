package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/entities"
)

func (s productStorage) UpdateProductOptions(ctx context.Context, optionID uint, options io.ProductOptionUpdateInput) error {
	db := wrap_gorm.GetDB()
	err := db.Model(entities.ProductOption{}).
		Where("id = ?", optionID).
		Updates(&options).Error
	if err != nil {
		return err
	}
	return nil
}
