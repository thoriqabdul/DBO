package controller

import (
	"net/http"
	"only-test/model"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Model *model.Model
}

func NewController(model *model.Model) Controller {
	return Controller{Model: model}
}

func (con *Controller) Healtz(c *gin.Context) {
	WriteSuccess(c, nil, "Successfuly", ResponseMeta{HTTPStatus: http.StatusOK})

}

type ResponseMeta struct {
	HTTPStatus int  `json:"http_status"`
	Total      *int `json:"total,omitempty"`
	Offset     *int `json:"offset,omitempty"`
	Limit      *int `json:"limit,omitempty"`
	Page       *int `json:"page,omitempty"`
	LastPage   *int `json:"last_page,omitempty"`
}

func WriteError(c *gin.Context, message string, errorCode int, err error) {

	c.JSON(errorCode, gin.H{
		"status":  "error",
		"message": message,
		"error":   err.Error(),
	})
}

func WriteSuccess(c *gin.Context, data interface{}, message string, meta ResponseMeta) {
	c.JSON(200, gin.H{
		"message": message,
		"data":    data,
		"meta":    meta,
	})
}
