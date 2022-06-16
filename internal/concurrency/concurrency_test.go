package concurrency

import (
	"fmt"
	"strings"
	"testing"
)

const prefix = "[Test Output]\t"

// Bad example how to use go functions - run to see the results.
// Not thread safe, run with -race to find the issue.
func TestUnexpectedResult(t *testing.T) {
	printLine()
	result := UnexpectedResult()
	fmt.Println(prefix, "Result:", result)
	fmt.Println(prefix, "Test finished...")
	printLine()
}

func TestUnexpectedResultFix(t *testing.T) {
	printLine()
	result := UnexpectedResultFix()
	fmt.Println(prefix, "Result:", result)
	fmt.Println(prefix, "Test finished...")
	printLine()
}

// Unexpected results since it is not thread safe, run with -race to find the issue.
func TestLetsMakeASmallChange(t *testing.T) {
	printLine()
	result := LetsMakeASmallChange()
	fmt.Println(prefix, "Result:", result)
	fmt.Println(prefix, "Test finished...")
	printLine()
}

// Good example with expected results and thread safe.
func TestFinallySomethingWorksAsExpected(t *testing.T) {
	printLine()
	result := FinallySomethingWorksAsExpected()
	fmt.Println(prefix, "Result:", result)
	fmt.Println(prefix, "Test finished...")
	printLine()
}

// Good example with expected results and thread safe.
func TestFinallySomethingWorksAsExpectedSafeCounter(t *testing.T) {
	printLine()
	result := FinallySomethingWorksAsExpectedSafeCounter()
	fmt.Println(prefix, "Result:", result)
	fmt.Println(prefix, "Test finished...")
	printLine()
}

// Bad endless go func example, terminating the run will terminate without printing the results.
func TestNonStoppingGoRoutine(t *testing.T) {
	printLine()
	result := NonStoppingGoRoutine()
	fmt.Println(prefix, "Result:", result)
	fmt.Println(prefix, "Test finished...")
	printLine()
}

// Using OS signals to catch termination signal to print out simpleCounter results.
// Good example how to make sure resources are closed when terminating running processes.
func TestNonStoppingGoRoutineWithShutdown(t *testing.T) {
	printLine()
	result := NonStoppingGoRoutineWithShutdown()
	fmt.Println(prefix, "Result:", result)
	fmt.Println(prefix, "Test finished...")
	printLine()
}

// Using OS signals to catch termination signal to print out simpleCounter results.
// Good example how to make sure resources are closed when terminating running processes.
func TestNonStoppingGoRoutineCorrectShutdown(t *testing.T) {
	printLine()
	result := NonStoppingGoRoutineCorrectShutdown()
	fmt.Println(prefix, "Result:", result)
	fmt.Println(prefix, "Test finished...")
	printLine()
}

// Using OS signals to catch termination signal to print out simpleCounter results.
// Bonus example, tiny change big impact!
func TestNonStoppingGoRoutineCorrectShutdownBonus(t *testing.T) {
	printLine()
	result := NonStoppingGoRoutineCorrectShutdownBonus()
	fmt.Println(prefix, "Result:", result)
	fmt.Println(prefix, "Test finished...")
	printLine()
}

func printLine() {
	fmt.Println(strings.Repeat("-", 50))
}
