package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
    "todoApp/tasks"
)

func TestGetTask(t *testing.T) {

	router := gin.Default()
	router.GET("/task:id",tasks.GetTask )
	w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/task:1", nil)
	router.ServeHTTP(w, req)

	expected := `{"Id":1,"Title":"Web","Text":"Golang","Date":"2023-06-11T00:00:00Z"}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}


