package restServer

import (
	"github.com/gin-gonic/gin"
	"github.com/miguelToscano/newsletter/internal/controllers"
)

func Start() {
	usersController := controllers.NewUsersController()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/users", usersController.GetUsers)

	r.Run(":8080") // Listen and serve on
}
