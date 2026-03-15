package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"arna/internal/handlers"
	"arna/internal/store"
)

func TestCreateTask(t *testing.T) {

	gin.SetMode(gin.TestMode)

	store := store.NewMemoryStore()
	handler := handlers.NewTaskHandler(store)

	router := gin.Default()
	router.POST("/tasks", handler.CreateTask)

	body := `{
		"title": "test task"
	}`

	req, _ := http.NewRequest(
		http.MethodPost,
		"/tasks",
		bytes.NewBuffer([]byte(body)),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 got %d", w.Code)
	}
}