package mobile

import (
	"errors"
	"fmt"
	"mobileService/models"
)

var catId1 = "Cat Id 1"
var catImage1 = "Cat 1 Image url"
var catTitle1 = "Cat 1 Title"
var catId2 = "Cat Id 2"
var catImage2 = "Cat 2 Image url"
var catTitle2 = "Cat 2 Title"
var ArticleId1 = "Art Id 1"
var ArticleSource1 = "Art 1 Source 1"
var ArticleStarred1 = false
var ArtcileTimestamp1 = "Art ts 1"
var ArticleTitle1 = "Art title 1"
var ArticleId2 = "Art Id 2"
var ArticleSource2 = "Art 2 Source 2"
var ArticleStarred2 = false
var ArtcileTimestamp2 = "Art ts 2"
var ArticleTitle2 = "Art title 2"

var returnListOfCategories = func() ([]*models.Category, error) {
	categoryList := make([]*models.Category, 0)
	category1 := &models.Category{
		ID:       &catId1,
		ImageURL: &catImage1,
		Title:    &catTitle1,
	}
	category2 := &models.Category{
		ID:       &catId2,
		ImageURL: &catImage2,
		Title:    &catTitle2,
	}
	categoryList = append(categoryList, category1)
	categoryList = append(categoryList, category2)

	return categoryList, nil
}

var returnListOfArticles = func() ([]*models.ArticleOverview, error) {
	ArticleList := make([]*models.ArticleOverview, 0)
	ArticleOverview1 := &models.ArticleOverview{
		ID:        &ArticleId1,
		Source:    &ArticleSource1,
		Starred:   &ArticleStarred1,
		Timestamp: &ArtcileTimestamp1,
		Title:     &ArticleTitle1,
	}
	ArticleOverview2 := &models.ArticleOverview{
		ID:        &ArticleId2,
		Source:    &ArticleSource2,
		Starred:   &ArticleStarred2,
		Timestamp: &ArtcileTimestamp2,
		Title:     &ArticleTitle2,
	}
	ArticleList = append(ArticleList, ArticleOverview1)
	ArticleList = append(ArticleList, ArticleOverview2)

	return ArticleList, nil
}
var returnArticlesForArticleId = func(articleId string) (models.ArticleOverview, error) {
	var err error
	var ObjectNotfoundError = errors.New(`Object Id is not present. Please enter a valid Id`)
	ArticleOverview1 := &models.ArticleOverview{
		ID:        &ArticleId1,
		Source:    &ArticleSource1,
		Starred:   &ArticleStarred1,
		Timestamp: &ArtcileTimestamp1,
		Title:     &ArticleTitle1,
	}
	ArticleOverview2 := &models.ArticleOverview{
		ID:        &ArticleId2,
		Source:    &ArticleSource2,
		Starred:   &ArticleStarred2,
		Timestamp: &ArtcileTimestamp2,
		Title:     &ArticleTitle2,
	}
	ArticleOverViewMap := map[string]*models.ArticleOverview{
		"Art_Id_1": ArticleOverview1,
		"Art_Id_2": ArticleOverview2,
	}
	var statusSuccessFlag int
	var finalArticleOverview models.ArticleOverview
	for key, _ := range ArticleOverViewMap {
		if key == articleId {
			fmt.Println("Values matched")
			finalArticleOverview = *ArticleOverViewMap[key]
			statusSuccessFlag++
		}
	}
	if statusSuccessFlag == 0 {
		err = ObjectNotfoundError
	}
	return finalArticleOverview, err
}

var setStarredForArticleId = func(articleId string, isStarred bool) (models.Response, error) {
	/*	ArticleOverview1 := &models.ArticleOverview{
				ID:        &ArticleId1,
				Source:    &ArticleSource1,
				Starred:   &ArticleStarred1,
				Timestamp: &ArtcileTimestamp1,
				Title:     &ArticleTitle1,
			}
			ArticleOverview2 := &models.ArticleOverview{
				ID:        &ArticleId2,
				Source:    &ArticleSource2,
				Starred:   &ArticleStarred2,
				Timestamp: &ArtcileTimestamp2,
				Title:     &ArticleTitle2,
			}
		ArticleOverViewMap := map[string]*models.ArticleOverview{
				"Art_Id_1": ArticleOverview1,
				"Art_Id_2": ArticleOverview2,
			}*/
	var response models.Response
	var statusSuccess = "Success"
	var statusFailure = "Failure"

	/*	for key, _ := range ArticleOverViewMap {
		if key == articleId {
			fmt.Println("Values matched")
			ArticleOverViewMap[key].Starred = &isStarred
			response.Message = "Value Updated OK"
			response.Status = &statusSuccess
			statusSuccessFlag++
		}
	}*/
	articleOverview, err := returnArticlesForArticleId(articleId)
	if err != nil {
		response.Message = "Invalid Object Id"
		response.Status = &statusFailure
	}
	articleOverview.Starred = &isStarred
	response.Message = "Value Updated OK"
	response.Status = &statusSuccess

	return response, err
}
