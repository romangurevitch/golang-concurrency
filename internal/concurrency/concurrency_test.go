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
	result := UnexpectedResult()
	printResultColor(red, sprintResult(result))
}

func TestUnexpectedResultFix(t *testing.T) {
	result := UnexpectedResultFix()
	printResultColor(green, sprintResult(result))
}

// Unexpected results since it is not thread safe, run with -race to find the issue.
func TestLetsMakeASmallChange(t *testing.T) {
	result := LetsMakeASmallChange()
	printResultColor(red, sprintResult(result))
}

// Good example with expected results and thread safe.
func TestFinallySomethingWorksAsExpected(t *testing.T) {
	result := FinallySomethingWorksAsExpected()
	printResultColor(green, sprintResult(result))
}

// Good example with expected results and thread safe.
func TestFinallySomethingWorksAsExpectedSafeCounter(t *testing.T) {
	result := FinallySomethingWorksAsExpectedSafeCounter()
	printResultColor(green, sprintResult(result))
}

// Bad endless go func example, terminating the run will terminate without printing the results.
func TestNonStoppingGoRoutine(t *testing.T) {
	fmt.Println(white, "Working, press ^C to stop", reset)

	result := NonStoppingGoRoutine()
	printResultColor(red, sprintResult(result))
}

// Using OS signals to catch termination signal to print out simpleCounter results.
// Good example how to make sure resources are closed when terminating running processes.
func TestNonStoppingGoRoutineWithShutdown(t *testing.T) {
	fmt.Println(white, "Working, press ^C to stop", reset)

	result, graceful := NonStoppingGoRoutineWithShutdown()
	printResultColor(red, sprintResult(result), sprintGraceful(graceful))
}

// Using OS signals to catch termination signal to print out simpleCounter results.
// Good example how to make sure resources are closed when terminating running processes.
func TestNonStoppingGoRoutineCorrectShutdown(t *testing.T) {
	fmt.Println(white, "Working, press ^C to stop", reset)

	result, graceful := NonStoppingGoRoutineCorrectShutdown()
	printResultColor(green, sprintResult(result), sprintGraceful(graceful))
}

// Using OS signals to catch termination signal to print out simpleCounter results.
// Bonus example, tiny change big impact!
func TestNonStoppingGoRoutineCorrectShutdownBonus(t *testing.T) {
	fmt.Println(white, "Working, press ^C to stop", reset)

	result, graceful := NonStoppingGoRoutineCorrectShutdownBonus()
	printResultColor(green, sprintResult(result), sprintGraceful(graceful))
}

func printResultColor(color string, messages ...string) {
	str := "\n\n" + sprintPrefixLine(color)
	for _, message := range messages {
		str += sprintColor(color, message, reset)
	}
	str += sprintSuffixLine(color) + "\n"

	fmt.Print(str)
}

func sprintColor(color string, a ...interface{}) string {
	return fmt.Sprintln(color, fmt.Sprint(a...), reset)
}

func sprintPrefixLine(color string) string {
	return fmt.Sprintln(color, strings.Repeat("-", 25), "Results", strings.Repeat("-", 25), reset)
}

func sprintSuffixLine(color string) string {
	return fmt.Sprintln(color, strings.Repeat("-", 22), "Test Finished", strings.Repeat("-", 22), reset)
}

func sprintResult(result int) string {
	return fmt.Sprint(prefix, "Counter value: ", result)
}

func sprintGraceful(graceful bool) string {
	return fmt.Sprint(prefix, "Graceful shutdown: ", graceful)
}
