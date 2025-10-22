package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate("index.tmpl", w, nil)
	})
	http.HandleFunc("/private", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate("private.tmpl", w, nil)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate("login.tmpl", w, nil)
	})
	http.ListenAndServe("0.0.0.0:4321", nil)
}

func renderTemplate(name string, w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles(
		"layouts/default-layout.tmpl",
		fmt.Sprintf("templates/%s", name),
	)
	if err != nil {
		w.WriteHeader(500)
		slog.Error("Error parsing template", "err", err)
		return
	}
	err = t.ExecuteTemplate(w, name, data)
	if err != nil {
		w.WriteHeader(500)
		slog.Error("Error executing template", "err", err)
		return
	}
}
