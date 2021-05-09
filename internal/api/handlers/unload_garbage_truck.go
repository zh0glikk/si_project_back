package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"net/http"
	"strconv"

	"github.com/si_project_back/internal/api/ctx"
)

func UnloadGarbageTruck(w http.ResponseWriter, r *http.Request) {
	if ctx.GarbageTruck(r).Fullness == 0 {
		ctx.Log(r).Info("GarbageTruck is empty")
		ape.Render(w, "GarbageTruck is empty")
		return
	}

	res := ctx.GarbageTruck(r).Dump()

	ctx.Loader(r).CollectedGarbage = append(ctx.Loader(r).CollectedGarbage, res...)

	ctx.Log(r).Info("Truck fullness: " + strconv.Itoa(ctx.GarbageTruck(r).Fullness))
	ctx.Log(r).Info("Vehicle was dumped")

	ape.Render(w, "Vehicle was dumped")
}
