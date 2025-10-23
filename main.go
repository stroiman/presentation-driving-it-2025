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
	renderNamedTemplate(name, name, w, data)
}

func renderNamedTemplate(fileName, templateName string, w http.ResponseWriter, data any) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("Error parsing template", "err", err)
			renderErrPage(w)
		}
	}()
	w.Header().Set("Content-Type", "text/html")
	t := template.Must(template.ParseGlob("layouts/*.tmpl"))
	t = template.Must(t.ParseFiles(
		fmt.Sprintf("templates/%s", fileName),
	))
	if err := t.ExecuteTemplate(w, templateName, data); err != nil {
		slog.Error("Error executing template", "err", err)
		renderErrPage(w)
	}
}

func renderErrPage(w http.ResponseWriter) {
	w.WriteHeader(500)
	fmt.Fprint(w, `<!doctype html><head><html></head><body><h1>Oh, dear!</h1></body></html>`)
}
