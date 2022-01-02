package handler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dchaconcarde/proyecto-structurado/internal/usuarios"
	"github.com/dchaconcarde/proyecto-structurado/pkg/web"
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

// ListUsers godoc
// @Summary List users
// @Tags Users
// @Description get users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users [get]
func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		us, err := u.service.GetAll()

		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}

		if us == nil {
			ctx.JSON(400, web.NewResponse(404, nil, "No existe ningún usuario en el contexto"))
		}
		ctx.JSON(200, web.NewResponse(200, us, ""))
	}
}

// StoreUsers godoc
// @Summary Store users
// @Tags Users
// @Description store users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param user body request true "User to store"
// @Success 200 {object} web.Response
// @Router /users [post]
func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}

		err := verifyFields(ctx, req)
		if err != nil {
			return
		}
		us, err := u.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, us, ""))
	}
}

// UpdateUsers godoc
// @Summary Update users
// @Tags Users
// @Description Update user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param user body request true "User to update"
// @Param string query string true "User ID to Update"
// @Success 200 {object} web.Response
// @Router /users/{id} [put]
func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, "Invalid ID"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}

		err = verifyFields(ctx, req)
		if err != nil {
			return
		}

		us, err := u.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, us, ""))

	}
}

// DeleteUsers godoc
// @Summary Delete users
// @Tags Users
// @Description delete user by id
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param string query string true "User ID to delete"
// @Success 200 {object} web.Response
// @Router /users/{id} [delete]
func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, "Invalid ID"))
			return
		}

		err = u.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El usuario %d ha sido eliminado", id), ""))
	}
}

type updateNameRequest struct {
	Apellido string
	Edad     int
}

// UpdateNameUsers godoc
// @Summary UpdateName users
// @Tags Users
// @Description update lastname and age for a user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param string query string true "User ID to update lastName and age"
// @Param user body updateNameRequest true "User lastName and age"
// @Success 200 {object} web.Response
// @Router /users/{id} [patch]
func (u *User) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, "Invalid ID"))
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}

		err = verifyFields(ctx, req)
		if err != nil {
			return
		}

		us, err := u.service.UpdateName(int(id), req.Apellido, req.Edad)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, us, ""))

	}
}

func verifyFields(ctx *gin.Context, req request) error {
	if req.Nombre == "" {
		ctx.JSON(400, web.NewResponse(401, nil, "El nombre de usuario es requerido"))
		return errors.New("Error")
	}
	if req.Apellido == "" {
		ctx.JSON(400, web.NewResponse(401, nil, "El apellido de usuario es requerido"))
		return errors.New("Error")
	}
	if req.Email == "" {
		ctx.JSON(400, web.NewResponse(401, nil, "El email de usuario es requerido"))
		return errors.New("Error")
	}
	if req.Edad == 0 {
		ctx.JSON(400, web.NewResponse(401, nil, "La edadde usuario es requerido"))
		return errors.New("Error")
	}
	if !req.Activo {
		ctx.JSON(400, web.NewResponse(401, nil, "El estado de usuario es requerido"))
		return errors.New("Error")
	}
	return nil

}
