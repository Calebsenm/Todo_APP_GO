package tasks

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todoApp/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
    "time"
)

type Task struct {
	Id    int       `json:"Id"`
	Title string    `json:"Title"`
	Text  string    `json:"Text"`
	Date  time.Time `json:"Date"`
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

// Correct 
// function make the post data

func PostTask(c *gin.Context) {
	var task Task
    err := c.ShouldBindJSON(&task);

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		fmt.Println(task);
        return
	}

	dab, err := db.Connet()

	if err != nil {
		log.Println(err)
	}

	insertQuery := "INSERT INTO tasks(id ,tittle,text , date) VALUES($1, $2, $3 , $4)"
	_, err = dab.Exec(insertQuery,task.Id,task.Title, task.Text, task.Date)


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task);
}

// To Do do the update to the database
func UpdateTask(c *gin.Context) {
    var task1 Task;

    err := c.ShouldBindJSON(&task1);

    if err != nil{
        fmt.Println("Error 1", err);
    }
    
    // the database conexion
    data , err := db.Connet();
    defer data.Close();

    if err != nil {
        fmt.Println("Error2",err)
    }
    
    prepareStatement , err := data.Prepare("UPDATE tasks SET tittle = $1, text = $2, date = $3 WHERE id=$4");
    if err != nil{
        fmt.Println("Error3", err )
    }
    
    prepareStatement.Exec(task1.Title, task1.Text, task1.Date, task1.Id); 
    c.JSON(http.StatusOK, task1)
}


func DeleteTask(c *gin.Context) {

    idText := c.Param("id");
    id , _ := strconv.Atoi(idText[1:]);
    
    data , err := db.Connet();
    if err != nil {
        fmt.Println("Error1",err);
    }
    defer data.Close();
    
    theData, err  :=  data.Exec("DELETE FROM tasks WHERE id=$1",id);
    if err != nil {
        fmt.Println("Error2",err);
    }

    fmt.Println(theData,"-", id);
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
    
}




