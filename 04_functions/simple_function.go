package main

import "fmt"

func printBanner() {
	fmt.Println("====================================")
	fmt.Println("	Server Monitoring Report		 ")
	fmt.Println("====================================")
}

func checkCPU(server string) {
	fmt.Println("Checking CPU for", server)
}

func getCPUUsage() float64 {
	return 75.5
}

func getCPUStatus(cpu float64) string {
	if cpu > 85 {
		return ("Critical")
	} else if cpu >= 60 {
		return ("Warning")
	}
	return ("Normal")
}

func getMemoryUsage() float64 {
	return 85.0
}

func getMemoryStatus(memory float64) string {
	if memory > 80 {
		return ("Critical")

	} else if memory >= 60 {
		return ("Warning")
	}
	return ("Normal")
}

func main() {
	printBanner()
	checkCPU("k8s-Node-01")

	cpu := getCPUUsage()
	cpustatus := getCPUStatus(cpu)

	memory := getMemoryUsage()
	memStatus := getMemoryStatus(memory)

	fmt.Println("Current CPU Usage:", cpu)
	fmt.Println("CPU Status:", cpustatus)
	fmt.Println("Curent Memory Usage:", memory)
	fmt.Println("Memory Status:", memStatus)
}
