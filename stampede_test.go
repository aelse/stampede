package stampede_test

import (
	"testing"
	"time"

	"github.com/aelse/stampede"
)

// For constant scaling factor and time (approx), increased cost should
// increase probability of a refresh.
func TestXFetchDistributionByCost(t *testing.T) {
	lastRefreshCount := 0
	for cost := time.Second; cost < 10*time.Second; cost += time.Second {
		shouldRefreshCount := 0
		expiry := 10 * time.Second
		for i := 0; i < 5000; i++ {
			if stampede.XFetch(expiry, cost, 1) {
				shouldRefreshCount++
			}
		}
		t.Logf("current %d, last %d", shouldRefreshCount, lastRefreshCount)
		if lastRefreshCount > shouldRefreshCount {
			t.Fatalf("Expected refresh count to be greater than previous round")
		}
		lastRefreshCount = shouldRefreshCount
	}
}

func BenchmarkXFetch(b *testing.B) {
	expiry := 10 * time.Second
	cost := time.Second
	scaling := 1.0
	for i := 0; i < b.N; i++ {
		stampede.XFetch(expiry, cost, scaling)
	}
}
