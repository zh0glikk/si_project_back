package requests

import (
	"encoding/json"
	"net/http"

	"github.com/si_project_back/internal/api/ctx"
	"github.com/si_project_back/internal/garbage"
)

type AddGarbageRequest struct {
	Weight 		int					`json:"weight"`
	GarbageType garbage.GarbageType	`json:"garbage_type"`
}

func NewAddGarbageRequest(r *http.Request) (*AddGarbageRequest, error) {
	var res AddGarbageRequest

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		ctx.Log(r).WithError(err).Error("failed to decode body")
		return nil, err
	}

	return &res, nil
}
