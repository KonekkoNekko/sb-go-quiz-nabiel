package controllers

import (
	"net/http"
	"sb-go-quiz-nabiel/database"
	"sb-go-quiz-nabiel/repository"
	"sb-go-quiz-nabiel/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleFindBooks(context *gin.Context) {
	responseCode := http.StatusOK
	bookList, retrievalError := repository.RetrieveBooks(database.DbConnection)

	if retrievalError.Message != "" {
		responseCode = retrievalError.Status
		context.JSON(responseCode, gin.H{
			"detail": retrievalError.Message,
		})
		return
	}

	context.JSON(responseCode, gin.H{
		"items": bookList,
	})
}

func HandleFindBook(context *gin.Context) {
	responseCode := http.StatusOK
	bookID, _ := strconv.Atoi(context.Param("id"))
	singleBook, retrievalError := repository.FindSingleBook(database.DbConnection, bookID)

	if retrievalError.Message != "" {
		responseCode = retrievalError.Status
		context.JSON(responseCode, gin.H{
			"detail": retrievalError.Message,
		})
		return
	}

	context.JSON(responseCode, gin.H{
		"item": singleBook,
	})
}

func HandleCreateBook(context *gin.Context) {
	responseCode := http.StatusOK
	currentUser, _ := context.Get("user")
	userRecord, _ := currentUser.(structs.User)
	var newBook structs.Book
	if bindError := context.ShouldBindJSON(&newBook); bindError != nil {
		responseCode = http.StatusBadRequest
		context.JSON(responseCode, gin.H{
			"detail": bindError.Error(),
		})
		return
	}

	if newBook.ReleaseYear < 1980 || newBook.ReleaseYear > 2024 {
		responseCode = http.StatusBadRequest
		context.JSON(responseCode, gin.H{
			"detail": "Release year must be within the range 1980-2024",
		})
		return
	}

	if newBook.TotalPage > 100 {
		newBook.Thickness = "tebal"
	} else {
		newBook.Thickness = "tipis"
	}

	newBook.CreatedBy = strconv.Itoa(userRecord.ID)
	newBook.ModifiedBy = strconv.Itoa(userRecord.ID)

	creationError := repository.StoreBook(database.DbConnection, newBook)

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

func HandleUpdateBook(context *gin.Context) {
	responseCode := http.StatusOK
	bookID, _ := strconv.Atoi(context.Param("id"))
	currentUser, _ := context.Get("user")
	userRecord, _ := currentUser.(structs.User)
	var updatedBook structs.Book
	if bindError := context.ShouldBindJSON(&updatedBook); bindError != nil {
		responseCode = http.StatusBadRequest
		context.JSON(responseCode, gin.H{
			"detail": bindError.Error(),
		})
		return
	}

	if updatedBook.ReleaseYear < 1980 || updatedBook.ReleaseYear > 2024 {
		responseCode = http.StatusBadRequest
		context.JSON(responseCode, gin.H{
			"detail": "Release year must be between 1980 and 2024",
		})
		return
	}

	if updatedBook.TotalPage > 100 {
		updatedBook.Thickness = "tebal"
	} else {
		updatedBook.Thickness = "tipis"
	}

	updatedBook.ID = bookID
	updatedBook.ModifiedBy = strconv.Itoa(userRecord.ID)

	updateError := repository.UpdateExistingBook(database.DbConnection, updatedBook)

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

func HandleDeleteBook(context *gin.Context) {
	responseCode := http.StatusOK
	bookID, _ := strconv.Atoi(context.Param("id"))
	deletionError := repository.EraseBook(database.DbConnection, bookID)

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