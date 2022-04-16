package helper

import (
	"gorestful/model/domain"
	"gorestful/model/web"
)

func CategoryToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func CategoriesToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, CategoryToCategoryResponse(category))
	}
	return categoryResponses
}
