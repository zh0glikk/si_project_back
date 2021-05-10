package handlers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/ape"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
	"github.com/si_project_back/internal/garbage"
)

func LoadConveyor(w http.ResponseWriter, r *http.Request) {
	if len(ctx.Loader(r).CollectedGarbage) == 0 {
		ctx.Log(r).Error("Nothing to load to conveyor.")
		ape.Render(w, "Nothing to load to conveyor.")
		return
	}

	garbage := ctx.Loader(r).Dump()

	ctx.Log(r).Info(fmt.Sprintf("Pushed garbage pack to conveyor with weight %v", calcWeight(garbage)))

	ctx.Conveyor(r).Garbage = append(ctx.Conveyor(r).Garbage, garbage...)

	ctx.Log(r).Info("Conveyor loaded")

	ape.Render(w, "Loaded to conveyor")
}

func calcWeight(garbage []garbage.Garbage) int {
	var res int

	for _, g := range garbage {
		res += g.Weight
		logrus.Info(g)
	}

	return res
}

