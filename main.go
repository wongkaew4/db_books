package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wongkaew4/godb_book/handle"
	"github.com/wongkaew4/godb_book/model"
)

func main() {
	r := gin.Default()
	b := handle.NewBook()

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, b.AllData())
	})

	r.POST("/book", func(c *gin.Context) {
		var v model.Book
		err := c.ShouldBindJSON(&v)

		if err != nil {
			c.JSON(400, gin.H{
				"message": "Error Format",
			})
			return

		}
		b.AddData(v)
		c.JSON(200, v)
	})

	r.Run(":4433") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
