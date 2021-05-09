package requests

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
)

type ChangeConveyorSpeedRequest struct {
	Speed float64 `json:"speed"`
}

func NewChangeConveyorSpeedRequest(r *http.Request) (*ChangeConveyorSpeedRequest, error) {
	var res ChangeConveyorSpeedRequest

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		ctx.Log(r).WithError(err).Error("failed to decode body")
		return nil, err
	}

	if res.Speed < 0 {
		return nil, errors.New("Invalid speed")
	}

	return &res, nil
}

