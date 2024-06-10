package retry_test

import (
	"errors"
	"testing"
	"time"

	"github.com/amnestia/tnderlike/pkg/retry"
	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {
	var (
		mockError = errors.New("mock error")
	)
	type args struct {
		maxAttempts     int
		expectedSuccess int
		backoff         time.Duration
	}
	type want struct {
		err bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "When_RetryInfinite_Success",
			args: args{
				maxAttempts:     -1,
				expectedSuccess: 20,
			},
			want: want{
				err: false,
			},
		},
		{
			name: "When_Retry_Failed",
			args: args{
				maxAttempts:     5,
				expectedSuccess: 10,
			},
			want: want{
				err: true,
			},
		},
		{
			name: "When_Retry_Success",
			args: args{
				maxAttempts:     5,
				expectedSuccess: 3,
			},
			want: want{
				err: false,
			},
		},
		{
			name: "When_Retry_SuccessMax",
			args: args{
				maxAttempts:     5,
				expectedSuccess: 5,
			},
			want: want{
				err: false,
			},
		},
		{
			name: "When_RetryWithBackoff_Success",
			args: args{
				maxAttempts:     4,
				expectedSuccess: 3,
				backoff:         1,
			},
			want: want{
				err: false,
			},
		},
		{
			name: "When_RetryWithBackoff_Failed",
			args: args{
				maxAttempts:     3,
				expectedSuccess: 10,
				backoff:         1,
			},
			want: want{
				err: true,
			},
		},
		{
			name: "When_RetryWithBackoff_SuccessMax",
			args: args{
				maxAttempts:     5,
				expectedSuccess: 5,
				backoff:         1,
			},
			want: want{
				err: false,
			},
		},
		{
			name: "When_NewPolicy_InvalidMaxAttempts",
			args: args{
				maxAttempts:     -10,
				expectedSuccess: 10,
				backoff:         1,
			},
			want: want{
				err: true,
			},
		},
		{
			name: "When_NewPolicy_InvalidBackoff",
			args: args{
				maxAttempts:     10,
				expectedSuccess: 10,
				backoff:         -1,
			},
			want: want{
				err: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentAttempt := 0
			policy, err := retry.NewPolicy(tt.args.maxAttempts, tt.args.backoff)
			if policy == nil && tt.want.err {
				assert.Error(t, err)
				return
			}
			err = policy.Execute(func() (bool, error) {
				if currentAttempt == tt.args.expectedSuccess {
					return false, nil
				}
				currentAttempt++
				return true, mockError
			})
			if tt.want.err {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
