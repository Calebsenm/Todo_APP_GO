package tasks

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"todoApp/db"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type NewTask struct {
	Title string    `json:"title" `
	Text  string    `json:"text" `
	Date  time.Time `json:"date" `
}

type Task struct {
	Id    int       `json:"id"`
	Title string    `json:"title"`
	Text  string    `json:"text"`
	Date  time.Time `json:"date"`
}

func GetTasks(c *gin.Context) {
	var tasks []Task
	// conexion to the database
	database, err := db.Connet()

	// make the Query
	data, err := database.Query("SELECT * FROM tasks")

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Pasó 3")
	// iteraty over the query
	for data.Next() {
		var ts Task
		data.Scan(&ts.Id, &ts.Title, &ts.Text, &ts.Date)
		tasks = append(tasks, ts)
		fmt.Println("XD  ", ts)
	}
	// json return
	c.JSON(http.StatusOK, tasks)

	defer database.Close()
}

func GetTask(c *gin.Context) {
	param := c.Param("id")
	// Separar el valor del prefijo ":"
	id1 := param[1:]
	id, _ := strconv.Atoi(id1)
	// conexion to the database
	database, err := db.Connet()

	// Realizar la consulta utilizando el ID obtenido
	query := "SELECT * FROM tasks"
	rows, err := database.Query(query)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al ejecutar la consulta"})
		return
	}
	defer rows.Close()

	// Iterar sobre los resultados de la consulta
	for rows.Next() {
		var ts Task
		err := rows.Scan(&ts.Id, &ts.Title, &ts.Text, &ts.Date)

		if ts.Id == id {

			c.JSON(http.StatusOK, ts)
			return
		} else {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear los resultados"})
			return
		}

	}

	// Si no se encontró ninguna tarea con el ID especificado
	c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})

}

func PostTask(c *gin.Context) {

	var task Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dab, err := db.Connet()

	if err != nil {
		log.Println(err)
	}

	insertQuery := "INSERT INTO tasks(tittle,text , date) VALUES($1, $2, $3)"
	_, err = dab.Exec(insertQuery, task.Title, task.Text, task.Title)

	if err != nil {
		fmt.Println("Error ", err)
	} 

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task);

}


func UpdateTask(c *gin.Context) {

    var task1 Task;
    data , err := db.Connet();


}
func DeleteTask(c *gin.Context) {

}




