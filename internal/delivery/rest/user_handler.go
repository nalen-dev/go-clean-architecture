package rest

import (
	"encoding/json"
	"go-clean-architecture/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func(h *handler) Login(c echo.Context) error {

	var request model.UserCredentials
	err := json.NewDecoder(c.Request().Body).Decode(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userSession, err := h.userUsecase.Login(c.Request().Context(), request)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": userSession,
	})
}