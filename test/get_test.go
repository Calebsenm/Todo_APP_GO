package tasks

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoApp/tasks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {

	router := gin.Default()
	router.GET("/tasks", tasks.GetTask)

	w := httptest.NewRecorder()
	assert.Equal(t, http.StatusOK, w.Code);

	fmt.Println("body", )
}
