package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/imantung/empl-clocking-backend-go/internal/app"
	"github.com/imantung/empl-clocking-backend-go/internal/app/infra/di"
	"go.uber.org/multierr"
)

func main() {
	exitAppSig := make(chan os.Signal, 1)
	signal.Notify(exitAppSig, syscall.SIGTERM, syscall.SIGINT)

	var err error
	go func() {
		defer func() { exitAppSig <- syscall.SIGTERM }()
		err = multierr.Append(err, di.Invoke(app.Start)) // NOTE: start app
	}()
	<-exitAppSig

	err = multierr.Append(err, di.Invoke(app.Stop)) // NOTE: stop app (gracefully)
	if err != nil {
		log.Fatal(err)
	}
}
