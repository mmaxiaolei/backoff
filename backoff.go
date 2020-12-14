package backoff

import (
	"time"
)

type backoff interface {
	Next() time.Duration
	Reset()
}
