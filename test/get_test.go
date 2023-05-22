package tasks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"todoApp/db"
    "todoApp/tasks"
)

func TestGetTasks(t *testing.T) {
	// Preparar el entorno
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/tasks", tasks.GetTasks)

	// Crear una solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	router.ServeHTTP(w, req)

	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificar el contenido de la respuesta JSON
	expectedJSON := `[]` // Cambiar aquí si se esperan datos específicos
	assert.JSONEq(t, expectedJSON, w.Body.String())
}

func TestConnectPost(t *testing.T) {
	// Realizar la prueba de conexión a la base de datos
	db, err := db.ConnectPost()

	// Verificar que no haya errores en la conexión
	assert.NoError(t, err)

	// Verificar que se haya obtenido una instancia válida de la base de datos
	assert.NotNil(t, db)

	// Cerrar la conexión
	db.Close()
}
