package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := ServeRouter()
	err := router.Run()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func ServeRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"poronga": 123, "sarasa": "jajs dice poronga"})
	})

	return r
}
