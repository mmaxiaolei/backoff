package backoff

import (
	"testing"
	"time"
)

func TestExponential_Next(t *testing.T) {
	e := NewExponentialBackoff(WithFactor(2), WithJitterFactor(0))
	equal(t, 500*time.Millisecond, e.Next())
	equal(t, 1000*time.Millisecond, e.Next())
}

func TestExponential_NextWithJitter(t *testing.T) {
	e := NewExponentialBackoff(WithJitterFactor(0.5), WithFactor(2))
	for i := 0; i < 20; i++ {
		println(e.Next() / time.Millisecond)
	}
}

func equal(t *testing.T, except time.Duration, actual time.Duration) {
	if except != actual {
		t.Fatalf("except: %dms,actual: %dms", except/time.Millisecond, actual/time.Millisecond)
	}
}
