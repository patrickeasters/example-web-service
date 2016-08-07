// example-web-service
// author: Patrick Easters

package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
    svc := os.Getenv("APP_SVC")
    host, err := os.Hostname()

    if(err != nil){
      log.Println("err:", err)
      http.Error(w, "Error getting hostname... weird.", 500)
    }

    if(len(svc) > 0){
      fmt.Fprintf(w, "Hi, I'm the %s service!\n\nHostname: %s", svc, host)
    }else{
      fmt.Fprintf(w, "Hi, nice to serve you today.\n\nHostname: %s", host)
    }

}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
