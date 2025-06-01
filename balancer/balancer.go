
//  balancer/balancer.go
// Core load balancer logic, routes requests using a selected algorithm.

package balancer

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"load-balancer/health"
)

type Strategy interface {
	Select(r *http.Request, healthyBackends []string) string
}

type LoadBalancer struct {
	strategy Strategy
	backends []string
	health   *health.HealthChecker
}

func NewLoadBalancer(algorithm string, backends []string, health *health.HealthChecker) (*LoadBalancer, error) {
	var strategy Strategy
	switch algorithm {
	case "round_robin":
		strategy = NewRoundRobin()
	case "least_conn":
		strategy = NewLeastConn()
	case "random":
		strategy = NewRandom()
	case "ip_hash":
		strategy = NewIPHash()
	case "weighted_rr":
		strategy = NewWeightedRR()
	case "custom":
		strategy = NewCustomAlgo()
	default:
		return nil, errors.New("unknown algorithm")
	}
	return &LoadBalancer{strategy: strategy, backends: backends, health: health}, nil
}

func (lb *LoadBalancer) Balance(w http.ResponseWriter, r *http.Request) {
	healthyBackends := lb.getHealthyBackends()
	if len(healthyBackends) == 0 {
		http.Error(w, "No healthy backends available", http.StatusServiceUnavailable)
		return
	}
	selected := lb.strategy.Select(r, healthyBackends)
	parsedURL, err := url.Parse(selected)
	if err != nil {
		http.Error(w, "Invalid backend URL", http.StatusInternalServerError)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(parsedURL)
	proxy.ServeHTTP(w, r)
}

func (lb *LoadBalancer) getHealthyBackends() []string {
	var healthy []string
	for _, b := range lb.backends {
		if lb.health.IsHealthy(b) {
			healthy = append(healthy, b)
		}
	}
	return healthy
}