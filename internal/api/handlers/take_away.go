package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
)

func TakeAway(w http.ResponseWriter, r *http.Request) {
	ctx.Truck(r).Dump()

	ctx.Log(r).Info("Truck dumped")

	ape.Render(w, "Truck dumped")
}

