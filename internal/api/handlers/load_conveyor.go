package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
)

func LoadConveyor(w http.ResponseWriter, r *http.Request) {
	garbage := ctx.Loader(r).Dump()

	ctx.Conveyor(r).Garbage = append(ctx.Conveyor(r).Garbage, garbage...)

	ape.Render(w, "Loaded to conveyor")
}

