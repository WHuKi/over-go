package utils

import (
	"time"

	"github.com/pkg/errors"
)

//ReTry  Re configures the settings to be used for retries
type ReTry struct {
	// Max specifies the maximum number of attempts. By default, there is no
	// maximum (inifinite retries).
	Max uint

	// Delay specifies the initial Delay between attempts. By default, the
	// initial Delay is set to 0.
	Delay time.Duration

	// MaxDelay specifies a maximum value for the Delay. By default, no
	// maximum is applied.
	MaxDelay time.Duration

	// Backoff specifies a multiplier to be applied to the Delay between
	// attempts. The default Backoff is 1.
	Backoff float64

	// Jitter specifies a fixed number of seconds added to the Delay between
	// attempts. By default, no Jitter is added.
	Jitter time.Duration

	// RetryableErrors specifies an array of merry.Error's that should be caught by the
	// retry. If not specified, all RetryableErrors are caught.
	RetryableErrors []error
}

// Func is a function type that takes no arguments and returns an error. This is meant
// to wrap the functionality that should be retried.
type Func func() error

// Try executes the provided function using the Re fields
func (r ReTry) Try(f Func) error {
	i := uint(0)
	delay := r.Delay
	for {
		err := f()
		if err == nil {
			break
		}

		err2 := r.checkErrors(err)
		if err2 != nil {
			err2 = errors.WithMessage(err, "retry: unexpected error:")
			return err2
		}

		i++
		if i == r.Max {
			err = errors.Wrapf(err, "retry: max retries reached (%v):", r.Max)
			return err
		}

		if delay > 0 {
			time.Sleep(delay)
		}

		if r.Backoff != 0 {
			delay = time.Duration(float64(delay)*r.Backoff) + r.Jitter
		} else {
			delay = delay + r.Jitter
		}

		if r.MaxDelay > 0 && delay > r.MaxDelay {
			delay = r.MaxDelay
		}
	}

	return nil
}

func (r ReTry) checkErrors(err error) error {
	if len(r.RetryableErrors) == 0 {
		return nil
	}

	for _, e := range r.RetryableErrors {
		if e == err {
			return nil
		}
	}

	return err
}
