// example-web-service
// author: Patrick Easters

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()

	if err != nil {
		log.Println("err:", err)
		http.Error(w, "Error getting hostname... weird.", 500)
	}

	if svc := os.Getenv("APP_SVC"); len(svc) > 0 {
		fmt.Fprintf(w, "Hi, I'm the %s service!\n\nHostname: %s\n", svc, host)
	} else {
		fmt.Fprintf(w, "Hi, nice to serve you today.\n\nHostname: %s\n", host)
	}
	if node := os.Getenv("NODE_NAME"); len(node) > 0 {
		fmt.Fprintf(w, "Node: %s\n", node)
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
