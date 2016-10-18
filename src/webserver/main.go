package main

import (
  "net/http"
  "os"
  "log"
)

func main() {
        log.Println("Listening on port 8081 and serving files in: "+ os.Getenv("SNAP")+"/web")

      log.Fatal(http.ListenAndServe(":8081", http.FileServer(http.Dir(os.Getenv("SNAP")+"/web"))))
    }
