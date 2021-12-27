package main

import (
	"github.com/dchaconcarde/proyecto-structurado/cmd/server/handler"
	"github.com/dchaconcarde/proyecto-structurado/internal/usuarios"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := usuarios.NewRepository()
	service := usuarios.NewService(repo)
	us := handler.NewUser(service)

	r := gin.Default()
	usersGroup := r.Group("/users")
	usersGroup.POST("/", us.Store())
	usersGroup.GET("/", us.GetAll())
	usersGroup.PUT("/:id", us.Update())
	usersGroup.DELETE("/:id", us.Delete())
	usersGroup.PATCH("/:id", us.UpdateName())

	r.Run()
}
