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
