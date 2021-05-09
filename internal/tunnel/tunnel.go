package tunnel

import (
	"github.com/si_project_back/internal/garbage"
)

type Tunnel struct {
	CollectedGarbage []garbage.Garbage
}

func NewTunnel() *Tunnel {
	return &Tunnel{}
}


