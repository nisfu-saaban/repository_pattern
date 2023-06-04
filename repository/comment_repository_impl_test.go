package repository

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	repository_pattern "github.com/nisfu-saaban/repository_pattern"
	"github.com/nisfu-saaban/repository_pattern/entity"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommanRepository(repository_pattern.GetConnectionDB())
	ctx := context.Background()

	comment := entity.Comment{
		Email:   "repositoryTest@Test.com",
		Comment: "repository test",
	}
	result, err := commentRepository.InsertComment(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
