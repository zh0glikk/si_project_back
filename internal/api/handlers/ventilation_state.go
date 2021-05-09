package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
)

type ventilationStateResponse struct {
	IsWorking bool `json:"is_working"`
}

func VentilationState(w http.ResponseWriter, r *http.Request) {
	response := ventilationStateResponse{
		IsWorking: ctx.Ventilation(r).IsWorking,
	}

	ctx.Log(r).Info(response.IsWorking)

	ape.Render(w, response)
}
