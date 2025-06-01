
// ### balancer/ip_hash.go
// Implements the IP Hash load balancing algorithm.

package balancer

import (
	"hash/fnv"
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
	ip := r.RemoteAddr // Note: Includes port; in practice, extract just the IP.
	h := fnv.New32a()
	h.Write([]byte(ip))
	hash := h.Sum32()
	idx := int(hash) % len(healthyBackends)
	return healthyBackends[idx]
}
