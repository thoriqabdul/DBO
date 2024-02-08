package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"only-test/model"

	"github.com/gin-gonic/gin"
)

func (con *Controller) CreateUser(c *gin.Context) {
	var payload model.User
	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		log.Fatal(err)
		WriteError(c, "Internal ServerError", http.StatusInternalServerError, err)
		return
	}

	res, err := con.Model.CreateUser(payload)
	if err != nil {
		// log.Fatal(err)
		WriteError(c, "Failed", http.StatusBadRequest, err)
		return
	}

	WriteSuccess(c, res, "Successfuly", ResponseMeta{HTTPStatus: http.StatusOK})
}
