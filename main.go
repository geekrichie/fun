package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"web/middleware"
	"web/router"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())
	router.RegisterRoute(r)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %s", err)
		} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("gracefully shutdown server fail: %v", err)
	}
	log.Println("Server exiting")
}
