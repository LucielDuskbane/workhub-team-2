package handlers

import (
	"net/http"
	"workhub/internal/dto"
	"workhub/internal/services"
	"workhub/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: services.NewAuthService(),
	}
}

// Register godoc
// @Summary Register user
// @Description Register new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register Request"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /auth/register [post]
func (h *AuthHandler) Register(
	c *gin.Context,
) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(
		&req,
	); err != nil {

		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	err := h.authService.
		Register(req)

	if err != nil {

		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"Register success",
		nil,
	)
}

// Login godoc
// @Summary Login user
// @Description Login and get JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login Request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /auth/login [post]
func (h *AuthHandler) Login(
	c *gin.Context,
) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(
		&req,
	); err != nil {

		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	response, err :=
		h.authService.
			Login(req)

	if err != nil {

		utils.ErrorResponse(
			c,
			http.StatusUnauthorized,
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Login success",
		response,
	)
}
