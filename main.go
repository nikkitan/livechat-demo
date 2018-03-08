package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gernest/utron"
	c "github.com/nikkitan/livechat-demo/controllers"
)

func main() {

	// Start the MVC App
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	// Register Controller
	app.AddController(c.NewLiveChatController)

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	log.Fatal(http.ListenAndServe(port, app))
}
