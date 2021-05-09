package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
)

func SortConveyor(w http.ResponseWriter, r *http.Request) {
	ctx.Conveyor(r).Sort()

	ape.Render(w, "Conveyor sorted")
}
