package main

import (
	"fmt"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"net/http"
	"time"
)

func main(){
	fmt.Printf("This Is Create Tag Api")
	router := routers.NewRouter()
	s := &http.Server{
		Addr					: ":8080",
		Handler					: router,
		ReadHeaderTimeout		: 10 * time.Second,
		WriteTimeout			: 10 * time.Second,
		MaxHeaderBytes			: 1 << 20,
	}

	s.ListenAndServe()
}
