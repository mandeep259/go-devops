package main

import "fmt"

func printBanner() {
	fmt.Printf("====================================\n")
	fmt.Printf("	Server Monitoring Report		 \n")
	fmt.Printf("====================================\n")
}

type Server struct {
	Hostname string
	CPU      float64
	Memory   float64
	Disk     float64
	Running  bool
}

func getStatus(value float64) string {
	if value >= 80 {
		return "CRITICAL"
	} else if value >= 60 {
		return "WARNING"
	} else {
		return "NORMAL"
	}
}

func (s *Server) getCPUStatus() string {
	return getStatus(s.CPU)
}

func (s *Server) getmemoryStatus() string {
	return getStatus(s.Memory)
}

func (s *Server) getDiskStatus() string {
	return getStatus(s.Disk)
}

func (s *Server) isHealthy() bool {
	return s.CPU <= 80 &&
		s.Memory <= 80 &&
		s.Disk <= 80
}

func (s *Server) increaseCPU() {
	s.CPU += 10.0
}

func (s *Server) restart() {
	fmt.Println("Restarting the Service...", s.Hostname)
	s.CPU = 0.0
	s.Running = true

}

func (s *Server) printReport() {
	fmt.Printf("Server Hostname: %s\n", s.Hostname)
	fmt.Printf("CPU: %.2f%% (%s)\n", s.CPU, s.getCPUStatus())
	fmt.Printf("Memory: %.2f%% (%s)\n", s.Memory, s.getmemoryStatus())
	fmt.Printf("Disk: %.2f%% (%s)\n", s.Disk, s.getDiskStatus())
	fmt.Printf("Server Health: %t\n", s.isHealthy())
	fmt.Println("-----------------------")

}

func main() {

	Server1 := Server{
		Hostname: "Server-1",
		CPU:      85.5,
		Memory:   85.0,
		Disk:     80.0,
		Running:  true,
	}

	printBanner()
	Server1.printReport()

	Server1.increaseCPU()
	fmt.Println("=== After Increase CPU ===")
	Server1.printReport()

	Server1.restart()
	fmt.Println("=== Afetr restart ===")
	Server1.printReport()
}
