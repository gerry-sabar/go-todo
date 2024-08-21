package main

import (
	"fmt"
	"go-todo/config"
	"go-todo/database"
	"go-todo/repository"
	"go-todo/server"
	"go-todo/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	config.LoadConfig()
	config.DB, err = database.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer config.DB.Close()

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoServiceImpl(todoRepository)
	UserRepository := repository.NewUserRepository()
	userService := service.NewUserServiceImpl(UserRepository)

	srv := server.RouterService{
		TodoService: todoService,
		UserService: userService,
	}

	r := server.SetupRouter(srv)
	r.Run(fmt.Sprintf(":%s", config.Cfg.APPPort))
}
