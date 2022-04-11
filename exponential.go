package exponential

import (
	"time"

	"github.com/go-toolbelt/jitter"
)

// Delay returns a function that iterative returns the next delay in an exponential series.
func Delay(initialDelay time.Duration, maxDelay time.Duration, factor float64) func() time.Duration {
	delay := initialDelay

	return func() time.Duration {
		jitteredDelay := jitter.UpTo(delay)

		delay = time.Duration(float64(initialDelay) * factor)
		if delay > maxDelay {
			delay = maxDelay
		}

		return jitteredDelay
	}
}
