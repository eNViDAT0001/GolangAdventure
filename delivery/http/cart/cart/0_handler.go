package cart

import (
	"github.com/eNViDAT0001/GolangAdventure/internal/cart/domain/cart"
)

type cartHandler struct {
	cartUC cart.UseCase
}

func NewCartHandler(cartUC cart.UseCase) cart.HttpHandler {
	return &cartHandler{cartUC: cartUC}
}
