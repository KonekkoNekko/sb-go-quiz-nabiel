package controllers

import (
	"net/http"
	"sb-go-quiz-nabiel/database"
	"sb-go-quiz-nabiel/repository"
	"sb-go-quiz-nabiel/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleFindCategories(context *gin.Context) {
	responseCode := http.StatusOK
	categoryList, retrievalError := repository.RetrieveCategories(database.DbConnection)

	if retrievalError.Message != "" {
		responseCode = retrievalError.Status
		context.JSON(responseCode, gin.H{
			"detail": retrievalError.Message,
		})
		return
	}

	context.JSON(responseCode, gin.H{
		"items": categoryList,
	})
}

func HandleCreateCategory(context *gin.Context) {
	responseCode := http.StatusOK
	var newCategory structs.Category
	if bindError := context.ShouldBindJSON(&newCategory); bindError != nil {
		responseCode = http.StatusBadRequest
		context.JSON(responseCode, gin.H{
			"detail": bindError.Error(),
		})
		return
	}

	currentUser, _ := context.Get("user")
	userRecord, _ := currentUser.(structs.User)

	newCategory.CreatedBy = strconv.Itoa(userRecord.ID)
	newCategory.ModifiedBy = strconv.Itoa(userRecord.ID)

	creationError := repository.StoreCategory(database.DbConnection, newCategory)

	if creationError.Message != "" {
		responseCode = creationError.Status
		context.JSON(responseCode, gin.H{
			"detail": creationError.Message,
		})
		return
	}

	context.JSON(responseCode, gin.H{
		"status": "success",
	})
}

func HandleUpdateCategory(context *gin.Context) {
	responseCode := http.StatusOK
	categoryID, _ := strconv.Atoi(context.Param("id"))
	var updatedCategory structs.Category
	if bindError := context.ShouldBindJSON(&updatedCategory); bindError != nil {
		responseCode = http.StatusBadRequest
		context.JSON(responseCode, gin.H{
			"detail": bindError.Error(),
		})
		return
	}

	currentUser, _ := context.Get("user")
	userRecord, _ := currentUser.(structs.User)

	updatedCategory.ID = categoryID
	updatedCategory.CreatedBy = strconv.Itoa(userRecord.ID)
	updatedCategory.ModifiedBy = strconv.Itoa(userRecord.ID)

	updateError := repository.UpdateCategory(database.DbConnection, updatedCategory)

	if updateError.Message != "" {
		responseCode = updateError.Status
		context.JSON(responseCode, gin.H{
			"detail": updateError.Message,
		})
		return
	}

	context.JSON(responseCode, gin.H{
		"status": "success",
	})
}

func HandleDeleteCategory(context *gin.Context) {
	responseCode := http.StatusOK
	categoryID, _ := strconv.Atoi(context.Param("id"))
	deletionError := repository.EraseCategory(database.DbConnection, categoryID)

	if deletionError.Message != "" {
		responseCode = deletionError.Status
		context.JSON(responseCode, gin.H{
			"detail": deletionError.Message,
		})
		return
	}

	context.JSON(responseCode, gin.H{
		"status": "success",
	})
}

func HandleFindCategory(context *gin.Context) {
	responseCode := http.StatusOK
	categoryID, _ := strconv.Atoi(context.Param("id"))
	singleCategory, retrievalError := repository.RetrieveCategory(database.DbConnection, categoryID)

	if retrievalError.Message != "" {
		responseCode = retrievalError.Status
		context.JSON(responseCode, gin.H{
			"detail": retrievalError.Message,
		})
		return
	}

	context.JSON(responseCode, gin.H{
		"item": singleCategory,
	})
}

func HandleFindBooksByCategory(context *gin.Context) {
	responseCode := http.StatusOK
	categoryID, _ := strconv.Atoi(context.Param("id"))
	booksInCategory, retrievalError := repository.RetrieveBooksForCategory(database.DbConnection, categoryID)

	if retrievalError.Message != "" {
		responseCode = retrievalError.Status
		context.JSON(responseCode, gin.H{
			"detail": retrievalError.Message,
		})
		return
	}

	context.JSON(responseCode, gin.H{
		"items": booksInCategory,
	})
}