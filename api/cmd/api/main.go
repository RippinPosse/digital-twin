package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"api/config"
	"api/internal/graph"
	"api/internal/graph/dataloader"
	"api/internal/hdfs"
	"api/internal/observability"
	"api/internal/server"
	"api/internal/service"
	"api/pkg/tvs"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	go func() {
		<-stop
		cancel()
	}()

	if err := godotenv.Load(); err != nil {
		err = fmt.Errorf("load env file: %w", err)
		logger.Error(err.Error())
	}

	c, err := config.New()
	if err != nil {
		err = fmt.Errorf("create config: %w", err)
		logger.Error(err.Error())

		return
	}

	tvsInteractor := tvs.New(c.TVS.Address, c.TVS.ClientID, c.TVS.ClientSecret)
	svc := service.New(tvsInteractor)
	dl := dataloader.New(&hdfs.HDFS{})

	graphHandler := graph.New(svc, dl)

	// Setup health probe routes
	routerInternal := chi.NewRouter()
	routerInternal.HandleFunc("/live", observability.LiveHandler)
	routerInternal.HandleFunc("/health", observability.HealthHandler)

	srvInternal := server.New(routerInternal, c.PortInternal)
	go srvInternal.Run(ctx)

	// Setup graphql routes
	routerExternal := chi.NewRouter()
	routerExternal.Use(dl.Middleware)
	routerExternal.Handle("/query", graphHandler)
	routerExternal.Handle("/", graph.Playground("/query"))

	srvExternal := server.New(routerExternal, c.Port)
	if err := srvExternal.Run(ctx); err != nil {
		err = fmt.Errorf("run external server: %w", err)
		logger.Error(err.Error())

		return
	}
}
