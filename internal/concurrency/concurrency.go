package concurrency

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/romangurevitch/golang-concurrency/internal/concurrency/counter"
)

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

// WorkingEndlessly is that a good idea?
func WorkingEndlessly() int {
	c := &counter.SafeCounter{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for {
			c.Inc()
		}
	}()

	wg.Wait()
	return c.Count()
}

// WorkingEndlesslyWithAWayOut is it good enough though?
func WorkingEndlesslyWithAWayOut() int {
	c := &counter.SafeCounter{}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			c.Inc()
		}
	}()

	fmt.Println("Waiting...")
	<-sigs

	return c.Count()
}

// WorkingEndlesslyWithAGoodWayOut yes?
func WorkingEndlesslyWithAGoodWayOut() int {
	wg := sync.WaitGroup{}
	c := &counter.SafeCounter{}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-sigs:
				fmt.Println("\nSignal received, shutting down!")
				return
			default:
				c.Inc()
				inlinePrint(c.Count())
			}

		}
	}()

	fmt.Println("Working, press ^C to stop")
	wg.Wait()

	return c.Count()
}

func inlinePrint(result int) {
	fmt.Print("\033[G\033[K", result)
}