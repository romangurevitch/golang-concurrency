package concurrency

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/romangurevitch/golang-concurrency/internal/concurrency/counter"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var white = "\033[97m"
var cursorBack = "\033[G\033[K"

// UnexpectedResult what did you expect?
func UnexpectedResult() int {
	c := &counter.SimpleCounter{}

	go func() {
		for i := 0; i < 1000; i++ {
			c.Inc()
		}
	}()

	return c.Count()
}

// UnexpectedResultFix is it fixed?
func UnexpectedResultFix() int {
	c := &counter.SimpleCounter{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			c.Inc()
		}
	}()

	wg.Wait()
	return c.Count()
}

// LetsMakeASmallChange ohh no!
func LetsMakeASmallChange() int {
	c := &counter.SimpleCounter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	return c.Count()
}

// FinallySomethingWorksAsExpected but is it?
func FinallySomethingWorksAsExpected() int {
	c := &counter.SimpleCounter{}
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			lock.Lock()
			c.Inc()
			lock.Unlock()
		}()
	}

	wg.Wait()
	return c.Count()
}

// FinallySomethingWorksAsExpectedSafeCounter but is it?
func FinallySomethingWorksAsExpectedSafeCounter() int {
	c := &counter.SafeCounter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	return c.Count()
}

// NonStoppingGoRoutine is that a good idea?
func NonStoppingGoRoutine() int {
	c := &counter.SafeCounter{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			inlinePrint(c.Inc())
		}
	}()

	wg.Wait()
	return c.Count()
}

// NonStoppingGoRoutineWithShutdown is it good enough though?
func NonStoppingGoRoutineWithShutdown() (int, bool) {
	c := &counter.SafeCounter{}
	gracefulShutdown := false

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer func() { gracefulShutdown = true }()

		for {
			inlinePrint(c.Inc())
		}
	}()

	<-sigs
	return c.Count(), gracefulShutdown
}

// NonStoppingGoRoutineCorrectShutdown yes?
func NonStoppingGoRoutineCorrectShutdown() (int, bool) {
	wg := sync.WaitGroup{}
	c := &counter.SafeCounter{}
	gracefulShutdown := false

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() { gracefulShutdown = true }()

		for {
			select {
			case <-sigs:
				return
			default:
				inlinePrint(c.Inc())
			}

		}
	}()

	wg.Wait()
	return c.Count(), gracefulShutdown
}

// NonStoppingGoRoutineCorrectShutdownBonus tiny change?
func NonStoppingGoRoutineCorrectShutdownBonus() (int, bool) {
	wg := sync.WaitGroup{}
	c := &counter.SafeCounter{}
	gracefulShutdown := false

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer func() { gracefulShutdown = true }() // <<< Tiny change
		defer wg.Done()                            // <<< Lets swap the defer commands

		for {
			select {
			case <-sigs:
				return
			default:
				inlinePrint(c.Inc())
			}

		}
	}()

	wg.Wait()
	return c.Count(), gracefulShutdown
}

func inlinePrint(result int) {
	fmt.Print(yellow, cursorBack, result, reset)
}
