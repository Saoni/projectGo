package mobile

import (
	"fmt"
	"mobileService/models"
	"mobileService/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

var statusFailure = "Failure"

func GetArticlesArticleID(params operations.GetArticlesArticleIDParams) middleware.Responder {
	fmt.Println("Article Id is :" + params.ArticleID)
	var ArticleBody2 = "Art 2 Body 2"
	var ArticleImage1 = "Art 1 Image 1"
	articleOverview, err := returnArticlesForArticleId(params.ArticleID)
	if err != nil {
		return operations.NewGetArticlesArticleIDNotFound().WithPayload(
			&models.Response{
				Message: "Invalid article id",
				Status:  &statusFailure,
			})
	}
	return operations.NewGetArticlesArticleIDOK().WithPayload(
		&models.Article{
			Body:     &ArticleBody2,
			ImageURL: &ArticleImage1,
			Overview: &articleOverview,
		})

}

func GetCategories(params operations.GetCategoriesParams) middleware.Responder {
	fmt.Println("Inside GetCategoriees..")
	listCategoryResponse, err := returnListOfCategories()
	if err != nil {
		return operations.NewGetCategoriesCategoryIDNotFound().WithPayload(
			&models.Response{
				Message: "Error in model function",
				Status:  &statusFailure,
			})
	}
	return operations.NewGetCategoriesOK().WithPayload(
		&models.ListCategoriesResponse{
			Categories: listCategoryResponse,
		})
}

func GetCategoriesCategoryID(params operations.GetCategoriesCategoryIDParams) middleware.Responder {
	fmt.Println("Inside GetCategoriesCategoryID..")
	listArticleResponse, err := returnListOfArticles()
	if err != nil {
		return operations.NewGetCategoriesCategoryIDNotFound().WithPayload(
			&models.Response{
				Message: "Invalid category id",
				Status:  &statusFailure,
			})
	}
	return operations.NewGetCategoriesCategoryIDOK().WithPayload(
		&models.ListArticlesResponse{
			Articles: listArticleResponse,
		})
}

func PatchArticlesArticleID(params operations.PatchArticlesArticleIDParams) middleware.Responder {
	fmt.Println(params.ArticleID)
	fmt.Println(params.SetStarredRequest.Starred)
	response, err := setStarredForArticleId(params.ArticleID, params.SetStarredRequest.Starred)
	if err != nil {
		return operations.NewPatchArticlesArticleIDNotFound().WithPayload(
			&models.Response{
				Message: "Invalid Article Id",
				Status:  &statusFailure,
			})
	}
	return operations.NewPatchArticlesArticleIDOK().WithPayload(
		&models.Response{
			Message: response.Message,
			Status:  response.Status,
		})
}
