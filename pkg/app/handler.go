package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hecomp/cars/internal/constants"
	"github.com/hecomp/cars/internal/models"
	"github.com/hecomp/cars/internal/telemetry/metrics"
	"github.com/hecomp/cars/pkg/services"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	ErrEmpty       = errors.New("empty id")
	ErrCarBody     = errors.New("car %s is invalid")
	ErrCreateCar   = errors.New("error creating car")
	ErrUpdateCar   = errors.New("error updating car")
	ErrNoData      = errors.New("no data")
	ErrCarNotFound = errors.New("car not found")

	CarCreatedSuccess = fmt.Sprintf("car created successfully!")
	CarUpdatedSuccess = fmt.Sprintf("car updated successfully!")
)

// CarsHandler defines all the handlers the CarsService needs.
type CarsHandler interface {
	GetCar(w http.ResponseWriter, r *http.Request)
	GetCars(w http.ResponseWriter, r *http.Request)
	CreateCar(w http.ResponseWriter, r *http.Request)
	UpdateCar(w http.ResponseWriter, r *http.Request)
	HealthHandler(w http.ResponseWriter, r *http.Request)
}

type carsHandler struct {
	services services.CarsService
	logger   *log.Logger
}

func NewHandler(logger *log.Logger, svc services.CarsService) CarsHandler {
	return &carsHandler{services: svc, logger: logger}
}

// GetCar godoc
//
//	@Summary	Get car
//	@Schemes
//	@Description	Reads a single car and returns it.
//	@Tags			read
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Car ID"
//	@Success		200	{object}	constants.UserResponse
//	@Failure		400	{object}	constants.ErrorResponse
//	@Failure		404	{object}	constants.ErrorResponse
//	@Router			/car/{id} [get]
func (c *carsHandler) GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "Application-Json")

	endpoint := "/car"
	start := time.Now()
	id := strings.TrimPrefix(r.URL.Path, "/car/")
	if id == "" {
		metrics.BadRequestCount.WithLabelValues(endpoint, id).Inc()
		c.logger.Println(ErrEmpty)
		w.WriteHeader(http.StatusBadRequest)
		response := constants.ErrorResponse{
			Err: ErrEmpty.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	car, uErr := c.services.GetCar(id)
	if uErr != nil {
		metrics.NotFoundCount.WithLabelValues(endpoint, id).Inc()
		c.logger.Println(ErrCarNotFound)
		w.WriteHeader(http.StatusNotFound)
		response := constants.ErrorResponse{
			Err: uErr.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	metrics.RequestDuration.WithLabelValues(endpoint, strconv.Itoa(http.StatusOK), car.Id).
		Observe(time.Since(start).Seconds())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&constants.UserResponse{
		Data: car,
	})
}

// GetCars godoc
//
//	@Summary	GetCar all cars
//	@Schemes
//	@Description	Reads and returns all the cars.
//	@Tags			read
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	constants.UserResponse
//	@Failure		404	{object}	constants.ErrorResponse
//	@Router			/cars [get]
func (c *carsHandler) GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "Application-Json")

	endpoint := "/cars"
	start := time.Now()
	cars := c.services.GetCars()
	if len(cars) == 0 {
		metrics.NotFoundCount.WithLabelValues(endpoint, "").Inc()
		c.logger.Println(ErrNoData)
		w.WriteHeader(http.StatusNotFound)
		response := constants.ErrorResponse{
			Message: ErrNoData.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	metrics.RequestDuration.WithLabelValues(endpoint, strconv.Itoa(http.StatusOK), "").
		Observe(time.Since(start).Seconds())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&constants.UserResponse{
		Data: cars,
	})
}

// CreateCar godoc
//
//	@Summary	Creates car
//	@Schemes
//	@Description	Creates a new car.
//	@Tags			write
//	@Accept			json
//	@Produce		json
//	@Param			car	body		models.Car	true	"New car"
//	@Success		201	{object}	constants.UserResponse
//	@Failure		400	{object}	constants.ErrorResponse
//	@Failure		500	{object}	constants.ErrorResponse
//	@Router			/create [post]
func (c *carsHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "Application-Json")

	endpoint := "/create"
	start := time.Now()
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		metrics.BadRequestCount.WithLabelValues(endpoint, "").Inc()
		c.logger.Println(ErrCarBody)
		w.WriteHeader(http.StatusBadRequest)
		response := constants.ErrorResponse{
			Err: ErrCarBody.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	var car models.Car
	err = json.Unmarshal(bytes, &car)
	if err != nil {
		metrics.UnmarshalFailCount.WithLabelValues(endpoint, car.Id).Inc()
		c.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response := constants.ErrorResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if _, err = c.services.Create(&car); err != nil {
		metrics.CreateFailCount.WithLabelValues(endpoint, car.Id).Inc()
		c.logger.Println(ErrCreateCar)
		w.WriteHeader(http.StatusInternalServerError)
		response := constants.ErrorResponse{
			Message: ErrCreateCar.Error(),
			Err:     err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	metrics.RequestDuration.WithLabelValues(endpoint, strconv.Itoa(http.StatusOK), car.Id).
		Observe(time.Since(start).Seconds())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&constants.UserResponse{
		Message: CarCreatedSuccess,
		Data:    car,
	})
}

// UpdateCar godoc
//
//	@Summary	Update car
//	@Schemes
//	@Description	Updates a new car.
//	@Tags			write
//	@Accept			json
//	@Produce		json
//	@Param			car	body		models.Car	true	"New car"
//	@Success		200	{object}	constants.UserResponse
//	@Failure		400	{object}	constants.ErrorResponse
//	@Failure		500	{object}	constants.ErrorResponse
//	@Router			/update [put]
func (c *carsHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "Application-Json")

	endpoint := "/update"
	start := time.Now()
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		metrics.BadRequestCount.WithLabelValues(endpoint, "").Inc()
		c.logger.Println(ErrCarBody)
		w.WriteHeader(http.StatusBadRequest)
		response := constants.ErrorResponse{
			Err: ErrCarBody.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	var car models.Car
	err = json.Unmarshal(bytes, &car)
	if err != nil {
		metrics.UnmarshalFailCount.WithLabelValues(endpoint, car.Id).Inc()
		c.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response := constants.ErrorResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if _, err = c.services.Update(&car); err != nil {
		metrics.UpdateFailCount.WithLabelValues(endpoint, car.Id).Inc()
		c.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response := constants.ErrorResponse{
			Message: ErrUpdateCar.Error(),
			Err:     err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	metrics.RequestDuration.WithLabelValues(endpoint, strconv.Itoa(http.StatusOK), car.Id).
		Observe(time.Since(start).Seconds())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&constants.UserResponse{
		Message: CarUpdatedSuccess,
		Data:    car,
	})
}

// HealthHandler check liveness check
//
//	@summary		The liveness endpoint determines the LIVE status of the service
//	@description	This endpoint will return a status to determine if the service is live or requires a restart
//	@tags			Health Check
//	@id				liveliness
//	@accept			json
//	@produce		json
//	@success		200	{object}	models.HealthResponse
//	@router			/health [get]
func (c *carsHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "Application-Json")
	c.logger.Println("Checking application health")
	w.WriteHeader(http.StatusOK)
	response := &models.HealthResponse{
		Status: "UP",
	}
	json.NewEncoder(w).Encode(response)
}
