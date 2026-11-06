package main
##prueba 
import (
	"fmt"
	"time"
	"os"
        "net/http"
)
type HansOn struct {
    Time    time.Time `json:"time"`
    Hostname string   `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	resp := HandsOn{
	   Time:      time.Now(),
           Hostname: os.Getent("HOSTNAME"),
	}
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

