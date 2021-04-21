package service

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"

	"github.com/si_project_back/internal/api"
	"github.com/si_project_back/internal/config"
)

type service struct {
	logger           *logan.Entry
	cfg 	config.Config
	ctx     context.Context
}

func NewService(ctx context.Context, cfg config.Config) *service {
	return &service{
		logger:          cfg.Log(),
		ctx: 			ctx,
		cfg: 			cfg,
	}
}

func (s *service) Run() error {
	r := api.Router(s.cfg)

	err := http.Serve(s.cfg.Listener(), r)
	if err != nil {
		return errors.Wrap(err, "server stopped with error")
	}

	return nil
}

