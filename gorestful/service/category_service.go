package service

import (
	"context"
	"gorestful/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, r web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, r web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
