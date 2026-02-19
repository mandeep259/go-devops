package main

import "fmt"

type Server struct {
	Hostname string
	CPU      float64
	Memory   float64
	Disk     float64
	Running  bool
}

func getStatus(value float64) string {
	if value > 80.0 {
		return "CRITICAL"
	} else if value > 60.0 {
		return "WARNING"
	}
	return "NORMAL"
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

func (s Server) printReport() {
	fmt.Printf("Server Hostname: (%s)\n", s.Hostname)
	fmt.Printf("CPU: %.2f%% (%s)\n", s.CPU, s.getCPUStatus())
	fmt.Printf("Memory: %.2f%% (%s)\n", s.Memory, s.getMemoryStatus())
	fmt.Printf("Disk: %.2f%% (%s)\n", s.Disk, s.getDiskStatus())
	fmt.Println("-----------------------")

}

func (s Server) isHealthy() bool {
	if s.CPU > 90.0 || s.Memory > 90.0 || s.Disk > 90.0 {
		return false
	}
	return true
}

func (s *Server) increaseCPU() {
	s.CPU += 10.0
}

func (s *Server) restart() {
	fmt.Println("Restarting the server...", s.Hostname)
	s.CPU = 0.0
	s.Running = true

}

func printBanner() {
	fmt.Printf("====================================\n")
	fmt.Printf("	Server Monitoring Report		 \n")
	fmt.Printf("====================================\n")
}

func main() {
	s := Server{
		Hostname: "server1",
		CPU:      85.0,
		Memory:   70.0,
		Disk:     90.0,
	}
	printBanner()
	s.printReport()
	s.increaseCPU()
	s.restart()
	fmt.Println("=== After Increase CPU ===")
	s.increaseCPU()
	fmt.Println("=== After Restart ===")
	fmt.Println("-------------")
	s.printReport()

}
