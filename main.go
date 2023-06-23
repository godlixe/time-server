package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func getTIme(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"server_time":"%v"}`, time.Now().Format("2006-01-02 15:04:05"))))
}

func main() {
	http.HandleFunc("/api/time-server", getTIme)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
