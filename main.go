package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase()
	defer DB.Close()

	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	e.GET("/", func(c *gin.Context) {
		todos := ReadToDoList()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	e.POST("/todos", func(c *gin.Context) {
		title := c.PostForm("title")
		status := c.PostForm("status")
		id, _ := CreateToDo(title, status)

		c.HTML(http.StatusOK, "task.html", gin.H{
			"Title":  title,
			"Status": status,
			"ID":     id,
		})
	})

	e.DELETE("/todos/:id", func(c *gin.Context) {
		param := c.Param("id")
		id, _ := strconv.ParseInt(param, 10, 64)
		DeleteToDo(id)
	})

	e.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
