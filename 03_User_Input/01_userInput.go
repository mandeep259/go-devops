package main

import "fmt"

func main() {
	var hostname string
	var cpu float64
	var memory float64
	var choice int

	fmt.Print("Enter the hostname: ")
	fmt.Scanln(&hostname)

	fmt.Print("Enter CPU usage: ")
	fmt.Scanln(&cpu)

	fmt.Print("Enter Memory Usage: ")
	fmt.Scanln(&memory)

	fmt.Println("1. Restart Service")
	fmt.Println("2. Checking logs")
	fmt.Println("3. Exit")

	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Restarting service...")
	case 2:
		fmt.Println("Checking logs...")
	case 3:
		fmt.Println("Exiting...")
	default:
		fmt.Println("Invalid choice!")
	}

	for i := 1; i < 5; i++ {
		fmt.Println("Health check attemp", i)

		if cpu >= 80 {
			fmt.Println("CPU usage is critical!")
		} else if cpu >= 60 {
			fmt.Println("CPU usage is warning!")
		} else {
			fmt.Println("CPU usage is normal.")
		}

		if memory >= 16 {
			fmt.Println("Memory usage is critical!")
		} else if memory >= 8 {
			fmt.Println("Memory usage is warning!")
		} else {
			fmt.Println("Memory usage is normal.")
		}

		fmt.Println("Monitoring Server:", hostname)
		fmt.Println("CPU Usage: ", cpu, "%")
		fmt.Println("Memory Usage: ", memory, "GB")
	}
}
