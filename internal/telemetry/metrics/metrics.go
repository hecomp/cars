package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	BadRequestCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_bad_request_count",
		Help: "The total number of bad request",
	}, []string{"endpoint", "car"})
	NotFoundCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_not_found_request_count",
		Help: "The total number of not found request",
	}, []string{"endpoint", "car"})
	UnmarshalFailCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_unmarshal_fail_request_count",
		Help: "The total number of unmarshal fail request",
	}, []string{"endpoint", "car"})
	CreateFailCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_create_create_request_count",
		Help: "The total number of unmarshal fail request",
	}, []string{"endpoint", "car"})
	UpdateFailCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_update_fail_request_count",
		Help: "The total number of unmarshal fail request",
	}, []string{"endpoint", "car"})
	RequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	}, []string{"endpoint", "status", "car"})
)
