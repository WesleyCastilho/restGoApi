package main

import (
	"github.com/WesleyCastilho/restGoApi/controllers"
	"github.com/WesleyCastilho/restGoApi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// r is a variable means route, its a Go community standard and i don't like it
	r := gin.Default()

	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
