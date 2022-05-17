package lib

import (
  "fmt"
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "github.com/joho/godotenv"
  "os"
)

func Start() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  port := os.Getenv("SERVER_PORT")

  r := mux.NewRouter()
  r.HandleFunc("/auth", RedirectUser)
  fmt.Println("Server is running on port " + port)
  log.Fatal(http.ListenAndServe(":" + port, r))
}
