package router

import "github.com/gin-gonic/gin"

import "api/card-deck-api/controllers"


// mode -> developement (true) / production (false)

const DEBUG bool = true


func SetupRouter () *gin.Engine {
	// set DEBUG mode to production / developement

	if !DEBUG {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin router instance

	routes := gin.Default()

	routes.POST("/create-deck", controllers.CreateDeck) // `CreateDeck` endpoint
	routes.GET("/open-deck/:deck_id", controllers.OpenDeck) // `OpenDeck` endpoint
	routes.GET("/draw-card", controllers.DrawCard) // `DrawCard` endoint

	return routes
}
