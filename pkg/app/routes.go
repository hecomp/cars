package app

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRoute(handler CarsHandler) *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/car/", handler.GetCar)          // GET
	mux.HandleFunc("/cars", handler.GetCars)         // GET
	mux.HandleFunc("/create", handler.CreateCar)     // POST
	mux.HandleFunc("/update", handler.UpdateCar)     // PUT
	mux.HandleFunc("/health", handler.HealthHandler) // GET
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	mux.Handle("/metrics", promhttp.Handler())

	return mux
}
