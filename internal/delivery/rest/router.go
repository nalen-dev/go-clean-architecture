package rest

import (
	"github.com/gin-gonic/gin"
)

func LoadRoutes(e *gin.Engine, handler *handler){
	// TODO: add auth midleware

	userGroup :=  e.Group("/user")
	userGroup.POST("/login", handler.Login)

}