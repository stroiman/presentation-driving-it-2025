package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(
		"0.0.0.0:4321",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w,
				`<!doctype html><html lang="en-US"><head>
<meta charset="utf-8" />
<meta name="viewport" content="width=device-width" />
</head><body><h1>Hello, world!</h1></body></html>`)
		}),
	)
}
