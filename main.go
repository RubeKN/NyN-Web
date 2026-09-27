package main

import (
	"fmt"
	"net/http"
)

func serveOpinion(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/opinion" || r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al parsear", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!DOCTYPE html>
						<html lang="es">
						<head>
							<meta charset="UTF-8">
							<meta name="viewport" content="width=device-width, initial-scale=1.0">
							<link rel="stylesheet" href="/css/styles.css">
							<title>Gracias por tu Opinion!</title>
						</head>
						<body>
							<h1>Gracias por tu opinion, %s!</h1>
							<a href="/">Volver</a>
						</body>
						</html>`,
		name)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "./index.html")
}

func main() {
	staticDir := "./static"
	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fileServer)

	http.HandleFunc("/opinion", serveOpinion)

	port := ":8080"
	fmt.Printf("Server running localhost%s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error %s", err)
	}
}
