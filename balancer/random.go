// ### balancer/random.go
// Implements the Random load balancing algorithm.

package balancer

import (
	"math/rand"
	"net/http"
)

type Random struct{}

func NewRandom() *Random {
	return &Random{}
}

func (randAlgo *Random) Select(r *http.Request, healthyBackends []string) string {
	if len(healthyBackends) == 0 {
		return ""
	}
	idx := rand.Intn(len(healthyBackends))
	return healthyBackends[idx]
}