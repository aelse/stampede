// Package stampede provides features enabling probabilistic cache invalidation.
package stampede

import (
	"math"
	"math/rand"
	"time"
)

// Using a fixed seed results in the same output every run but we do not care here.
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

/*
The XFetch algorithm as described in `Optimal Probabilistic Cache Stampede Prevention`.

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

// XFetch implements the selection component of the XFetch algorithm.
// expiry: is the time.Duration until the cache value expires
// ∆ -> cost: time it takes to regenerate the cached value
// β -> scaling (1 is a reasonable default): can be increased to more aggressively avoid stampedes
func XFetch(expiry time.Duration, cost time.Duration, scaling float64) bool {
	c := float64(cost)
	d := c * scaling * math.Log(rnd.Float64())
	delta := time.Duration(int(d))
	if delta > expiry {
		return true
	}
	return false
}
