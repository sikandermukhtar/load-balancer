
// ### health/healthcheck.go
// Implements health checks for backend servers by pinging their /health endpoints.

package health

import (
	"net/http"
	"time"
)

type HealthChecker struct {
	backends map[string]bool // backend URL to health status
}

func NewHealthChecker(backends []string) *HealthChecker {
	hc := &HealthChecker{
		backends: make(map[string]bool),
	}
	for _, b := range backends {
		hc.backends[b] = false // initially assume unhealthy
	}
	return hc
}

func (hc *HealthChecker) Start() {
	for {
		for backend := range hc.backends {
			go hc.checkHealth(backend)
		}
		time.Sleep(10 * time.Second) // Check every 10 seconds
	}
}

func (hc *HealthChecker) checkHealth(backend string) {
	resp, err := http.Get(backend + "/health")
	if err != nil {
		hc.backends[backend] = false
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		hc.backends[backend] = true
	} else {
		hc.backends[backend] = false
	}
}

func (hc *HealthChecker) IsHealthy(backend string) bool {
	return hc.backends[backend]
}