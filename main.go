package main 

import(
	"todoApp/tasks"
	"github.com/gin-gonic/gin"
)

func main(){
	
	router := gin.Default();
    
    router.GET("/",func (c *gin.Context ){
        c.File("./fronted/index.html");
    });

	router.GET("/tasks",tasks.GetTasks);
	router.GET("/task:id",tasks.GetTask);
	router.POST("/task",tasks.PostTask);
	router.PUT("/task/:id",tasks.UpdateTask);
	router.DELETE("/task/:id",tasks.DeleteTask);

	router.Run("localhost:8000");
}
