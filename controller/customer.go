package controller

import (
	"encoding/json"
	"net/http"
	"only-test/model"
	"only-test/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (con *Controller) CreateCustomer(c *gin.Context) {
	var payload model.Customer
	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		WriteError(c, "Internal ServerError", http.StatusInternalServerError, err)
		return
	}

	res, err := con.Model.CreateCustomer(payload)
	if err != nil {
		// log.Fatal(err)
		WriteError(c, "Failed", http.StatusBadRequest, err)
		return
	}
	WriteSuccess(c, res, "Successfuly", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (con *Controller) DetailCustomer(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	res, err := con.Model.DetailCustomer(int(id))
	if err != nil {
		// log.Fatal(err)
		WriteError(c, "Failed", http.StatusNotFound, err)
		return
	}
	WriteSuccess(c, res, "Successfuly", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (con *Controller) DeleteCCustomer(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	err := con.Model.DeleteCustomer(int(id))
	if err != nil {
		// log.Fatal(err)
		WriteError(c, "Failed", http.StatusBadRequest, err)
		return
	}
	WriteSuccess(c, nil, "Successfuly", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (con *Controller) UpdateCustomer(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var payload model.Customer
	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		WriteError(c, "Internal ServerError", http.StatusInternalServerError, err)
		return
	}

	payload.ID = id

	res, err := con.Model.UpdateCustomer(payload)
	if err != nil {
		// log.Fatal(err)
		WriteError(c, "Failed", http.StatusBadRequest, err)
		return
	}
	WriteSuccess(c, res, "Successfuly", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (con *Controller) ListCustomer(c *gin.Context) {

	var limit, page int

	limit, _ = strconv.Atoi(c.Query("limit"))
	page, _ = strconv.Atoi(c.Query("page"))

	var query request.CustomerQuery
	query.Limit = limit
	query.Page = page
	query.Search = c.Query("search")
	res, err := con.Model.ListCustomer(query)
	if err != nil {
		// log.Fatal(err)
		WriteError(c, "Failed", http.StatusBadRequest, err)
		return
	}
	offset := query.GetOffset()
	WriteSuccess(c, res, "Successfuly",
		ResponseMeta{
			HTTPStatus: http.StatusOK,
			Limit:      &query.Limit,
			Page:       &query.Page,
			Offset:     &offset})

}
