package routes

import (
	"github.com/MashCyber/ext-apis/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// routes
	r.GET("/pizza/:id", handlers.FetchPizza)

	//Norris
	r.GET("/norris/random", handlers.NorrisRandom)
	// r.GET("/norris/categ:category", handlers.NorrisCategory)
}
