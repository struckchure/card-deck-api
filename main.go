package main

import "api/card-deck-api/router"


// app entry

func main () {
	// run server

  router.SetupRouter().Run("localhost:8080")
}
