package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/entities"
)

func (s productStorage) CreateSpecification(ctx context.Context, input io.ProductSpecificationCreateForm) (specID uint, err error) {
	db := wrap_gorm.GetDB()
	err = db.Model(entities.ProductSpecification{}).
		Create(&input).Error
	if err != nil {
		return 0, err
	}
	return input.ID, err
}
