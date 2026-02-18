package main

import "fmt"

func printBanner() {
	fmt.Println("====================================")
	fmt.Println("         Server Monitoring Report            ")
	fmt.Println("====================================")
}

func getCPUUsage(cpuUsage float64) float64 {
	return cpuUsage
}

func getCPUStatus(cpu float64) string {
	if cpu > 80 {
		return "Critical"
	} else if cpu >= 60 {
		return "Warning"
	}
	return "Normal"
}

func getMemoryStatus(memory float64) string {
	if memory > 85 {
		return "Critical"
	} else if memory >= 60 {
		return "Warning"
	}
	return "Normal"
}

func main() {
	printBanner()

	cpu := getCPUUsage(85)
	cpuStatus := getCPUStatus(cpu)

	memory := 89.0
	memoryStatus := getMemoryStatus(memory)

	fmt.Println("CPU Usage:", cpu, "%")
	fmt.Println("CPU Status:", cpuStatus)
	fmt.Println("Memory Status:", memoryStatus)
}
