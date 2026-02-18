package main

import (
	"fmt"
)

func main() {
	// var name string = "Mandeep"

	// fmt.Println("Hello", name)

	hostname := "k8s-node-1"
	cpuUsage := 60
	memUsage := 4096
	serviceRunning := true

	fmt.Println("Host:", hostname)
	fmt.Println("CPU Usage:", cpuUsage, "%")
	fmt.Println("Memory Usage:", memUsage, "MB")
	fmt.Println("Service Running:", serviceRunning)

	a := 10
	b := 3

	fmt.Println("Sum:", a+b)
	fmt.Println("Difference", a-b)
	fmt.Println("Product:", a*b)
	fmt.Println("Quotient:", a/b)
	fmt.Println("Remainder:", a%b)

	if cpuUsage > 80 {
		fmt.Println("CPU usage is Critical!")
	} else if cpuUsage >= 60 {
		fmt.Println("CPU usage is Warning!")
	} else {
		fmt.Println("CPU usage is Normal.")
	}

	running := true
	healthy := true

	if running && healthy {
		fmt.Println("Service is running and healthy.")
	} else {
		fmt.Println("Service is not running or not healthy.")
	}
}
