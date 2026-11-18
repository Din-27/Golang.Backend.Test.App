package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"runtime": "Go stdlib",
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		testMsg := os.Getenv("TEST_MESSAGE")
		if testMsg == "" {
			testMsg = "not set"
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message":           "Deployment Test Success - Go",
			"runtime":           runtime.Version(),
			"hostname":          hostname,
			"port":              port,
			"host":              r.Host,
			"x_forwarded_proto": r.Header.Get("X-Forwarded-Proto"),
			"env_test":          testMsg,
		})
	})

	fmt.Printf("Go server listening on :%s\n", port)
	http.ListenAndServe("0.0.0.0:"+port, nil)
}
