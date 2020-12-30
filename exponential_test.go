package backoff

import (
	"fmt"
	"testing"
	"time"
)

func TestExponential_Next(t *testing.T) {
	e := NewExponentialBackoff(
		WithMinInterval(10*time.Millisecond),
		WithFactor(2),
		WithJitterFactor(0))
	equal(t, 10*time.Millisecond, e.Next())
	equal(t, 20*time.Millisecond, e.Next())
	equal(t, 40*time.Millisecond, e.Next())
	equal(t, 80*time.Millisecond, e.Next())
	equal(t, 160*time.Millisecond, e.Next())
	equal(t, 320*time.Millisecond, e.Next())
	equal(t, 640*time.Millisecond, e.Next())
	equal(t, 1280*time.Millisecond, e.Next())
}

func TestExponential_NextWithJitter(t *testing.T) {
	e := NewExponentialBackoff(
		WithMinInterval(10*time.Millisecond),
		WithJitterFactor(0.5),
		WithFactor(2))
	for i := 0; i < 8; i++ {
		fmt.Printf("%dms\n", e.Next()/time.Millisecond)
	}
}

func equal(t *testing.T, except time.Duration, actual time.Duration) {
	fmt.Printf("%dms\n", actual/time.Millisecond)
	if except != actual {
		t.Fatalf("except: %dms,actual: %dms", except/time.Millisecond, actual/time.Millisecond)
	}
}
