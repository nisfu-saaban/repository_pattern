package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/nisfu-saaban/repository_pattern/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommanRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (r *commentRepositoryImpl) InsertComment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlQuery := "INSERT INTO comments(email,comment) VALUES(?,?)"
	result, err := r.DB.ExecContext(ctx, sqlQuery, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (r *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	sqlQuery := "SELEC id,email,comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := r.DB.QueryContext(ctx, sqlQuery, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("id " + strconv.Itoa(int(id)) + " Not Found")
	}

}

func (r *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlQuery := "SELEC id,email,comment FROM comments"
	rows, err := r.DB.QueryContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
