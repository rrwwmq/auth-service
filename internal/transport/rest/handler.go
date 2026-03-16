package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrwwmq/auth-service/internal/service"
)

type Handler struct {
	authService *service.AuthService
}

func NewHandler(authService *service.AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
	}

	return r
}

func (h *Handler) register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	if err := h.authService.Register(req.Email, req.Password); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusCreated, MessageResponse{"Успешная регистрация!"})
}

func (h *Handler) login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	if err := h.authService.Login(req.Email, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{"Вы успешно вышли в систему!"})
}
