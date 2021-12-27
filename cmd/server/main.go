package main

import (
	"github.com/dchaconcarde/proyecto-structurado/cmd/server/handler"
	"github.com/dchaconcarde/proyecto-structurado/internal/usuarios"
	"github.com/dchaconcarde/proyecto-structurado/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "products.json")
	repo := usuarios.NewRepository(db)
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
