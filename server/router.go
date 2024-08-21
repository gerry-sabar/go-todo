package server

import (
	"go-todo/handler"
	"go-todo/service"

	"github.com/gin-gonic/gin"
)

type RouterService struct {
	TodoService service.TodoService
	UserService service.UserService
}

func SetupRouter(srv RouterService) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		todoHandler := handler.NewTodoHandler(srv.TodoService)
		v1.GET("/todos", todoHandler.GetAll())           // list items
		v1.POST("/todos", todoHandler.CreateItem())      // create item
		v1.GET("/todos/:id", todoHandler.GetById())      // get an item by ID
		v1.PUT("/todos/:id", todoHandler.EditItemById()) // edit an item by ID
		v1.DELETE("/todos/:id", todoHandler.Delete())    // delete an item by ID

		userHandler := handler.NewUserHandler(srv.UserService)
		v1.GET("/users", userHandler.GetAll())           // list users
		v1.POST("/users", userHandler.CreateItem())      // create user
		v1.GET("/users/:id", userHandler.GetUserById())  // get a user by ID
		v1.PUT("/users/:id", userHandler.EditUserById()) // edit a user by ID
		v1.DELETE("/users/:id", userHandler.Delete())    // delete a user by ID

	}

	return router
}
