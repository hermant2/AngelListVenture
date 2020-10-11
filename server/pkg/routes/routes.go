package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/hermant2/angelventureserver/pkg/applogger"
	"github.com/hermant2/angelventureserver/pkg/routes/internal/prorate"
	"github.com/hermant2/angelventureserver/pkg/usecase"
)

func Router() *chi.Mux {
	router := chi.NewRouter()

	cors := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: false,
		MaxAge:           300})
	router.Use(middleware.Logger,
		cors,
		middleware.Recoverer,
		render.SetContentType(render.ContentTypeJSON),
		middleware.Compress(5, "application/serializer"))
	handleV1Routes(router)

	return router
}

func handleV1Routes(router *chi.Mux) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/prorate", prorateRouter())
	})
}

func prorateRouter() *chi.Mux {
	controller := prorate.NewController(usecase.NewProrateService(), applogger.Instance())
	router := chi.NewRouter()
	router.Post("/", controller.CalculateAllocation)
	return router
}
