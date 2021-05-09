package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"strconv"

	"github.com/si_project_back/internal/api/ctx"
	"github.com/si_project_back/internal/api/requests"
	"github.com/si_project_back/internal/garbage"
)

func AddGarbage(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewAddGarbageRequest(r)
	if err != nil {
		ape.Render(w, problems.BadRequest(err))
		ctx.Log(r).WithError(err).Error("bad request")
		return
	}

	isOk := ctx.GarbageTruck(r).AddGarbage(garbage.Garbage{
		Weight:      req.Weight,
		GarbageType: req.GarbageType,
	})

	ctx.Log(r).Info("Current Fullness/Capacity: " + strconv.Itoa(ctx.GarbageTruck(r).Fullness) + "/" + strconv.Itoa(ctx.GarbageTruck(r).Capacity))

	if !isOk {
		ape.Render(w, problems.InternalError())
		ctx.Log(r).Info("Vehicle is overweight")
		return
	}
}
