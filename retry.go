package retry

import (
	"context"
	"errors"
	"fmt"
)

type task func() error

// Retry defines a task and retry count
// return nil when first run sucessfully
// otherwise, repeat it
func Retry(ctx context.Context, fn task, count int) error {
	if count < 1 {
		return errors.New("retry count must greater than zero")
	}

	var err error
	for {
		if count == 0 {
			// has no retry times
			return fmt.Errorf("has no retry times remained, last error is %v", err)
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("retry stopped due to context done, %v", ctx.Err())
		default:
			err = fn()
		}

		if err != nil {
			count--
		} else {
			break
		}
	}

	return err
}
