package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
)

func TransformGarbage(w http.ResponseWriter, r *http.Request) {
	if !ctx.Ventilation(r).IsWorking {
		ctx.Log(r).Error("Ventilation is disabled.")
		ape.Render(w, "Ventilation is disabled.")
		return
	}

	if !ctx.Conveyor(r).IsEnabled {
		ctx.Log(r).Error("Conveyor is disabled.")
		ape.Render(w, "Conveyor is disabled.")
		return
	}

	result := ctx.Conveyor(r).Transform()
	if result == nil {
		ctx.Log(r).Error("Can't transform garbage. Conveyor is empty or is not sorted.")
		ape.Render(w, "Can't transform garbage. Conveyor is empty or is not sorted.")
		return
	}

	ctx.Loader(r).TransformedGarbage = append(ctx.Loader(r).TransformedGarbage, result...)

	ctx.Log(r).Info("Garbage transformed")

	ape.Render(w, "Garbage transformed")
}
