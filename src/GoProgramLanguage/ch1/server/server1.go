// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler1) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler echoes the Path component of the request URL r.
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//
//func main() {
//	r := gin.Default()
//	r.GET("/", func(c *gin.Context) {
//		c.String(http.StatusOK, fmt.Sprintf("URL: %q\n", c.Request.URL.Path))
//	})
//	r.Run()
//}
