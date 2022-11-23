package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Status    int    `json:"status"`
	Success   bool   `json:"success"`
	Data      any    `json:"data"`
	TimeStamp string `json:"time_stamp"`
}

func SendSuccessResponse(c *gin.Context, status int, data any) {
	if status == 0 {
		status = 200
	}

	successRes := SuccessResponse{
		Status:    status,
		Success:   true,
		Data:      data,
		TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	switch status {
	case 200:
		StatusOk(c, successRes)
	case 201:
		StatusCreated(c, successRes)
	default:
		StatusOk(c, successRes)
	}

}

func StatusCreated(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, data) //201
}

func StatusOk(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data) //200

}
