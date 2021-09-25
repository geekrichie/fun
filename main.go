package main

import (
	"net/http"
	"time"
	"web/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.RegisterRoute(r)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
