package controller

import (
	"encoding/json"
	"net/http"
	"only-test/request"
	"only-test/utils"
	"only-test/view"

	"github.com/gin-gonic/gin"
)

func (con *Controller) Login(c *gin.Context) {
	var payload request.LoginRequest

	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		WriteError(c, "Internal ServerError", http.StatusInternalServerError, err)
		return
	}

	user, err, _ := con.Model.CheckUser(payload)
	if err != nil {
		WriteError(c, "Email and Password Invalid", http.StatusUnauthorized, err)
		return
	}

	isValid, err := utils.ComparePassword(user.Password, []byte(payload.Password))
	if !isValid {
		WriteError(c, "Email and Password Invalid", http.StatusUnauthorized, err)
		return
	}

	token, err := utils.CreateToken(user.ID)
	if err != nil {
		WriteError(c, "Error Create Token", http.StatusInternalServerError, err)
		return
	}

	res := view.LoginResponse{
		User:  user,
		Token: token,
	}

	WriteSuccess(c, res, "Login Successfully", ResponseMeta{HTTPStatus: http.StatusOK})
}
