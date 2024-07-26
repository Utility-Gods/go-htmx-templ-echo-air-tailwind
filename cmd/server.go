package main

import (
	"fmt"
	"net/http"

	"github.com/Utility-Gods/photoship-go/internal/app/template"
	"github.com/a-h/templ"
)

func main() {
	component := template.Home("PhotoShip")
	http.Handle("/", templ.Handler(component))
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
