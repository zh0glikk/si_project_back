package handlers

import (
	"fmt"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
	"github.com/si_project_back/internal/api/requests"
)

func ChangeConveyorSpeed(w http.ResponseWriter, r *http.Request)  {
	res, err := requests.NewChangeConveyorSpeedRequest(r)
	if err != nil {
		ctx.Log(r).WithError(err).Error("bad request")
		ape.Render(w, problems.BadRequest(err))
		return
	}

	ctx.Conveyor(r).Speed = res.Speed

	ctx.Log(r).Info(fmt.Sprintf("speed changed to %s", res.Speed))

	ape.Render(w, "Speed changed.")
}
