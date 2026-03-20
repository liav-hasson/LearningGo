package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/liav-hasson/learning-code/go/http-server/internal"
)

func main() {
	http.HandleFunc("/", internal.GetRoot)
	http.HandleFunc("/hello", internal.GetHello)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed!")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
