package concurrency

import (
	"fmt"
	"github.com/romangurevitch/golang-concurrency/internal/concurrency/counter"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// UnexpectedResult what did you expect?
func UnexpectedResult() int {
	c := &counter.Simple{}

	go func() {
		for i := 0; i < 1000; i++ {
			c.Counter++
		}
	}()

	return c.Counter
}

// IncorrectResult ohh no!
func IncorrectResult() int {
	c := &counter.Simple{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			c.Counter++
		}()
	}

	wg.Wait()
	return c.Counter
}

// FinallySomethingWorksAsExpected but is it?
func FinallySomethingWorksAsExpected() int {
	c := &counter.Simple{}
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			lock.Lock()
			c.Counter++
			lock.Unlock()
		}()
	}

	wg.Wait()
	return c.Counter
}

// WorkingEndlessly is that a good idea?
func WorkingEndlessly() int {
	c := &counter.Simple{}
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	wg.Add(1)
	go func() {
		for {
			lock.Lock()
			c.Counter++
			lock.Unlock()
		}
	}()

	wg.Wait()
	return c.Counter
}

// WorkingEndlesslyWithAWayOut is it good enough though?
func WorkingEndlesslyWithAWayOut() int {
	lock := sync.Mutex{}
	c := &counter.Simple{}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			lock.Lock()
			c.Counter++
			lock.Unlock()
		}
	}()

	fmt.Println("Waiting...")
	<-sigs

	lock.Lock()
	result := c.Counter
	lock.Unlock()

	return result
}

// WorkingEndlesslyWithAGoodWayOut yes?
func WorkingEndlesslyWithAGoodWayOut() int {
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	c := &counter.Simple{}

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
				lock.Lock()
				c.Counter++
				inlinePrint(c.Counter)
				lock.Unlock()
			}

		}
	}()

	fmt.Println("Working, press ^C to stop")
	wg.Wait()

	return c.Counter
}

// WorkingUntilContextIsDone yes?
func WorkingUntilContextIsDone(ctx context.Context) int {
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	c := &counter.Simple{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("\nContext is done, shutting down!")
				return
			default:
				lock.Lock()
				c.Counter++
				inlinePrint(c.Counter)
				lock.Unlock()
			}

		}
	}()

	fmt.Println("Working as long as the context is not done")
	wg.Wait()

	return c.Counter
}

func inlinePrint(result int) {
	fmt.Print("\033[G\033[K", result)
}
