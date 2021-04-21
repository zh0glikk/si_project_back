package api

import (
	"github.com/si_project_back/internal/api/ctx"
	"github.com/si_project_back/internal/config"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func Router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(cfg.Log()),
		ape.LoganMiddleware(cfg.Log()),
		ape.CtxMiddleWare(
			ctx.SetDB(cfg.DB()),
			ctx.SetLog(cfg.Log()),
		),
	)

	r.Route("/", func(r chi.Router) {

	})

	return r
}
