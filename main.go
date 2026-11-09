
package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

// Creamos una estructura donde devolvemos el tiempo y hostname
type HandsOn struct {
	Time     time.Time `json:"time"`
	Hostname string    `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Creamos un objeto con nuestra respuesta
	resp := HandsOn{
		Time:     time.Now(),
		Hostname: os.Getenv("HOSTNAME"),
	}

	// Convertimos el objeto a JSON
	jsonResp, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":9090", nil)
}
