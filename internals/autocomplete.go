// A FAIRE APRES PAGE DE CHAQUE ARTISTES.

package internals

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func Autocomplete() {

	http.HandleFunc("/suggestions", func(w http.ResponseWriter, r *http.Request) {
		searchValue := strings.TrimSpace(r.URL.Query().Get("q"))
		if searchValue == "" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		data, err := ioutil.ReadFile("suggetions.json")
		if err != nil {
			http.Error(w, "Failed to read suggestions", http.StatusInternalServerError)
			return
		}

		var suggestions []string
		if err := json.Unmarshal(data, &suggestions); err != nil {
			http.Error(w, "Failed to parse suggestions", http.StatusInternalServerError)
			return
		}

		var filtered []string
		lowerSearch := strings.ToLower(searchValue)
		for _, temp := range suggestions {
			if strings.HasPrefix(strings.ToLower(temp), lowerSearch) {
				filtered = append(filtered, temp)
			}
		}

		resp, err := json.Marshal(filtered)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	})

}
