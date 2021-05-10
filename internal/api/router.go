package api

import (
	"github.com/si_project_back/internal/api/ctx"
	"github.com/si_project_back/internal/api/handlers"
	"github.com/si_project_back/internal/config"
	"github.com/si_project_back/internal/conveyor"
	"github.com/si_project_back/internal/loader"
	"github.com/si_project_back/internal/vehicle"
	"github.com/si_project_back/internal/ventilation"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func Router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(cfg.Log()),
		ape.LoganMiddleware(cfg.Log()),
		ape.CtxMiddleWare(
			ctx.SetLog(cfg.Log()),
			ctx.SetVentilation(ventilation.New()),
			ctx.SetLoader(loader.NewLoader()),
			ctx.SetGarbageTruck(vehicle.NewGarbageTruck(cfg.Vehicle().Capacity)),
			ctx.SetTruck(vehicle.NewTruck(cfg.Vehicle().Capacity)),
			ctx.SetConveyor(conveyor.NewConveyor(cfg.Conveyor().Speed)),
		),
	)

	r.Route("/", func(r chi.Router) {
		r.Route("/driver", func(r chi.Router) {
			r.Post("/add-garbage", handlers.AddGarbage)
			r.Post("/take-away", handlers.TakeAway)
		})

		r.Route("/loader", func(r chi.Router) {
			r.Post("/unload-garbage-truck", handlers.UnloadGarbageTruck)
			r.Post("/load-truck", handlers.LoadTruck)
			r.Post("/load-conv", handlers.LoadConveyor)
		})

		r.Route("/operator", func(r chi.Router) {
			r.Post("/change-conveyor-state", handlers.ChangeConveyorState)
			r.Post("/change-conveyor-speed", handlers.ChangeConveyorSpeed)

			r.Post("/transform-garbage", handlers.TransformGarbage)

			r.Post("/change-ventilation-state", handlers.ChangeVentilationState)
		})

		r.Route("/sorter", func(r chi.Router) {
			r.Post("/sort-conveyor", handlers.SortConveyor)
		})
	})

	return r
}
