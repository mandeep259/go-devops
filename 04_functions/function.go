package main

import "fmt"

func sayHello() {
	fmt.Println("Hello Mandeep, Welcome to Go for Devops!")
}

func checkCPU(server string) {
	fmt.Println("Checking CPU for", server)
}

func getCPUUsage(cpuUsage float64) float64 {
	//return cpuUsage
	fmt.Println("CPU Usage:", cpuUsage, "%")
	return cpuUsage
}

func checkCPU1(cpu float64) string {
	if cpu > 80 {
		return "Critical"
	} else if cpu >= 60 {
		return "Warning"
	} else {
		return "Normal"
	}
}

func printBanner() {
	fmt.Println("====================================")
	fmt.Println("         Server Monitoring Report            ")
	fmt.Println("====================================")
}

func checkMemory(memory float64) string {
	if memory > 85 {
		return "Critical"
	} else if memory >= 60 {
		return "Warning"
	} else {
		return "Normal"
	}
}

func main() {
	printBanner()
	sayHello()
	checkCPU("k8s-node-1")
	cpu := getCPUUsage(90)
	fmt.Println("Final CPU Usage:", cpu, "%")
	CPUstatus := checkCPU1(81)
	fmt.Println("CPU Status:", CPUstatus)
	memoryStatus := checkMemory(85)
	fmt.Println("Memory Status:", memoryStatus)
}
