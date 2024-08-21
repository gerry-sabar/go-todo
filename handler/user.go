package handler

import (
	"go-todo/domain"
	"go-todo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User interface {
	Delete() gin.HandlerFunc
	EditUserById() gin.HandlerFunc
	GetUserById() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	CreateUser() gin.HandlerFunc
}

type UserService struct {
	User domain.User
}

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.UserService.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func (h *UserHandler) EditUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userReq := domain.User{}
		if err := c.ShouldBind(&userReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.UserService.Update(id, userReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}

func (h *UserHandler) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := h.UserService.GetById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func (h *UserHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := h.UserService.GetAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": items})
	}
}

func (h *UserHandler) CreateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		userReq := domain.User{}
		if err := c.ShouldBind(&userReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := h.UserService.Create(userReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
