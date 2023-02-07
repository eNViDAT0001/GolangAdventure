package usecase

import (
	userPKG "github.com/eNViDAT0001/GolangAdventure/internal/user/domain/user"
	"github.com/eNViDAT0001/GolangAdventure/internal/verification/domain/jwt"
)

type jwtUseCase struct {
	//clientSto clientPKG.Storage
	userSto  userPKG.Storage
	tokenSto jwt.Storage
}

func NewJwtUseCase(
	//clientSto clientPKG.Storage,
	userSto userPKG.Storage,
	tokenSto jwt.Storage,

) jwt.UseCase {
	return &jwtUseCase{
		//clientSto: clientSto,
		userSto:  userSto,
		tokenSto: tokenSto,
	}
}
