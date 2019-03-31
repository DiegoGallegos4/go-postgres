package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	l "github.com/treastech/logger"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

type Gym struct {
	Name string `validate:"required"`
}

var validate *validator.Validate

func main() {
	router := chi.NewRouter()
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	logger, _ := config.Build()

	validate = validator.New()

	defer logger.Sync()
	router.Use(middleware.RequestID)
	router.Use(l.Logger(logger))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var gym Gym
		json.NewDecoder(r.Body).Decode(&gym)
		if err := validate.Struct(gym); err != nil {
			logger.Warn("Error", zap.Error(err))
		}
	})
	logger.Info("Running server in port 5000")
	if err := http.ListenAndServe(":5000", router); err != nil {
		fmt.Print(err)
	}
}
