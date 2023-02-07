package convert

import (
	ioHttpHandler "github.com/eNViDAT0001/GolangAdventure/delivery/http/product/io"
	ioUC "github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product/usecase/io"
	"github.com/jinzhu/copier"
)

func UpdateDescriptionsReqToUpdateDescriptionsForm(input *ioHttpHandler.ProductDescriptionsWithFileUpdateReq) (ioUC.ProductDescriptionsUpdate, error) {
	var result ioUC.ProductDescriptionsUpdate
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}
	return result, nil
}
func CreateDescriptionsReqToCreateDescriptionsForm(input *ioHttpHandler.ProductDescriptionsWithFileCreateReq) (ioUC.ProductDescriptionsWithFileCreate, error) {
	var result ioUC.ProductDescriptionsWithFileCreate
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}

	return result, nil
}
