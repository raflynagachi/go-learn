package repository

import (
	"context"
	"database/sql"
	"errors"
	"gorestful/helper"
	"gorestful/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "INSERT INTO category(name) VALUES (?)"

	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "UPDATE category SET NAME = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sql := "SELECT id, name FROM category WHERE id = ?"
	result, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}
	if result.Next() {
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "SELECT id, name FROM category"
	result, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)

	var categories []domain.Category
	for result.Next() {
		category := domain.Category{}
		err = result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}
	return categories
}
