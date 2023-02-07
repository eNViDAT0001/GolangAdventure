package address

import (
	address "github.com/eNViDAT0001/GolangAdventure/internal/address/domain/address"
)

type addressHandler struct {
	addressUC address.UseCase
}

func NewAddressHandler(addressUC address.UseCase) address.HttpHandler {
	return &addressHandler{
		addressUC: addressUC,
	}
}
