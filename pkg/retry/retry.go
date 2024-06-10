package retry

import (
	"fmt"
	"time"

	"github.com/amnestia/tnderlike/pkg/logger"
)

// Policy retry policy with the config
type Policy struct {
	attempt     int
	maxAttempts int
	backoff     time.Duration
	timer       *time.Timer
}

type Retryable func() (shouldRetry bool, err error)

// NewPolicy create new retry policy with maxAttempts and backoff(in milliseconds)
// backoff will be multiplied for each retry attempts until maxAttempts or success
// example:
// with backoff of 60(milliseconds)
// second retry backoff time will be 1 seconds 20 milliseconds(2 times of the backoff)
// third retry backoff time will be 1 seconds 80 milliseconds(3 times of the backoff)
func NewPolicy(maxAttempts int, backoff time.Duration) (*Policy, error) {
	if maxAttempts != -1 && maxAttempts < 1 {
		return nil, fmt.Errorf("invalid attempts, ex: -1(infinite) or >0(should be greater than zero)")
	}
	if backoff < 0 {
		return nil, fmt.Errorf("invalid backoff, ex: >=0(should be greater or equal to zero)")
	}
	return &Policy{
		maxAttempts: maxAttempts,
		backoff:     backoff,
	}, nil
}

func (p *Policy) nextTimer() {
	if p.backoff < 1 || p.attempt > p.maxAttempts {
		return
	}
	backoff := p.backoff * time.Duration(p.attempt) * time.Millisecond
	if p.timer == nil {
		p.timer = time.NewTimer(backoff)
	} else {
		p.timer.Reset(backoff)
	}
	<-p.timer.C
}

func (p *Policy) isFinished() bool {
	if p.maxAttempts != -1 && p.attempt > p.maxAttempts {
		return true
	}
	p.nextTimer()
	return false
}

// Execute run function with the configured retry policy
func (p *Policy) Execute(f Retryable) error {
	var (
		shouldRetry bool
		err         error
	)
	for {
		if p.isFinished() {
			break
		}
		shouldRetry, err = f()
		if !shouldRetry || err == nil {
			break
		}
		logger.Logger.Error().Err(logger.ErrorWrap(fmt.Errorf("Retry attempt #%d: %s", p.attempt, err), "retry.Execute"))
		p.attempt++
	}
	return err
}
