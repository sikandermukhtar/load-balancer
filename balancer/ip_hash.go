// ### balancer/ip_hash.go
// Implements the IP Hash load balancing algorithm.

package balancer

import (
	"hash/fnv"
	"net"
	"net/http"
)

type IPHash struct{}

func NewIPHash() *IPHash {
	return &IPHash{}
}

func (iph *IPHash) Select(r *http.Request, healthyBackends []string) string {
	if len(healthyBackends) == 0 {
		return ""
	}

	// Extract IP only (exclude port)
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// fallback to original if parsing fails
		ip = r.RemoteAddr
	}

	h := fnv.New32a()
	h.Write([]byte(ip))
	hash := h.Sum32()
	idx := int(hash) % len(healthyBackends)

	return healthyBackends[idx]
}
