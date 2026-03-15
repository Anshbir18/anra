package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"arna/internal/models"
	"arna/internal/store"
	"arna/internal/validation"
)

type TaskHandler struct {
	store store.TaskStore
}

func NewTaskHandler(s store.TaskStore) *TaskHandler {
	return &TaskHandler{store: s}
}

type CreateTaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

// CreateTask godoc
// @Summary Create task
// @Description Create a new task
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body CreateTaskRequest true "Task"
// @Success 201 {object} models.Task
// @Failure 400 {object} map[string]string
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if err := validation.ValidateTitle(req.Title); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	status := req.Status
	if status == "" {
		status = "todo"
	}

	if !validation.ValidateStatus(status) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid status value",
		})
		return
	}

	task := models.Task{
		ID:     uuid.New().String(),
		Title:  req.Title,
		Status: models.TaskStatus(status),
	}

	h.store.Create(task)

	c.JSON(http.StatusCreated, task)
}

// ListTasks godoc
// @Summary List tasks
// @Description Get all tasks
// @Tags tasks
// @Produce json
// @Success 200 {array} models.Task
// @Router /tasks [get]
func (h *TaskHandler) ListTasks(c *gin.Context) {
	tasks := h.store.List()

	c.JSON(http.StatusOK, tasks)
}