package main

import (
	"fmt"
	"net/http"

	"github.com/monstrong/proyektnaya-practica/src/pkg/handlers"
)

var portNumber = ":8080"

func main() {
	// Обработчик статических файлов (CSS, JS, изображения)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/journal", handlers.Journal)
	http.HandleFunc("/team", handlers.Team)
	http.HandleFunc("/resources", handlers.Resources)
	http.HandleFunc("/", handlers.ServerPage)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
