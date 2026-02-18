package main

import "fmt"

type Server struct {
	Hostname string
	CPU      float64
	Memory   float64
	Disk     float64
}

func printBanner() {
	fmt.Println("====================================")
	fmt.Println("	Server Monitoring Report		 ")
	fmt.Println("====================================")
}

func getStatus(value float64) string {
	if value > 80.0 {
		return "CRITICAL"
	} else if value > 60.0 {
		return "WARNING"
	}
	return "NORMAL"
}

func (s Server) printReport() {
	fmt.Printf("Server Hostname: %s\n", s.Hostname)
	fmt.Printf("CPU: %.2f%% (%s)\n", s.CPU, s.getCPUStatus())
	fmt.Printf("Memory: %.2f%% (%s)\n", s.Memory, s.getMemoryStatus())
	fmt.Printf("Disk: %.2f%% (%s)\n", s.Disk, s.getDiskStatus())
	fmt.Println("Overall Healthy:", s.isHealthy())
	fmt.Println("------------------------------")
}

func (s Server) getCPUStatus() string {
	return getStatus(s.CPU)
}

func (s Server) getMemoryStatus() string {
	return getStatus(s.Memory)
}
func (s Server) getDiskStatus() string {
	return getStatus(s.Disk)
}

func (s Server) isHealthy() bool {
	return s.CPU <= 80.0 &&
		s.Memory <= 80.0 &&
		s.Disk <= 80.0
}

func main() {
	printBanner()
	Server1 := Server{
		Hostname: "K8s-Node-01",
		CPU:      75.5,
		Memory:   80.0,
		Disk:     60.0,
	}
	Server2 := Server{
		Hostname: "K8s-Node-02",
		CPU:      85.0,
		Memory:   70.0,
		Disk:     90.0,
	}
	Server1.printReport()
	Server2.printReport()
}
