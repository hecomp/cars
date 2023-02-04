package main

import (
	"context"
	"errors"
	"flag"
	"github.com/hecomp/cars/docs"
	"github.com/hecomp/cars/pkg/app"
	"github.com/hecomp/cars/pkg/repository"
	"github.com/hecomp/cars/pkg/services"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	version = "1.0.0"
)

// @title			GetCars CarsService
// @description	This is a Goland server that manages cars.
// @version		1.0.0
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @contact.name	cars
// @host			localhost:9000
// @BasePath		/
func main() {

	var httpAddr = flag.String("http.addr", "localhost:9000", "Address for HTTP (JSON) server")

	flag.Parse()

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Cars API"
	docs.SwaggerInfo.Description = "This is a Cars server."
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = *httpAddr
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	logger := log.New(os.Stdout, "cars ", log.LstdFlags)

	r := repository.NewRepository()
	s := services.NewCarsService(r)
	h := app.NewHandler(logger, s)
	route := app.NewRoute(h)

	ctx := context.Background()
	srv := &http.Server{
		Handler:      route,
		Addr:         *httpAddr,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, httpAddr, l.Addr().String())
			return ctx
		},
	}

	// start the server
	go func() {
		logger.Println("Starting server on port 9000")

		err := srv.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			logger.Printf("server closed")
		} else if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)

}
