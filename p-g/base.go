package main

import (
	"fmt"
	"time" // Added time package for weekday functionality
)

// 1. Basic syntax
func basicSyntax() {
	// Variables
	var name string = "Go Learner"
	age := 25 // Type inference
	const PI = 3.14159
	const (
		StatusOK = iota
		StatusError
	)

	// Data structures
	numbers := []int{1, 2, 3, 4, 5}
	user := map[string]string{
		"name": "Alice",
		"role": "Developer",
	}

	fmt.Println(name, age, PI, StatusOK, StatusError, numbers, user)
}

// 2. Flow control
func flowControl() {
	// If-else
	if num := 9; num < 0 {
		fmt.Println("Negative")
	} else if num < 10 {
		fmt.Println("Single digit")
	} else {
		fmt.Println("Multiple digits")
	}

	// Switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// For loop
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	
	// Range loop
	colors := []string{"red", "green", "blue"}
	for index, value := range colors {
		fmt.Printf("\nIndex: %d, Color: %s", index, value)
	}
}

// 3. Functions
func add(a, b int) int {
	return a + b
}

func multiReturn() (int, string) {
	return 42, "answer"
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 4. Methods and OOP
type Rectangle struct {
	Width, Height float64
}

// Method with pointer receiver
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Interface
type Geometry interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Interface implementation check
func calculate(g Geometry) {
	fmt.Printf("\nArea: %.2f, Perimeter: %.2f", g.Area(), g.Perimeter())
}

// 5. Error handling
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// 6. Advanced features
func closures() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	basicSyntax()
	flowControl()
	
	fmt.Println("\nAdd:", add(3, 4))
	val, str := multiReturn()
	fmt.Println(val, str)
	fmt.Println("Fibonacci(7):", fibonacci(7))

	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("Rectangle Area:", rect.Area())
	calculate(&rect) // Use pointer to satisfy interface implementation

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("\nError:", err)
	} else {
		fmt.Println("\nDivision result:", result)
	}

	// Closure usage
	counter := closures()
	fmt.Println("Counter:", counter(), counter(), counter())
}