
package main

import (
    "encoding/json"
    "net/http"
	"net/http/httptest"
	"testing"
    "todoApp/db"
    "todoApp/tasks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

)

func TestGetTasks(t *testing.T) {
	// Crea un enrutador de Gin
	router := gin.Default()

	// Registra el manejador para la ruta GET /tasks
	router.GET("/tasks", tasks.GetTasks)

	// Crea un ResponseRecorder falso para capturar la respuesta
	w := httptest.NewRecorder()

	// Crea una solicitud GET de prueba para /tasks
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Envía la solicitud al enrutador
	router.ServeHTTP(w, req)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)


    // 
    var ListTasks  []db.Task
    json.Unmarshal(w.Body.Bytes(), &ListTasks);
    assert.NotEmpty(t, ListTasks);

}




