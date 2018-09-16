package tools

import "time"

// RetryIn retries passed function `call`
// until it end with success, else repeat again.
func RetryIn(interval time.Duration, call func() bool) {
	for {
		if ok := call(); ok {
			return
		}
		time.Sleep(interval)
	}
}

// RetryIncrementally retries passed function `call` until it end with success,
// else repeat again and increments retrying interval up to 2 hour.
func RetryIncrementally(interval time.Duration, call func() bool) {
	var counter int64

	for {
		counter++
		if ok := call(); ok {
			return
		}
		time.Sleep(incrementInterval(interval, counter))
	}
}

func incrementInterval(interval time.Duration, counter int64) time.Duration {
	quartile := int64(interval) / 4
	addendum := quartile * counter
	result := interval + time.Duration(addendum)

	if result > 2*time.Hour {
		result = 2 * time.Hour
	}
	return result
}
