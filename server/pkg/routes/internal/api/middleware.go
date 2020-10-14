package api

import (
	"github.com/go-chi/cors"
	"github.com/hermant2/angelventureserver/pkg/env"
	"net/http"
	"os"
)

func CorsMiddlewareHandler() func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv(env.AllowedOrigin)},
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: false,
		MaxAge:           300})
}
