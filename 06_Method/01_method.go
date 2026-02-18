package main

import "fmt"

func printBanner() {
	fmt.Println("====================================")
	fmt.Println("	Server Monitoring Report		 ")
	fmt.Println("====================================")
}

type Server struct {
	Hostname string
	CPU      float64
	Memory   float64
	Disk     float64
}

func printReport(s Server) {
	fmt.Printf("Server Hostname: %s\n", s.Hostname)
	fmt.Printf("Server CPU: %.2f%%\n", s.CPU)
	fmt.Printf("Server Memory: %.2f%%\n", s.Memory)
	fmt.Printf("Server Disk: %.2f%%\n", s.Disk)
	fmt.Println("CPU Status:", s.getCPUstatus())
	fmt.Println("Memory Status:", s.getMemoryStatus())
	fmt.Println("Disk Status:", s.getDiskStatus())
	fmt.Println("isHealthy", s.isHealthy())
	fmt.Println("------------------------------")
}

func (s Server) getCPUstatus() string {
	if s.CPU > 80.0 {
		return "CRITICAL"
	} else if s.CPU > 60.0 {
		return "WARNING"
	} else {
		return "NORMAL"
	}
}

func (s Server) getMemoryStatus() string {
	if s.Memory > 80.0 {
		return "CRITICAL"
	} else if s.Memory > 60.0 {
		return "WARNING"
	}
	return "NORMAL"
}
func (s Server) getDiskStatus() string {
	if s.Disk > 80.0 {
		return "CRITICAL"
	} else if s.Disk > 60.0 {
		return "WARNING"
	}
	return "NORMAL"
}

func (s Server) isHealthy() bool {
	if s.CPU <= 80.0 && s.Memory <= 80.0 && s.Disk <= 80.0 {
		return true
	}
	return false
}

func main() {
	printBanner()
	Server1 := Server{
		Hostname: "k8s-Node-01",
		CPU:      75.5,
		Memory:   59.0,
		Disk:     81.0,
	}
	Server2 := Server{
		Hostname: "k8s-Node-02",
		CPU:      60.0,
		Memory:   70.0,
		Disk:     50.0,
	}
	// Server1.printReport()
	// Server2.printReport()
	printReport(Server1)
	printReport(Server2)
}
