// Package p contains an HTTP Cloud Function.
package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
)

func SortAlpha(w http.ResponseWriter, r *http.Request) {
	type item struct {
		Title string
	}

	dec := json.NewDecoder(r.Body)

	// read open bracket
	t, err := dec.Token()
	if err == io.EOF {
		fmt.Fprint(w, "sort titles alphabetically :)")
		return
	} else if err != nil {
		log.Printf("json.NewDecoder: %v %v", err, t)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// while the array contains values
	for dec.More() {
		var i item
		// decode an array value (Message)
		err := dec.Decode(&i)
		if err != nil {
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		fmt.Fprintln(w, html.EscapeString(i.Title))
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Printf("json.NewDecoder: %v %v", err, t)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

func main() {
	// Determine port for HTTP service.
	log.Print("starting server...")
	http.HandleFunc("/", SortAlpha)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
