// ### health/healthcheck.go
// Implements health checks for backend servers by pinging their /health endpoints.

package health

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type HealthChecker struct {
	backends map[string]bool
	mu       sync.RWMutex
}

func NewHealthChecker(backends []string) *HealthChecker {
	hc := &HealthChecker{
		backends: make(map[string]bool),
	}
	for _, b := range backends {
		hc.backends[b] = false
	}
	return hc
}

func (hc *HealthChecker) Start() {
	for {
		for backend := range hc.backends {
			go hc.checkHealth(backend)
		}
		time.Sleep(100 * time.Second)
	}
}
func (hc *HealthChecker) checkHealth(backend string) {
    resp, err := http.Get(backend + "/health")
    hc.mu.Lock()
    defer hc.mu.Unlock()
    if err != nil {
        log.Printf("Health check failed for %s: %v", backend, err)
        hc.backends[backend] = false
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        log.Printf("Backend %s is healthy", backend)
        hc.backends[backend] = true
    } else {
        log.Printf("Backend %s returned status %d", backend, resp.StatusCode)
        hc.backends[backend] = false
    }
}
func (hc *HealthChecker) IsHealthy(backend string) bool {
	hc.mu.RLock()
	defer hc.mu.RUnlock()
	return hc.backends[backend]
}
