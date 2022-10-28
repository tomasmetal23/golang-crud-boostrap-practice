package main

import (
	"database/sql"
	"html/template"

	//"log"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Funcion para la conexion y autenticacion con la Base de datos
func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Password := "nEis4bGZe"
	Nombre := "crud"

	conexion, err := sql.Open(Driver, Usuario+":"+Password+"@tcp(10.89.2.2)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion

}

// plantilla
var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)

	http.HandleFunc("/borrar", Borrar)

	fmt.Println("Running Service...")

	http.ListenAndServe(":8080", nil)
}

// Funcion borrar
func Borrar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	fmt.Println(idEmpleado)

	conexionEstablecida := conexionBD()
	//Instruccion SQL para borrar
	borrarRegistros, err := conexionEstablecida.Prepare("DELETE FROM empleados WHERE id=?")

	if err != nil {
		panic(err.Error())
	}
	insertarRegistros.Exec(idEmpleado)

	http.Redirect(w, r, "/", 301)

}

type Empleado struct {
	Id     int
	Nombre string
	Correo string
}

func Inicio(w http.ResponseWriter, r *http.Request) {

	conexionEstablecida := conexionBD()

	registros, err := conexionEstablecida.Query("SELECT * FROM empleados")

	if err != nil {
		panic(err.Error())
	}
	empleado := Empleado{}
	arregloEmpleado := []Empleado{}

	for registros.Next() {
		var id int
		var nombre, correo string
		err = registros.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

		arregloEmpleado = append(arregloEmpleado, empleado)

	}
	//para mostrar el arreglo empleado en la consola ##fmt.Println(arregloEmpleado)

	//fmt.Fprintf(w, "Hello Saiyans!!!")
	plantillas.ExecuteTemplate(w, "inicio", arregloEmpleado)

}

func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)

}
func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")

		conexionEstablecida := conexionBD()

		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre,correo) VALUES('?','?') ")

		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(nombre, correo)

		http.Redirect(w, r, "/", 301)

	}

}
