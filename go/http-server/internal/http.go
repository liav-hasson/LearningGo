package internal

import (
	"fmt"
	"io"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a \"/\" request")
	io.WriteString(w, "Welcome to my website!")
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a \"/hello\" request")
	io.WriteString(w, "Hello, World!")
}
