package utils

import (
	"sync/atomic"
)

var servedCounter int64

func IncrementServedCounter() {
	atomic.AddInt64(&servedCounter, 1)
}

func GetServedCount() int64 {
	return atomic.LoadInt64(&servedCounter)
}
