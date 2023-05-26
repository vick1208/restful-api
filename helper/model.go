package helper

import (
	"vick1208/restful-api/model/domain"
	"vick1208/restful-api/model/web"
)

func ToCategoryResp(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
func ToCategoryResps(categories []domain.Category) []web.CategoryResponse {
	var categoryResps []web.CategoryResponse
	for _, category := range categories {
		categoryResps = append(categoryResps, ToCategoryResp(category))
	}
	return categoryResps
}
