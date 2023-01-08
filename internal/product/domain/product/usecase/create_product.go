package usecase

import (
	"context"
	"github.com/eNViDAT0001/Backend/internal/product/domain/product/usecase/convert"
	ioUC "github.com/eNViDAT0001/Backend/internal/product/domain/product/usecase/io"
)

func (u *productUseCase) CreateProduct(ctx context.Context, productDetail ioUC.ProductDetailCreateForm) (productID uint, err error) {
	productID, err = u.productSto.CreateProduct(ctx, productDetail.Product)
	if err != nil {
		return 0, err
	}
	for i, _ := range productDetail.Media {
		productDetail.Media[i].ProductID = productID
	}
	mediaSto, err := convert.MediaCreateFormToMediaCreateInput(productDetail.Media)
	if err != nil {
		return productID, err
	}

	err = u.productSto.CreateProductMedia(ctx, mediaSto)
	if err != nil {
		return productID, err
	}

	for i, _ := range productDetail.Descriptions {
		productDetail.Descriptions[i].ProductID = productID
	}

	err = u.productSto.CreateBulkProductDescriptions(ctx, productDetail.Descriptions)

	return productID, nil
}