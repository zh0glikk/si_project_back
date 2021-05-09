package handlers

import (
	"fmt"
	"gitlab.com/distributed_lab/ape"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
)

func ChangeVentilationState(w http.ResponseWriter, r *http.Request)  {
	ctx.Ventilation(r).IsWorking = !ctx.Ventilation(r).IsWorking

	ctx.Log(r).Info(fmt.Sprintf("State changed to %v", ctx.Ventilation(r).IsWorking))

	ape.Render(w, "State changed")
}
