// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"sort"
	"net/http"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func SortAlpha(w http.ResponseWriter, r *http.Request) {
	type item struct {
		Title string
	}


	var titles = []string{}
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
		titles = append(titles, i.Title)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Printf("json.NewDecoder: %v %v", err, t)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	sort.Strings(titles)
	for _, ti := range titles {
		fmt.Fprintln(w, html.EscapeString(ti))
	}
}
