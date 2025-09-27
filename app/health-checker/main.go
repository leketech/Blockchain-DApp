package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// HealthStatus represents the health status of the application
type HealthStatus struct {
	Status    string  `json:"status"`
	Timestamp string  `json:"timestamp"`
	Database  string  `json:"database"`
	CPUUsage  float64 `json:"cpu_usage"`
	MemUsage  float64 `json:"mem_usage"`
}

// Metrics
var (
	healthChecksTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "health_checks_total",
			Help: "Total number of health checks",
		},
		[]string{"status"},
	)

	cpuUsageGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cpu_usage_percent",
			Help: "Current CPU usage percentage",
		},
	)

	memUsageGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "memory_usage_percent",
			Help: "Current memory usage percentage",
		},
	)
)

func init() {
	// Register metrics
	prometheus.MustRegister(healthChecksTotal)
	prometheus.MustRegister(cpuUsageGauge)
	prometheus.MustRegister(memUsageGauge)
}

func main() {
	// Set up HTTP handlers
	http.HandleFunc("/health", healthHandler)
	http.Handle("/metrics", promhttp.Handler())

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting dApp Health Checker on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Get health status
	status := getHealthStatus()

	// Update metrics
	if status.Status == "healthy" {
		healthChecksTotal.WithLabelValues("healthy").Inc()
	} else {
		healthChecksTotal.WithLabelValues("unhealthy").Inc()
	}

	cpuUsageGauge.Set(status.CPUUsage)
	memUsageGauge.Set(status.MemUsage)

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send response
	json.NewEncoder(w).Encode(status)
}

func getHealthStatus() HealthStatus {
	status := HealthStatus{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
		Database:  "unknown",
		CPUUsage:  getCPUUsage(),
		MemUsage:  getMemUsage(),
	}

	// Check database connectivity
	if checkDatabase() {
		status.Database = "connected"
	} else {
		status.Database = "disconnected"
		status.Status = "unhealthy"
	}

	return status
}

func checkDatabase() bool {
	// Get database connection details from environment
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Create connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return false
	}
	defer db.Close()

	// Ping database
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return false
	}

	return true
}

func getCPUUsage() float64 {
	// Simulate CPU usage (in a real app, you would get this from system metrics)
	return 25.5
}

func getMemUsage() float64 {
	// Simulate memory usage (in a real app, you would get this from system metrics)
	return 45.2
}