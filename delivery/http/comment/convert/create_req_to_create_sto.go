package convert

import (
	ioHttpHandler "github.com/eNViDAT0001/GolangAdventure/delivery/http/comment/io"
	ioSto "github.com/eNViDAT0001/GolangAdventure/internal/product/domain/comment/storage/io"
	"github.com/jinzhu/copier"
)

func CreateReqToCreateCommentInput(input *ioHttpHandler.CreateCommentReq) (ioSto.CreateComment, error) {
	var result ioSto.CreateComment
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}
	return result, nil
}
