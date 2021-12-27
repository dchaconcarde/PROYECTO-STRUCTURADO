package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dchaconcarde/proyecto-structurado/internal/usuarios"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Email    string  `json:"email"`
	Edad     int     `json:"edad"`
	Altura   float64 `json:"altura"`
	Activo   bool    `json:"activo"`
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

func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !verifyToken(ctx) {
			ctx.JSON(401, gin.H{"error": "no tiene autorización"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		us, err := u.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, us)

	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !verifyToken(ctx) {
			ctx.JSON(401, gin.H{"error": "no tiene autorización"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		err = u.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El usuario %d ha sido eliminado", id)})
	}
}

func (u *User) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !verifyToken(ctx) {
			ctx.JSON(401, gin.H{"error": "no tiene autorización"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		us, err := u.service.UpdateName(int(id), req.Nombre, req.Edad)
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
