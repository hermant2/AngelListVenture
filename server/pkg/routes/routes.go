package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/hermant2/angelventureserver/pkg/applogger"
	"github.com/hermant2/angelventureserver/pkg/routes/internal/api"
	"github.com/hermant2/angelventureserver/pkg/routes/internal/prorate"
	"github.com/hermant2/angelventureserver/pkg/usecase"
)

func Router() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger,
		api.CorsMiddlewareHandler(),
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


