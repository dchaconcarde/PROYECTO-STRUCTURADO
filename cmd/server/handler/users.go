package handler

import (
	"net/http"

	"github.com/dchaconcarde/proyecto-structurado/cmd/internal/usuarios"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre   string  `json:"nombre" binding:"required"`
	Apellido string  `json:"apellido" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Edad     int     `json:"edad" binding:"required"`
	Altura   float64 `json:"altura" binding:"required"`
	Activo   bool    `json:"activo" binding:"required"`
}

type User struct {
	service usuarios.Service
}

func NewUser(service usuarios.Service) *User {
	return &User{
		service: service,
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !verifyToken(ctx) {
			ctx.JSON(401, gin.H{"error": "no tiene autorización"})
			return
		}

		us, err := u.service.GetAll()

		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		if us == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No existe ningún usuario en el contexto"})
		}
		ctx.JSON(200, us)
	}
}

func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !verifyToken(ctx) {
			ctx.JSON(401, gin.H{"error": "no tiene autorización"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		us, err := u.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, us)
	}
}

func verifyToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	return token == usuarios.TokenToPass
}
