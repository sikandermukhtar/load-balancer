package main

import (
	"log"
	"net/http"
	"load-balancer/balancer"
	"load-balancer/config"
	"load-balancer/health"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize health checker
	healthChecker := health.NewHealthChecker(cfg.Backends)

	// Initialize load balancer with selected algorithm
	lb, err := balancer.NewLoadBalancer(cfg.Algorithm, cfg.Backends, healthChecker)
	if err != nil {
		log.Fatal("Failed to create load balancer:", err)
	}

	// Start health checking in a goroutine
	go healthChecker.Start()

	// Set up HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lb.Balance(w, r)
	})

	// Start HTTP server
	log.Printf("Starting load balancer on %s", cfg.ListenAddr)
	log.Fatal(http.ListenAndServe(cfg.ListenAddr, nil))
}
