package tools

import (
	"fmt"
	"testing"
	"time"
)

func TestRetryIn(t *testing.T) {
	var counter int
	call := func() bool {
		counter++
		fmt.Println(time.Now().UTC(), counter)
		return counter == 13
	}
	RetryIn(time.Second, call)
}

func TestRetryIncrementally(t *testing.T) {
	var counter int
	call := func() bool {
		counter++
		fmt.Println(time.Now().UTC(), counter)
		return counter == 13
	}
	RetryIncrementally(time.Second, call)
}

func TestRetryIncrementallyUntil(t *testing.T) {
	var counter int
	call := func() bool {
		counter++
		fmt.Println(time.Now().UTC(), counter)
		return counter == 13
	}

	ok := RetryIncrementallyUntil(time.Second, time.Minute, call)
	if !ok {
		t.Fatal("Required: true; got: false")
	}

	fmt.Println()
	counter = 0
	ok = RetryIncrementallyUntil(time.Second, 13*time.Second, call)
	if ok {
		t.Fatal("Required: false; got: true")
	}
}