package bootstrap

import (
	"fmt"

	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() {

	e := gin.Default()
	configureRoutes(e)

	err := e.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server Started")
}

func configureRoutes(e *gin.Engine) {
	g := e.Group("/api/v1")
	{
		// exemplo
		// em go nós temos sempre o nosso contexto quando precisaremos
		g.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Pong",
			})
		})

	}
}
