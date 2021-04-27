package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", ServerRespose)
	fmt.Print("[*] Servidor corriendo en el puerto 8080")
    http.ListenAndServe(":8080", nil)
}

func ServerRespose(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hola desde Go!!!")
}