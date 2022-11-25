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

type SuccessResponseWithToken struct {
	Status    int    `json:"status"`
	Success   bool   `json:"success"`
	Data      any    `json:"data"`
	Token     string `json:"token"`
	TimeStamp string `json:"time_stamp"`
}

type SuccessResponseWithMessage struct {
	Status    int    `json:"status"`
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	TimeStamp string `json:"time_stamp"`
}

func SendSuccessResponse(c *gin.Context, status int, data any, opts ...string) {

	var successRes any

	if status == 0 {
		status = 200
	}

	if len(opts) != 0 {
		switch opts[0] {

		case "token":
			successRes = SuccessResponseWithToken{
				Status:    status,
				Success:   true,
				Data:      data,
				Token:     string(opts[1]),
				TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
			}
		case "message":
			successRes = SuccessResponseWithMessage{
				Status:    status,
				Success:   true,
				Message:   string(opts[1]),
				TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
			}
		default:
			successRes = SuccessResponseWithMessage{
				Status:    status,
				Success:   true,
				Message:   string(opts[1]),
				TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
			}
		}

	} else {
		successRes = SuccessResponse{
			Status:    status,
			Success:   true,
			Data:      data,
			TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
		}

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
