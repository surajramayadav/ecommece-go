package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status    int       `json:"status"`
	Success   bool      `json:"success"`
	Message   string    `json:"message"`
	TimeStamp time.Time `json:"time_stamp"`
}

func (errorResponse *ErrorResponse) defaultValues() {
	errorResponse.Success = false
	errorResponse.TimeStamp = time.Now()
}

func SendErrorResponse(c *gin.Context, status int, message string) {

	errorData := ErrorResponse{
		Status:  status,
		Message: message,
	}
	errorData.defaultValues()

	switch status {
	case 400:
		StatusBadRequest(c, errorData)
	case 401:
		StatusUnauthorized(c, errorData)
	case 404:
		StatusNotFound(c, errorData)
	case 500:
		StatusInternalServerError(c, errorData)
	default:
		StatusBadRequest(c, errorData)
	}
}

func StatusBadRequest(c *gin.Context, data ErrorResponse) {
	c.JSON(http.StatusBadRequest, data) //400
}

func StatusNotFound(c *gin.Context, data ErrorResponse) {
	c.JSON(http.StatusNotFound, data) //404
}

func StatusUnauthorized(c *gin.Context, data ErrorResponse) {
	c.JSON(http.StatusUnauthorized, data) //401
}

func StatusInternalServerError(c *gin.Context, data ErrorResponse) {
	c.JSON(http.StatusInternalServerError, data) // 500
}
