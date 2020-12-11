package backoff

import (
	"context"
	"fmt"
	"math"
	"time"
)

const (
	defaultFactor         float64       = 2.0
	defaultMinimumBackoff time.Duration = 1
)

func New(min, max time.Duration, factor float64) *Backoff {
	return &Backoff{
		min:    min,
		max:    max,
		factor: factor,
	}
}

type Backoff struct {
	// min represents the minimal backoff time
	min time.Duration

	// Max represents the maximal backoff time
	max time.Duration

	// factor represents the factor the backoff time grows exponentially
	factor float64

	attempt int
}

// Wait waits for the required time or returns when the context is cancelled.
func (b *Backoff) Wait(ctx context.Context) error {
	b.attempt++
	duration := b.duration(b.attempt)

	select {
	case <-ctx.Done():
		return fmt.Errorf("backoff: cancelled via context: %s", ctx.Err())
	case <-time.After(duration):
	}
	return nil
}

func (b *Backoff) duration(attempt int) time.Duration {
	min := defaultMinimumBackoff
	if b.min != time.Duration(0) {
		min = b.min
	}
	factor := defaultFactor
	if b.factor != 0 {
		factor = b.factor
	}
	d := float64(min) * math.Pow(factor, float64(attempt))
	if b.max != time.Duration(0) && d > float64(b.max) {
		d = float64(b.max)
	}
	return time.Duration(d)
}
