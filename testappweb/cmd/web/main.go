package main

import (
	"fmt"
	"net/http"

	"github.com/valdoth/mygoprogram/config"
	"github.com/valdoth/mygoprogram/internal/handlers"
)

func main() {
	var appConfig config.Config

	templateCache, err := handlers.CreateTemplateCache()

	if err != nil {
		panic(err)
	}

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8000"

	handlers.CreateTemplates(&appConfig)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("http://localhost:8000 - Server started on port: ", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}
