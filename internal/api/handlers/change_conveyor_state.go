package handlers

import (
	"fmt"
	"gitlab.com/distributed_lab/ape"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
)

func ChangeConveyorState(w http.ResponseWriter, r *http.Request)  {
	ctx.Conveyor(r).IsEnabled = !ctx.Conveyor(r).IsEnabled

	ctx.Log(r).Info(fmt.Sprintf("State changed to %v", ctx.Conveyor(r).IsEnabled))

	ape.Render(w, "State changed")
}

