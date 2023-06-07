package tasks

import (


	"net/http"
	"net/http/httptest"
	"testing"
	"todoApp/tasks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {

	router := gin.Default();
	router.GET("/tasks", tasks.GetTask,nil);
	w := httptest.NewRecorder();

    req, _ := http.NewRequest("GET","/tasks",nil);
    router 



	assert.Equal(t, http.StatusOK, w.Code);

}
