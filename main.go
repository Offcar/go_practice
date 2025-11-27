package main

import (
	// "fmt"
	"net/http"

	/*
		Importando handlers o cualquier otro elemento 
		- usar una carpetita o haciendo algo
		- se importan nombre de la app descrita en go.mod
			- en este caso web_app/handlers
	*/
	"web_app/handlers"
)

// Entendiendo conceptos básicos de server
// a partir de server http se deja escuchando la app a una url en particular
// recordemos que leer información de un server es conceptualmente igual a leer un archivo
// en vez de bytes leídos desde un archivo, son bytes leídos dede una conexión al puerto expuesto.

// Dia 1: hacer que la app reciba un curl al localhost e imprima hola mundo

func main() {
	mux := http.NewServeMux()

	// Cargando files
	fs := http.FileServer(http.Dir("static"))
	if fs != nil {
		mux.Handle("/static/", http.StripPrefix("/static/", fs))
	}

	// Acá por probar, importaré todos los handlers
	// y su foo respectiva o controller en /handlers
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/secret", handlers.SecretHandler)

	http.ListenAndServe(":8080", mux)
}
