package main

import (
	addressHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/address"
	appAccessionHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/app/app_accession"
	mediaHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/app/app_file"
	cartHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/cart/cart"
	cartItemsHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/cart/cart_items"

	orderHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/order/order"
	orderItemsHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/order/order_items"

	commentHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/comment"
	productHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/product"
	bannerHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/store/banner"
	categoryHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/store/category"
	favoriteHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/store/favorite"
	providerHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/store/provider"

	userHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/user"
	jwtHttpHandlerPKG "github.com/eNViDAT0001/GolangAdventure/delivery/http/verification/jwt"

	userPKG "github.com/eNViDAT0001/GolangAdventure/internal/user/domain/user"
	userStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/user/domain/user/storage"
	userUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/user/domain/user/usecase"

	addressPKG "github.com/eNViDAT0001/GolangAdventure/internal/address/domain/address"
	addressStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/address/domain/address/storage"
	addressUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/address/domain/address/usecase"

	categoryPKG "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/category"
	categoryStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/category/storage"
	categoryUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/category/usecase"

	providerPKG "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/provider"
	providerStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/provider/storage"
	providerUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/provider/usecase"

	favoritePKG "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/favorite"
	favoriteStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/favorite/storage"
	favoriteUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/favorite/usecase"

	bannerPKG "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/banner"
	bannerStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/banner/storage"
	bannerUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/banner/usecase"

	cartPKG "github.com/eNViDAT0001/GolangAdventure/internal/cart/domain/cart"
	cartStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/cart/domain/cart/storage"
	cartUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/cart/domain/cart/usecase"

	cartItemsPKG "github.com/eNViDAT0001/GolangAdventure/internal/cart/domain/cart_item"
	cartItemsStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/cart/domain/cart_item/storage"
	cartItemsUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/cart/domain/cart_item/usecase"

	orderPKG "github.com/eNViDAT0001/GolangAdventure/internal/order/domain/order"
	orderStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/order/domain/order/storage"
	orderUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/order/domain/order/usecase"

	orderItemsPKG "github.com/eNViDAT0001/GolangAdventure/internal/order/domain/order_item"
	orderItemsStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/order/domain/order_item/storage"
	orderItemsUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/order/domain/order_item/usecase"

	productPKG "github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product"
	productStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product/storage"
	productUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product/usecase"

	commentPKG "github.com/eNViDAT0001/GolangAdventure/internal/product/domain/comment"
	commentStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/product/domain/comment/storage"
	commentUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/product/domain/comment/usecase"

	mediaPKG "github.com/eNViDAT0001/GolangAdventure/internal/file_storage/domain/media"
	mediaStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/file_storage/domain/media/storage"
	mediaUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/file_storage/domain/media/usecase"

	appAccessionPKG "github.com/eNViDAT0001/GolangAdventure/internal/app/domain/app_accession"
	appAccessionStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/app/domain/app_accession/storage"
	appAccessionUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/app/domain/app_accession/usecase"

	jwtPKG "github.com/eNViDAT0001/GolangAdventure/internal/verification/domain/jwt"
	jwtStoPKG "github.com/eNViDAT0001/GolangAdventure/internal/verification/domain/jwt/storage"
	jwtUCPKG "github.com/eNViDAT0001/GolangAdventure/internal/verification/domain/jwt/usecase"

	"github.com/google/wire"
)

var IteratorCollection = wire.NewSet(

	userHttpHandlerPKG.NewUserHandler,
	userUCPKG.NewUserUseCase,
	userStoPKG.NewUserStorage,

	addressHttpHandlerPKG.NewAddressHandler,
	addressUCPKG.NewAddressUseCase,
	addressStoPKG.NewAddressStorage,

	categoryHttpHandlerPKG.NewCategoryHandler,
	categoryUCPKG.NewCategoryUseCase,
	categoryStoPKG.NewCategoryStorage,

	mediaHttpHandlerPKG.NewMediaHandler,
	mediaUCPKG.NewMediaUseCase,
	mediaStoPKG.NewMediaStorage,

	productHttpHandlerPKG.NewProductHandler,
	productUCPKG.NewProductUseCase,
	productStoPKG.NewProductStorage,

	commentHttpHandlerPKG.NewCommentHandler,
	commentUCPKG.NewCommentUseCase,
	commentStoPKG.NewCommentStorage,

	providerHttpHandlerPKG.NewProviderHandler,
	providerUCPKG.NewProviderUseCase,
	providerStoPKG.NewProviderStorage,

	favoriteHttpHandlerPKG.NewFavoriteHandler,
	favoriteUCPKG.NewFavoriteUseCase,
	favoriteStoPKG.NewFavoriteStorage,

	bannerHttpHandlerPKG.NewBannerHandler,
	bannerUCPKG.NewBannerUseCase,
	bannerStoPKG.NewBannerStorage,

	cartHttpHandlerPKG.NewCartHandler,
	cartUCPKG.NewCartUseCase,
	cartStoPKG.NewCartStorage,

	cartItemsHttpHandlerPKG.NewCartItemHandler,
	cartItemsUCPKG.NewCartItemUseCase,
	cartItemsStoPKG.NewCartItemStorage,

	orderHttpHandlerPKG.NewOrderHandler,
	orderUCPKG.NewOrderUseCase,
	orderStoPKG.NewOrderStorage,

	orderItemsHttpHandlerPKG.NewOrderItemHandler,
	orderItemsUCPKG.NewOrderItemUseCase,
	orderItemsStoPKG.NewOrderItemStorage,

	appAccessionHttpHandlerPKG.NewAppAccessionHandler,
	appAccessionUCPKG.NewAppAccessionUseCase,
	appAccessionStoPKG.NewAppAccessionStorage,

	jwtHttpHandlerPKG.NewJwtHandler,
	jwtUCPKG.NewJwtUseCase,
	jwtStoPKG.NewJwtStorage,

	NewHandlerCollection,
)

type HandlerCollection struct {
	userHandler      userPKG.HttpHandler
	addressHandler   addressPKG.HttpHandler
	categoryHandler  categoryPKG.HttpHandler
	appAccessHandler appAccessionPKG.HttpHandler
	jwtHandler       jwtPKG.HttpHandler
	providerHandler  providerPKG.HttpHandler
	favoriteHandler  favoritePKG.HttpHandler
	bannerHandler    bannerPKG.HttpHandler
	productHandler   productPKG.HttpHandler
	commentHandler   commentPKG.HttpHandler
	mediaHandler     mediaPKG.HttpHandler
	cartHandler      cartPKG.HttpHandler
	cartItemHandler  cartItemsPKG.HttpHandler
	orderHandler     orderPKG.HttpHandler
	orderItemHandler orderItemsPKG.HttpHandler
}

func NewHandlerCollection(
	userHandler userPKG.HttpHandler,
	addressHandler addressPKG.HttpHandler,
	categoryHandler categoryPKG.HttpHandler,
	appAccessHandler appAccessionPKG.HttpHandler,
	jwtHandler jwtPKG.HttpHandler,
	providerHandler providerPKG.HttpHandler,
	favoriteHandler favoritePKG.HttpHandler,
	productHandler productPKG.HttpHandler,
	commentHandler commentPKG.HttpHandler,
	mediaHandler mediaPKG.HttpHandler,
	bannerHandler bannerPKG.HttpHandler,
	cartHandler cartPKG.HttpHandler,
	cartItemHandler cartItemsPKG.HttpHandler,
	orderHandler orderPKG.HttpHandler,
	orderItemHandler orderItemsPKG.HttpHandler,
) *HandlerCollection {
	return &HandlerCollection{
		userHandler:      userHandler,
		appAccessHandler: appAccessHandler,
		categoryHandler:  categoryHandler,
		jwtHandler:       jwtHandler,
		addressHandler:   addressHandler,
		providerHandler:  providerHandler,
		productHandler:   productHandler,
		commentHandler:   commentHandler,
		mediaHandler:     mediaHandler,
		bannerHandler:    bannerHandler,
		cartHandler:      cartHandler,
		cartItemHandler:  cartItemHandler,
		orderHandler:     orderHandler,
		orderItemHandler: orderItemHandler,
		favoriteHandler:  favoriteHandler,
	}
}
