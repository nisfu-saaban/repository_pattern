package repository

import (
	"context"

	"github.com/nisfu-saaban/repository_pattern/entity"
)

type CommentRepository interface {
	InsertComment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
