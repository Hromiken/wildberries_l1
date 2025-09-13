package main

import (
	"context"
	"time"
)

func main() {
	mySleepAfter(10 * time.Second)
	mySleepTimer(10 * time.Second)
	mySleepContext(10 * time.Second)
}

func mySleepAfter(duration time.Duration) {
	<-time.After(duration)
}

func mySleepContext(duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	<-ctx.Done()
}

func mySleepTimer(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
	timer.Stop()
}
