package balancer

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"time"
	"load-balancer/health"
)

type Strategy interface {
	Select(r *http.Request, healthyBackends []string) string
}

type LoadBalancer struct {
	strategy   Strategy
	backends   []string
	health     *health.HealthChecker
	proxyCache map[string]*httputil.ReverseProxy
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
    default:
        return nil, errors.New("unknown algorithm")
    }

    proxyCache := make(map[string]*httputil.ReverseProxy)
    for _, backend := range backends {
        u, err := url.Parse(backend)
        if err != nil {
            return nil, err
        }
        // Normalize path to avoid trailing slash issues
        u.Path = path.Clean(u.Path)
        proxy := httputil.NewSingleHostReverseProxy(u)
        proxy.Transport = &http.Transport{
            DialContext:           (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
            ResponseHeaderTimeout: 30 * time.Second,
        }
        proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
            log.Printf("Proxy error for %s to %s: %v", r.URL, backend, err)
            http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
        }
        proxy.Director = func(req *http.Request) {
            req.URL.Scheme = u.Scheme
            req.URL.Host = u.Host
            req.Host = u.Host // Explicitly set the Host header
            // Simplify path handling to avoid artifacts like "/."
            req.URL.Path = path.Clean(req.URL.Path)
            if req.URL.Path == "" || req.URL.Path == "." {
                req.URL.Path = "/"
            }
            // Log the full request for debugging
            log.Printf("Proxying request: %s %s Headers: %+v", req.Method, req.URL.String(), req.Header)
        }
        proxyCache[backend] = proxy
    }

    return &LoadBalancer{
        strategy:   strategy,
        backends:   backends,
        health:     health,
        proxyCache: proxyCache,
    }, nil
}

func (lb *LoadBalancer) Balance(w http.ResponseWriter, r *http.Request) {
	healthyBackends := lb.getHealthyBackends()
	if len(healthyBackends) == 0 {
		log.Println("No healthy backends available")
		http.Error(w, "No healthy backends available", http.StatusServiceUnavailable)
		return
	}
	selected := lb.strategy.Select(r, healthyBackends)
	proxy, exists := lb.proxyCache[selected]
	if !exists {
		log.Printf("No proxy found for backend %s", selected)
		http.Error(w, "Invalid backend", http.StatusInternalServerError)
		return
	}
	log.Printf("Forwarding request %s to %s", r.URL, selected)
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