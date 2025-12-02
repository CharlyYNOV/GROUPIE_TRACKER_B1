package internals

import (
	"fmt"
	"log"
	"net/http"
)

func SearchBar(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Error while parsing search bar input : %v", err)
		http.Error(w, "Error while parsing search bar input", http.StatusInternalServerError)
	}

	name := r.PostFormValue("name")
	fmt.Fprintf(w, "Hello, %s!", name)

}
