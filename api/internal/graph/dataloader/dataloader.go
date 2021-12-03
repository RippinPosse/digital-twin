package dataloader

import (
	"context"
	"net/http"
	"time"

	"api/internal/hdfs"
)

const loadersKey = "dataloaders"

type Loaders struct {
	UserLoader *UserLoader
}

type Dataloader struct {
	hdfs *hdfs.HDFS
}

func New(hdfs *hdfs.HDFS) *Dataloader {
	return &Dataloader{
		hdfs: hdfs,
	}
}

func (d *Dataloader) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wait := 250 * time.Microsecond

		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UserLoader: &UserLoader{
				fetch:    FetchhUsers,
				wait:     wait,
				maxBatch: 100,
			},
		})

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (d *Dataloader) For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

