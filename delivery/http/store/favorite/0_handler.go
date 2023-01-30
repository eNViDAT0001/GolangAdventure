package banner

import (
	"github.com/eNViDAT0001/Backend/internal/store/domain/favorite"
)

type favoriteHandler struct {
	favoriteUC favorite.UseCase
}

func NewFavoriteHandler(favoriteUC favorite.UseCase) favorite.HttpHandler {
	return &favoriteHandler{favoriteUC: favoriteUC}
}
