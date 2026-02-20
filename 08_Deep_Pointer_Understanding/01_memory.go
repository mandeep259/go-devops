package main

import "fmt"

func increase(value *int) {
	*value += 10
}

type Server struct {
	CPU float64
}

func main() {

	// cpu := 80
	// ptr := (&cpu)

	s := Server{CPU: 50}
	// s.CPU = 75
	ptr := &s // This will cause a compilation error because ptr is of type *int and cannot point to a Server struct.
	ptr.CPU = 75
	fmt.Println("CPU usage is", ptr.CPU, "%")

	// var ptr *int
	// *ptr = 100

	// fmt.Println("CPU usage is", cpu, "%")
	// fmt.Println("CPU usage is", *ptr, "%")
	// fmt.Println("CPU usage is", cpu, "%")
	fmt.Println(*ptr)
	// fmt.Println(ptr)

	x := 10
	ptrX := &x

	*ptrX = 20

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of ptrX:", ptrX)
	fmt.Println("Dereferenced value of ptrX:", *ptrX)

	value := 5
	fmt.Println("Before increase:", value)
	increase(&value)
	fmt.Println("After increase:", value)

}
