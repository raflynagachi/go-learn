package repository

import (
	"context"
	databaselearn "database_learn"
	"database_learn/entity"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(databaselearn.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test repository comment",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(databaselearn.GetConnection())
	ctx := context.Background()
	id := 20
	result, err := commentRepository.FindById(ctx, int32(id))
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(databaselearn.GetConnection())
	ctx := context.Background()
	limit := 5
	result, err := commentRepository.FindAll(ctx, limit)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
