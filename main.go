package main

import (
	"github.com/mhsbz/xiaohan/internal/router"
	"net/http"
)

func main() {
	r := router.NewRouter()
	r.RegisterRoutes()
	if http.ListenAndServe(":8080", r.SwaggerAPI.Serve(nil)) != nil {
		panic("Failed to start server")
	}
}
