package storage

import (
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/comment"
)

type commentStorage struct {
}

func NewCommentStorage() comment.Storage {
	return &commentStorage{}
}
