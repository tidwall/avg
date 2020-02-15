package avg

import (
	"sync"
	"time"
)

// Avg represents a moving average.
type Avg struct {
	mu    sync.Mutex
	last  int64
	vals  [10]int64
	count int64
	idx   uint64
	avg   float64
}

// Add count and return current average.
func (a *Avg) Add(count int) float64 {
	now := time.Now().UnixNano()
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.last == 0 {
		a.last = now
	}
	a.count += int64(count)
	dif := now - a.last
	if dif > 1e9/int64(len(a.vals)) {
		a.vals[a.idx%uint64(len(a.vals))] = a.count
		a.idx++
		var total int64
		for i := 0; i < len(a.vals); i++ {
			total += a.vals[i]
		}
		a.avg = float64(total) / (float64(dif) / 1e9) / float64(len(a.vals))
		a.count = 0
		a.last = now
	}
	return a.avg
}
