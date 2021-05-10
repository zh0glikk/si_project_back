package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
)

func SortConveyor(w http.ResponseWriter, r *http.Request) {
	res := ctx.Conveyor(r).Sort()

	if !res {
		ctx.Log(r).Error("Failed to sort conveyor. Conveyor is empty.")
		ape.Render(w, "Failed to sort conveyor. Conveyor is empty.")
		return
	}

	ctx.Log(r).Info("Garbage sorted")

	ape.Render(w, "Garbage sorted")
}
