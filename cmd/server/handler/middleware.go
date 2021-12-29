package handler

import (
	"os"

	"github.com/dchaconcarde/proyecto-structurado/pkg/web"
	"github.com/gin-gonic/gin"
)

func NewMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token == "" {
		ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "ingresar token"))
	} else if token != os.Getenv("TOKEN") {
		ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token inválido"))
	}
	ctx.Next()
}
