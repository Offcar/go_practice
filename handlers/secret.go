package handlers

import (
	"fmt"
	"net/http"
)

func SecretHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		msg := "Sale pa alla qlo"
		fmt.Println(msg)
		http.Error(w, msg, 500)
		return
	}

	if r.Body != nil {
		fmt.Println("chan viene sin na?")
	}

	fmt.Println("SUPER SECRET REQUEST, QUE PASA SI LLEGA GET?")
}
