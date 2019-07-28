package main

import (
	"net/http"
  "log"
  "encoding/json"
)

func main() {
  fs := http.FileServer(http.Dir("ui/build"))
  http.Handle("/", fs)

  http.HandleFunc("/api", api)

  log.Println("Listening on  localhost:8080")
  http.ListenAndServe(":8080", nil)
}

type ApiResponse struct {
  Name    string
  Version float32
  Description string
}

func api(w http.ResponseWriter, r *http.Request) {
  response := ApiResponse{"go-react api", 1.0, "Welcome to the go-react api!"}

  js, err := json.Marshal(response)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}