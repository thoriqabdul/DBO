package controller

import (
	"encoding/json"
	"net/http"
	"only-test/model"
	"only-test/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (con *Controller) CreateOrder(c *gin.Context) {
	var payload model.Order
	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		WriteError(c, "Internal ServerError", http.StatusInternalServerError, err)
		return
	}

	res, err := con.Model.CreateOrder(payload)
	if err != nil {
		// log.Fatal(err)
		WriteError(c, "Failed", http.StatusBadRequest, err)
		return
	}
	WriteSuccess(c, res, "Successfuly", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (con *Controller) DetailOrder(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	res, err := con.Model.DetailOrder(int(id))
	if err != nil {
		// log.Fatal(err)
		WriteError(c, "Failed", http.StatusNotFound, err)
		return
	}
	WriteSuccess(c, res, "Successfuly", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (con *Controller) DeleteOrder(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	err := con.Model.DeleteOrder(int(id))
	if err != nil {
		// log.Fatal(err)
		WriteError(c, "Failed", http.StatusBadRequest, err)
		return
	}
	WriteSuccess(c, nil, "Successfuly", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (con *Controller) UpdateOrder(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var payload model.Order
	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		WriteError(c, "Internal ServerError", http.StatusInternalServerError, err)
		return
	}

	payload.ID = id

	res, err := con.Model.UpdateOrder(payload)
	if err != nil {
		// log.Fatal(err)
		WriteError(c, "Failed", http.StatusBadRequest, err)
		return
	}
	WriteSuccess(c, res, "Successfuly", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (con *Controller) ListOrder(c *gin.Context) {

	var limit, page int

	limit, _ = strconv.Atoi(c.Query("limit"))
	page, _ = strconv.Atoi(c.Query("page"))

	var query request.OrderQuery
	query.Limit = limit
	query.Page = page
	query.Search = c.Query("search")
	res, err := con.Model.ListOrder(query)
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
