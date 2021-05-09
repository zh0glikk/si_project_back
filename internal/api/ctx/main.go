package ctx

import (
	"context"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"

	"github.com/si_project_back/internal/conveyor"
	"github.com/si_project_back/internal/db"
	"github.com/si_project_back/internal/loader"
	"github.com/si_project_back/internal/tunnel"
	"github.com/si_project_back/internal/vehicle"
	"github.com/si_project_back/internal/ventilation"
)

type key int

const (
	keyLog key = iota
	keyDB
	keyVent
	keyGarbageTruck
	keyTruck
	keyLoader
	keyTunnel
	keyConveyor
)

func SetLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, keyLog, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(keyLog).(*logan.Entry)
}

func SetDB(entry *pgdb.DB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, keyDB, entry)
	}
}

func DB(r *http.Request) *pgdb.DB {
	return r.Context().Value(keyDB).(*pgdb.DB)
}

func SetVentilation(v *ventilation.Ventilation) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, keyVent, v)
	}
}

func Ventilation(r *http.Request) *ventilation.Ventilation {
	return r.Context().Value(keyVent).(*ventilation.Ventilation)
}

func SetGarbageTruck(v *vehicle.GarbageTruck) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, keyGarbageTruck, v)
	}
}

func GarbageTruck(r *http.Request) *vehicle.GarbageTruck {
	return r.Context().Value(keyGarbageTruck).(*vehicle.GarbageTruck)
}

func SetTruck(t *vehicle.Truck) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, keyTruck, t)
	}
}

func Truck(r *http.Request) *vehicle.Truck {
	return r.Context().Value(keyTruck).(*vehicle.Truck)
}

func SetLoader(v *loader.Loader) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, keyLoader, v)
	}
}

func Loader(r *http.Request) *loader.Loader {
	return r.Context().Value(keyLoader).(*loader.Loader)
}

func SetTunnel(v *tunnel.Tunnel) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, keyTunnel, v)
	}
}

func Tunnel(r *http.Request) *tunnel.Tunnel {
	return r.Context().Value(keyTunnel).(*tunnel.Tunnel)
}

func SetConveyor(conveyor *conveyor.Conveyor) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, keyConveyor, conveyor)
	}
}

func Conveyor(r *http.Request) *conveyor.Conveyor {
	return r.Context().Value(keyConveyor).(*conveyor.Conveyor)
}

func GarbageQ(r *http.Request) *db.GarbageQ {
	return db.NewGarbageQ(r.Context().Value(keyDB).(*pgdb.DB).Clone())
}


