package main

import (
	"book-shelf/api/router"
	"book-shelf/config"
	"book-shelf/util/validator"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	c := config.NewConfig()
	v := validator.New()

	var logLevel logger.LogLevel
	if c.Server.Debug {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}

	dbString := fmt.Sprintf(fmtDBString, c.DB.Host, c.DB.User, c.DB.Password, c.DB.Name, c.DB.Port)
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{Logger: logger.Default.LogMode(logLevel)})
	if err != nil {
		log.Fatal("DB connection start failure")
		return
	}

	r := router.New(db, v)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
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
