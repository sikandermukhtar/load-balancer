
// ### balancer/least_conn.go
// Implements the Least Connections load balancing algorithm (placeholder).
package balancer

import "net/http"

type LeastConn struct{}

func NewLeastConn() *LeastConn {
	return &LeastConn{}
}

func (lc *LeastConn) Select(r *http.Request, healthyBackends []string) string {
	if len(healthyBackends) == 0 {
		return ""
	}
	// In a real implementation, track active connections and select the backend with the fewest.
	// For now, return the first healthy backend.
	return healthyBackends[0]
}