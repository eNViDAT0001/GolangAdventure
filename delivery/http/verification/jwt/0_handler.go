package jwt

import (
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product"
	"github.com/eNViDAT0001/GolangAdventure/internal/store/domain/provider"
	"github.com/eNViDAT0001/GolangAdventure/internal/user/domain/user"
	"github.com/eNViDAT0001/GolangAdventure/internal/verification/domain/jwt"
)

type jwtHandler struct {
	jwtUC      jwt.UseCase
	userUC     user.UseCase
	providerUC provider.UseCase
	productUC  product.UseCase
}

func NewJwtHandler(jwtUC jwt.UseCase, userUC user.UseCase, providerUC provider.UseCase, productUC product.UseCase) jwt.HttpHandler {
	return &jwtHandler{jwtUC: jwtUC, userUC: userUC, providerUC: providerUC, productUC: productUC}
}
