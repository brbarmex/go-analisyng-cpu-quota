package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	go func(){
		for{
			go func ()  {
				os.ReadFile("/etc/")	
			}()
		}
	}()

	go func(){
		for{
			go func ()  {
				os.ReadFile("/etc/")	
			}()
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Starting HTTP server on :8080")
	http.ListenAndServe(":8080", nil)
}

