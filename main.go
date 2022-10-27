package main

import (
	"html/template"
	//"log"
	"fmt"
	"net/http"
)

// plantilla
var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)

	fmt.Println("Running Service...")

	http.ListenAndServe(":8080", nil)
}
func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello Saiyans!!!")
	plantillas.ExecuteTemplate(w, "inicio", nil)

}

func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)

}
