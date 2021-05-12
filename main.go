package main

import "github.com/gin-gonic/gin"

func main() {
	// Initialize the router.
	router := gin.Default()

	// Show a message indicating successful deployment.
	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Successfully deployed!",
		})
	})

	// Host the stellar.toml file.
	router.StaticFile("/.well-known/stellar.toml", "./stellar.toml")

	// Run the app.
	router.Run()
}
