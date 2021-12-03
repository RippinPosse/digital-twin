package main

import (
	"context"
	"data-processor/config"
	"data-processor/internal/kerberos"
	"data-processor/internal/observability"
	"data-processor/internal/server"
	"fmt"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net"
	"os"
	"os/signal"
	"strconv"
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

	c, err := config.New()
	if err != nil {
		err = fmt.Errorf("new config: %w", err)
		logger.Error(err.Error())

		return
	}

	address := net.JoinHostPort(c.Krb.Host, strconv.Itoa(c.Krb.Port))

	k, err := kerberos.New(address, c.Krb.Username,)

	routerInternal := chi.NewRouter()
	routerInternal.HandleFunc("/health", observability.HealthHandler)
	routerInternal.HandleFunc("/live", observability.LiveHandler)

	srvInternal := server.New(routerInternal, c.PortInternal)

	if err := srvInternal.Run(ctx); err != nil {
		err = fmt.Errorf("run internal server: %w", err)
		logger.Error(err.Error())

		return
	}
}
