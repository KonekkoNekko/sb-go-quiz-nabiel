package middleware

import (
	"database/sql"
	"net/http"
	"sb-go-quiz-nabiel/structs"

	"github.com/gin-gonic/gin"
)

func Authenticate(dbs *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		name, secret, present := context.Request.BasicAuth()

		if !present {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"detail": "Authentication required",
			})
			return
		}

		var account structs.User

		statement := "SELECT * FROM users WHERE username = $1 AND password = $2"
		err := dbs.QueryRow(statement, name, secret).Scan(&account.ID, &account.Username, &account.Password, &account.CreatedAt, &account.CreatedBy, &account.ModifiedAt, &account.ModifiedBy)

		if err != nil {
			if err == sql.ErrNoRows {
				context.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect credentials"})
			} else {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			context.Abort()
			return
		}

		context.Set("account", account)
		context.Next()
	}
}