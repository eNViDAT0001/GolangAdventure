package category

import (
	"github.com/eNViDAT0001/GolangAdventure/internal/store/domain/category"
)

type categoryHandler struct {
	categoryUC category.UseCase
}

func NewCategoryHandler(categoryUC category.UseCase) category.HttpHandler {
	return &categoryHandler{categoryUC: categoryUC}
}
