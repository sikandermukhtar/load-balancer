package balancer

import (
	"net/http"
	"sync"
)

// WeightedRR implements the Weighted Round-Robin load balancing algorithm
// with equal default weights (can be customized by modifying the weights map).
type WeightedRR struct {
	mu            sync.Mutex
	currentIndex  int
	currentWeight int
	weights       map[string]int
	maxWeight     int
	gcdWeight     int
}

// NewWeightedRR creates a new WeightedRR strategy instance.
func NewWeightedRR() *WeightedRR {
	return &WeightedRR{
		currentIndex:  -1,
		currentWeight: 0,
		weights:       make(map[string]int),
	}
}

// Select chooses the next backend based on weighted round robin.
// All healthy backends default to weight=1. To customize weights,
// modify wrr.weights before the first Select call.
func (wrr *WeightedRR) Select(r *http.Request, healthyBackends []string) string {
	wrr.mu.Lock()
	defer wrr.mu.Unlock()

	if len(healthyBackends) == 0 {
		return ""
	}

	// Initialize weights and metrics on first call or when backends change
	if len(wrr.weights) != len(healthyBackends) {
		wrr.weights = make(map[string]int)
		for _, b := range healthyBackends {
			wrr.weights[b] = 1 // default weight
		}
		wrr.maxWeight = computeMax(wrr.weights)
		wrr.gcdWeight = computeGCD(wrr.weights)
		wrr.currentIndex = -1
		wrr.currentWeight = wrr.maxWeight
	}

	for {
		// Round-robin advance
		wrr.currentIndex = (wrr.currentIndex + 1) % len(healthyBackends)
		if wrr.currentIndex == 0 {
			wrr.currentWeight -= wrr.gcdWeight
			if wrr.currentWeight <= 0 {
				wrr.currentWeight = wrr.maxWeight
				if wrr.currentWeight == 0 {
					return ""
				}
			}
		}

		backend := healthyBackends[wrr.currentIndex]
		if wrr.weights[backend] >= wrr.currentWeight {
			return backend
		}
	}
}

// computeMax returns the maximum weight in the map.
func computeMax(weights map[string]int) int {
	max := 0
	for _, w := range weights {
		if w > max {
			max = w
		}
	}
	return max
}

// computeGCD computes the greatest common divisor of all weights.
func computeGCD(weights map[string]int) int {
	var g int
	for _, w := range weights {
		if g == 0 {
			g = w
		} else {
			g = gcd(g, w)
		}
	}
	return g
}

// gcd calculates greatest common divisor using Euclidean algorithm.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
