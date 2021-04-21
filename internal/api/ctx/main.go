package ctx

import (
	"context"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
)

type key int

const (
	keyLog key = iota
	keyDB
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
