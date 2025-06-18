package balancer

import (
	"net/http"
	"sync"
)

// LeastConn implements the Least Connections load balancing algorithm.
type LeastConn struct {
	mu              sync.Mutex
	activeConnCount map[string]int
}

// NewLeastConn creates a new LeastConn strategy instance.
func NewLeastConn() *LeastConn {
	return &LeastConn{
		activeConnCount: make(map[string]int),
	}
}

// Select returns the backend with the fewest active connections among healthyBackends.
// It also increments the active connection count for the selected backend.
func (lc *LeastConn) Select(r *http.Request, healthyBackends []string) string {
	if len(healthyBackends) == 0 {
		return ""
	}

	lc.mu.Lock()
	defer lc.mu.Unlock()

	// Ensure all healthy backends are tracked
	for _, b := range healthyBackends {
		if _, exists := lc.activeConnCount[b]; !exists {
			lc.activeConnCount[b] = 0
		}
	}

	// Pick the backend with the minimum active connections
	selected := healthyBackends[0]
	minConns := lc.activeConnCount[selected]
	for _, b := range healthyBackends[1:] {
		if lc.activeConnCount[b] < minConns {
			selected = b
			minConns = lc.activeConnCount[b]
		}
	}

	// Mark one more active connection
	lc.activeConnCount[selected]++
	return selected
}

// Done should be called when a request to a backend is completed.
// It decrements the active connection count for that backend.
func (lc *LeastConn) Done(backend string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if count, exists := lc.activeConnCount[backend]; exists && count > 0 {
		lc.activeConnCount[backend]--
	}
}
