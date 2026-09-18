package main

import (
	"github.com/MashCyber/ext-apis/routes"
	"github.com/gin-gonic/gin"
)

// use gin to request 3rd-party api and respond with the api response

func main() {
	r := gin.Default()

	//register routes

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
