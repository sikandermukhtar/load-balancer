
// ### balancer/weighted_rr.go
// Implements the Weighted Round Robin load balancing algorithm (placeholder).

package balancer

import "net/http"

type WeightedRR struct{}

func NewWeightedRR() *WeightedRR {
	return &WeightedRR{}
}

func (wrr *WeightedRR) Select(r *http.Request, healthyBackends []string) string {
	if len(healthyBackends) == 0 {
		return ""
	}
	// In a real implementation, use weights to select backends.
	// For now, return the first healthy backend.
	return healthyBackends[0]
}
