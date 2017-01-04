package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dring1/jwt-oauth/config"
	"github.com/dring1/jwt-oauth/middleware"
	"github.com/dring1/jwt-oauth/routes"
	"github.com/dring1/jwt-oauth/services"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	svcs, err := services.New(c)
	if err != nil {
		log.Fatalln(err)
	}
	middlewares, err := middleware.New(svcs)
	if err != nil {
		log.Fatalln(err)
	}
	// Init router
	rs, err := routes.NewRoutes(&routes.Config{
		Services:     svcs,
		Middlewares:  middlewares,
		ClientID:     c.GitHubClientID,
		ClientSecret: c.GitHubClientSecret,
	})
	if err != nil {
		log.Fatalln(err)
	}
	router, err := routes.NewRouter(rs)
	if err != nil {
		log.Fatalln(err)
	}
	// Apply middlewares
	globalMiddlewares := []middleware.Middleware{
		middleware.NewApacheLoggingHandler(c.LoggingEndpoint),
	}
	globalMiddlewares = append(globalMiddlewares, middleware.DefaultMiddleWare()...)

	log.Printf("Serving on port :%d", c.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", c.Port), middleware.Handlers(router, globalMiddlewares...))
	log.Fatal(err)
}
