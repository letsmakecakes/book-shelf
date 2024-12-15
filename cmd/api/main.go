package main

import (
	"book-shelf/api/router"
	"book-shelf/config"
	"errors"
	"log"
	"net/http"
)

//  @title          MYAPP API
//  @version        1.0
//  @description    This is a sample RESTful API with a CRUD

//  @contact.name   Adwaith Rajeev
//  @contact.url    adwaithrajeev@gmail.com

//  @license.name   MIT License
//  @license.url

// @host       localhost:8080
// @basePath   /v1
func main() {
	c := config.NewConfig()

	r := router.New()
	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  c.Server.ReadTimeout,
		WriteTimeout: c.Server.WriteTimeout,
		IdleTimeout:  c.Server.IdleTimeout,
	}

	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server startup failed")
	}
}
