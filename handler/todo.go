package handler

import (
	"net/http"
	"strconv"

	"go-todo/domain"
	"go-todo/service"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Delete() gin.HandlerFunc
	EditItemById() gin.HandlerFunc
	GetById() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	CreateItem() gin.HandlerFunc
}

type TodoService struct {
	TodoItem domain.Todo
}

type TodoHandler struct {
	TodoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *TodoHandler {
	return &TodoHandler{TodoService: todoService}
}

func (h *TodoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.TodoService.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func (h *TodoHandler) EditItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		itemReq := domain.Todo{}
		if err := c.ShouldBind(&itemReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.TodoService.Update(id, itemReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}

func (h *TodoHandler) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo, _ := h.TodoService.GetById(id)
		c.JSON(http.StatusOK, gin.H{"data": todo})
	}
}

func (h *TodoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		items, _ := h.TodoService.GetAll()
		c.JSON(http.StatusOK, gin.H{"data": items})
	}
}

func (h *TodoHandler) CreateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		itemReq := domain.Todo{}
		if err := c.ShouldBind(&itemReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := h.TodoService.Create(itemReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
