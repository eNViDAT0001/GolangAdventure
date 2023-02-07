package usecase

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/internal/store/entities"
)

func (u *categoryUseCase) GetCategoryChildrenByCategoryID(ctx context.Context, categoryID uint) ([]entities.Category, error) {
	return u.categorySto.GetCategoryChildrenTreeWithCategoryID(ctx, categoryID)
}
