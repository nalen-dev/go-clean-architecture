package rest

import "github.com/labstack/echo/v4"

func LoadRoutes(e *echo.Echo, handler *handler){
	// TODO: add auth midleware

	userGroup :=  e.Group("/user")
	userGroup.POST("/login",handler.Login)

}