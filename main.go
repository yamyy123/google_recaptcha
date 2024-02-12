package main

import (
	"fmt"
	"log"
	"recaptcha/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Newroute(router)
	fmt.Println("Server started at :8080")
	log.Fatal(router.Run(":8080"))
}
