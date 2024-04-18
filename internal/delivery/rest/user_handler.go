package rest

import (
	"encoding/json"
	"go-clean-architecture/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func(h *handler) Login(c *gin.Context) {

	var request model.UserCredentials
	err := json.NewDecoder(c.Request.Body).Decode(&request)

	if err != nil {
		 c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	userSession, err := h.userUsecase.Login(c.Request.Context(), request)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	 c.JSON(http.StatusOK, map[string]interface{}{
		"data": userSession,
	})
	return
}