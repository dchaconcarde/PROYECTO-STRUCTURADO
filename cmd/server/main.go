package main

import (
	"os"

	"github.com/dchaconcarde/proyecto-structurado/cmd/server/handler"
	"github.com/dchaconcarde/proyecto-structurado/docs"
	"github.com/dchaconcarde/proyecto-structurado/internal/usuarios"
	"github.com/dchaconcarde/proyecto-structurado/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Users.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @licence.name Apache 2.0
// @license.url http://www.apache.org/licences/LICENCE-2.0.html

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "users.json")
	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	us := handler.NewUser(service)
	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	usersGroup := r.Group("/users")
	usersGroup.Use(handler.NewMiddleware)
	usersGroup.POST("/", us.Store())
	usersGroup.GET("/", us.GetAll())
	usersGroup.PUT("/:id", us.Update())
	usersGroup.DELETE("/:id", us.Delete())
	usersGroup.PATCH("/:id", us.UpdateName())

	r.Run()
}
