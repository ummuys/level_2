package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	mCH := make(chan interface{})
	only := sync.Once{}
	closeCh := func(c chan interface{}) {
		only.Do(func() {
			close(c)
		})
	}

	for _, ch := range channels {
		ch := ch
		go func() {
			select {
			case <-ch:
				closeCh(mCH)
			case <-mCH:
				break
			}
		}()
	}
	return mCH
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
