package stampede_test

import (
	"testing"
	"time"

	"github.com/aelse/stampede"
)

// For constant scaling factor and time (approx), increased cost should
// increase probability of a refresh.
func TestShouldRefreshDistributionByCost(t *testing.T) {
	expiry := 10 * time.Second
	lastRefreshCount := 0
	for cost := time.Second; cost < 10*time.Second; cost += time.Second {
		shouldRefreshCount := 0
		for i := 0; i < 5000; i++ {
			if stampede.ShouldRefresh(expiry, cost, 1) {
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

// For constant scaling factor and cost, reduced ttl (expiry) should
// increase probability of a refresh.
func TestShouldRefreshDistributionByExpiry(t *testing.T) {
	cost := 100 * time.Millisecond
	lastRefreshCount := 0
	for ttl := time.Second; ttl >= 0; ttl -= 100 * time.Millisecond {
		t.Logf("ttl %v", ttl)
		shouldRefreshCount := 0
		for i := 0; i < 5000; i++ {
			if stampede.ShouldRefresh(ttl, cost, 2) {
				shouldRefreshCount++
			}
		}
		t.Logf("current %d, last %d", shouldRefreshCount, lastRefreshCount)
		if lastRefreshCount > shouldRefreshCount {
			t.Errorf("Expected refresh count to be greater than previous round")
		}
		lastRefreshCount = shouldRefreshCount
	}
}

func BenchmarkShouldRefresh(b *testing.B) {
	expiry := 10 * time.Second
	cost := time.Second
	scaling := 1.0
	for i := 0; i < b.N; i++ {
		stampede.ShouldRefresh(expiry, cost, scaling)
	}
}
