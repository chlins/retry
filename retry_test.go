package retry

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	// normal condition
	var changed bool
	fn := func() error {
		changed = true
		return nil
	}

	err := Do(context.TODO(), fn, 1)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if !changed {
		t.Fatalf("not changed")
	}

	// remained time
	var n int
	fn = func() error {
		n++
		if n != 3 {
			return fmt.Errorf("n is %d", n)
		}
		return nil
	}
	err = Do(context.TODO(), fn, 2)
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Fatalf("expected err")
	}

	n = 0
	err = Do(context.TODO(), fn, 3)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	// ctx timeout
	ctx, cancel := context.WithCancel(context.Background())
	fn = func() error {
		time.Sleep(1 * time.Second)
		return errors.New("err")
	}
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
	err = Do(ctx, fn, 10)
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Fatalf("expected err")
	}
}
