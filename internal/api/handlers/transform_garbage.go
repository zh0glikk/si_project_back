package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
)

func TransformGarbage(w http.ResponseWriter, r *http.Request) {
	result := ctx.Conveyor(r).Transform()

	ctx.Loader(r).TransformedGarbage = append(ctx.Loader(r).TransformedGarbage, result...)

	ctx.Log(r).Info("Garbage transformed")

	ape.Render(w, "Garbage transformed")
}
