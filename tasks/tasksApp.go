package tasks

import (
	//"net/http"
	"github.com/gin-gonic/gin"
    "time"
)
type NewTask struct{
	Tittle      string `json:"tittle"`
	Text string `json:"text"`
    Date        time.Time `json:"date"`
}

func GetTasks( c *gin.Context ){

}

func GetTask( c *gin.Context){

}

func PostTask(c *gin.Context){

}
func UpdateTask( c *gin.Context){

}
func DeleteTask( c *gin.Context){

}
