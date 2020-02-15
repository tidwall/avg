package avg

import (
	"math/rand"
	"testing"
	"time"
)

func TestAvg(t *testing.T) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	var a Avg
	start := time.Now()
	total := 0
	for time.Since(start) < time.Second/8 {
		count := rand.Intn(100000)
		total += count
		a.Add(count)
	}
	elapsed := time.Since(start)
	avg1 := a.Add(0)
	avg2 := float64(total) / elapsed.Seconds() / 8
	if avg1/avg2 < 0.60 {
		t.Fatalf("below an acceptable 60%% (got %f%%)", avg1/avg2*100)
	}
}
