// Package stampede provides features enabling probabilistic cache invalidation.
package stampede

import (
	"math"
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

/*
The stampede package provides `ShouldRefresh` implementing probabilistic cache refresh.
This relies on an implementation of the XFeteh algorithm.

XFetch is described in `Optimal Probabilistic Cache Stampede Prevention`.

function XFetch(key, ttl; β = 1)
	value, ∆, expiry ← CacheRead(key)
	if !value or Time() − ∆β log(rand()) ≥ expiry then
		start ← Time()
		value ← RecomputeValue()
		∆ ← Time() – start
		CacheWrite(key, (value, ∆), ttl)
	end
	return value
end
*/

// ShouldRefresh implements the selection component of the XFetch algorithm.
// expiry: is the time.Duration until the cache value expires. Should always be positive
// ∆ -> cost: time it takes to regenerate the cached value
// β -> scaling (1 is a reasonable default): can be increased to more aggressively avoid stampedes
func ShouldRefresh(ttl time.Duration, delta time.Duration, beta float64) bool {
	return time.Duration(-1*float64(delta)*beta*math.Log(rnd.Float64())) >= ttl
}
