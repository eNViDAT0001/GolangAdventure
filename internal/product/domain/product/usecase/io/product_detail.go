package io

import (
	"github.com/eNViDAT0001/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Backend/internal/product/entities"
)

type ProductDetailCreateForm struct {
	Product        io.ProductCreateForm
	Media          []CreateProductMediaForm
	Specifications []io.ProductSpecificationCreateForm
	Descriptions   []io.ProductDescriptionsCreateForm
}

type ProductDetailForm struct {
	Product      entities.Product
	Media        []entities.ProductMedia
	Descriptions []entities.ProductDescriptions
	SpecTree     []io.ProductSpecificationFullDetail
}
