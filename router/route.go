package router

import (
	"only-test/controller"
	"only-test/model"
	"only-test/request"
	"only-test/utils"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	db, err := utils.ConnectSql()
	if err != nil {
		panic("Failed to connect to the database")
	}
	//migrate
	db.AutoMigrate(&model.User{}, &model.Customer{}, &model.Order{})
	//end migrate
	models := model.NewModel(db)
	//inite user
	req := model.User{
		Name:     "Admmi",
		Email:    "admin@gmail.com",
		Password: "admin123",
	}
	var login request.LoginRequest
	login.Email = req.Email
	_, _, exist := models.CheckUser(login)
	if !exist {
		_, _ = models.CreateUser(req)
	}
	///end init user

	con := controller.NewController(&models)

	r := gin.Default()

	r.GET("/healtz", con.Healtz)
	r.POST("/login", con.Login)
	protect := r.Group("/")
	protect.Use(con.TokenValid)
	protect.POST("/user", con.CreateUser)

	customer := protect.Group("/customer")
	customer.GET("/", con.ListCustomer)
	customer.POST("/", con.CreateCustomer)
	customer.GET("/:id", con.DetailCustomer)
	customer.PUT("/:id", con.UpdateCustomer)
	customer.DELETE("/:id", con.DeleteCCustomer)

	order := protect.Group("/order")
	order.GET("/", con.ListOrder)
	order.POST("/", con.CreateOrder)
	order.GET("/:id", con.DetailOrder)
	order.PUT("/:id", con.UpdateOrder)
	order.DELETE("/:id", con.DeleteOrder)

	return r
}
