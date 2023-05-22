package tasks

import (
	"net/http"
	"github.com/gin-gonic/gin"
    "time"
    "log"
    "todoApp/db"
    "fmt"
)
type NewTask struct{
    Title       string `json:"title" binding:"required"` 
    Text        string `json:"text" binding:"required"`
    Date        time.Time `json:"date" binding: "required"`
}

type Task struct{
    Title      string     `json:"title"`
    Text       string     `json:"text"`
    Date       time.Time  `json:"date"`
}


func GetTasks( c *gin.Context ){
    var tasks []Task 
    // conexion to the database 
    database , err := db.ConnectPost();
    defer database.Close();
    
    // make the Query  
    data, err := database.Query("SELECT * FROM tasks");
    
    if err != nil {
        log.Println(err);
    }
    
    // iteraty over the query 
    for data.Next(){
        var ts Task 
        data.Scan(&ts.Title, &ts.Text , &ts.Date)
        tasks = append(tasks, ts )
        fmt.Println("XD  ",ts)
    }

    // json return 
    c.JSON(http.StatusOK, tasks);
}

func GetTask( c *gin.Context){

}

func PostTask(c *gin.Context){

}
func UpdateTask( c *gin.Context){

}
func DeleteTask( c *gin.Context){

}


