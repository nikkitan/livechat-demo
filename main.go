package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gernest/utron"
	c "github.com/nikkitan/livechat-demo/controllers"
)

func main() {
	app, err := utron.NewMVC()
	port := fmt.Sprintf(":%d", app.Config.Port)

	if err != nil {
		log.Fatal(err)
	}

	app.AddController(c.NewCommonController)
	app.AddController(c.NewLiveChatController)
	app.AddController(c.NewChatWgtController)

	log.Fatal(http.ListenAndServe(port, app))
}
