package comment

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/paging"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/comment/storage/io"
)

type Storage interface {
	ListCommentByProductID(ctx context.Context, productID uint, filter paging.ParamsInput) ([]io.CommentDetail, error)
	CountListCommentByProductID(ctx context.Context, productID uint, filter paging.ParamsInput) (int64, error)
	GetCommentDetailByID(ctx context.Context, commentID uint) (io.CommentDetail, error)
	CreateComment(ctx context.Context, comment io.CreateComment) (commentID uint, err error)
	CreateCommentMedia(ctx context.Context, media []io.CreateCommentMedia) error
}
