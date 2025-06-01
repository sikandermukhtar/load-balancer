// ### balancer/round_robin.go
// Implements the Round Robin load balancing algorithm.

package balancer

import (
	"sync"
	"net/http"	
)
	

type RoundRobin struct {
	index int
	mu    sync.Mutex
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{index: -1}
}

func (rr *RoundRobin) Select(r *http.Request, healthyBackends []string) string {
	if len(healthyBackends) == 0 {
		return ""
	}
	rr.mu.Lock()
	defer rr.mu.Unlock()
	rr.index = (rr.index + 1) % len(healthyBackends)
	return healthyBackends[rr.index]
}
