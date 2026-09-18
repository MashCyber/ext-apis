package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MashCyber/ext-apis/models"
	"github.com/gin-gonic/gin"
)

func FetchPizza(c *gin.Context) {
	// /pizza/if.json
	status, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldnt parse id!",
		})
		return
	}

	// url + param
	url := fmt.Sprintf("https://status.pizza/%d.json", status)

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Failed to reach external API",
		})
		return
	}

	defer resp.Body.Close()

	var pizza models.Pizza
	if err := json.NewDecoder(resp.Body).Decode(&pizza); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse external response"})
		return
	}

	c.JSON(http.StatusOK, pizza)
}
