package handlers

import (
	"fmt"
	"net/http"
	"html/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------ INDEX")


	fmt.Println("-- template")
	t, err := template.ParseFiles("./templates/index.html")

	if err != nil {
		msg := "Error buscando template index"
		fmt.Println(msg)
		http.Error(w, msg , 500)
	}

	// arreglo opcional instanciado y definido acá
	// data := struct{types de los datos}{datos}

	// input -> (wr io.Writer, data any)
	err = t.Execute(w, nil)
	if err != nil {
		msg := "Error renderizando template index"
		fmt.Println(msg)
		http.Error(w, msg , 500)
	}
}
