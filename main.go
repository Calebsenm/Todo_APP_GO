package main

import (
	"github.com/gin-gonic/gin"
	"todoApp/tasks"
)

func main() {

	router := gin.Default()

	router.Static("src", "./src/")

	router.GET("/", func(c *gin.Context) {
		c.File("./src/")
	})

	router.GET("/tasks", tasks.GetTasks)
	router.GET("/task:id", tasks.GetTask)
	router.POST("/task", tasks.PostTask)
	router.PUT("/task/:id", tasks.UpdateTask)
	router.DELETE("/task/:id", tasks.DeleteTask)

	router.Run("localhost:8000")

}
