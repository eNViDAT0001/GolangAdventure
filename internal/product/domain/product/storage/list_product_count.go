package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/paging/paging_query"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/entities"
)

func (s productStorage) ListCountProduct(ctx context.Context, input io.ListProductInput) (total int64, err error) {
	var count int64
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Product{})

	paging_query.SetCountListPagingQuery(&input.Paging, entities.Product{}.TableName(), query)

	if len(input.ProductIDs) > 0 {
		query = query.Where("Product.id IN ?", input.ProductIDs)
	}
	if len(input.CategoryIDs) > 0 {
		query = query.Where("Product.category_id IN ?", input.CategoryIDs)
	}

	err = query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
