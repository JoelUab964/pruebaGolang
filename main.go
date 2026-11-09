package main
import (
	"fmt"
	"time"
	"os"
    "net/http"
)
//creamos una estructura donde devolvemos el tiempo y hostname 
type HansOn struct {
    Time     time.Time `json:"time"`
    Hostname string   `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//creamos un objeto con nuestra respuesta 
	resp := HandsOn{
	    Time: time.Now(),
        Hostname: os.Getent("HOSTNAME"),
	}
	//usamos una libreria para volvera string
	jsonResp, err := json.Marshal(&resp)
	if err != nil {
	    w.Write([]byte("Error"))
	    return
	}

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
   w.Write([]byte(resp))
}
func main() {
        http.HandleFunc("/", ServeHTTP)
        http.ListenAndServe(":9090", nil)
}

