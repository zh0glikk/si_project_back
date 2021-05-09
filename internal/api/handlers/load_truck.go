package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"net/http"
	"strconv"

	"github.com/si_project_back/internal/api/ctx"
)

func LoadTruck(w http.ResponseWriter, r *http.Request) {
	garbage := ctx.Loader(r).DumpTransformedGarbage()

	for _, g := range garbage {
		isOk := ctx.Truck(r).AddGarbage(g)
		if !isOk {
			ctx.Log(r).Info("Vehicle is overweight")
			ape.Render(w, "Vehicle is overweight")
			return
		}
		ctx.Log(r).Info("Current Fullness/Capacity: " + strconv.Itoa(ctx.Truck(r).Fullness) + "/" + strconv.Itoa(ctx.Truck(r).Capacity))
	}

	ctx.Log(r).Info("Truck loaded")

	ape.Render(w, "Truck loaded")
}

