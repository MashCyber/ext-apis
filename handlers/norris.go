package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MashCyber/ext-apis/models"
	"github.com/gin-gonic/gin"
)

func NorrisRandom(c *gin.Context) {
	url := "https://api.chucknorris.io/jokes/random"

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Failed to reach external API",
		})
		return
	}
	defer resp.Body.Close()

	var random models.Norris
	if err := json.NewDecoder(resp.Body).Decode(&random); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to parse external respose.",
		})
		return
	}

	//success
	c.JSON(http.StatusOK, random)
}
