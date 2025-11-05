package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.Run(":8081")
}