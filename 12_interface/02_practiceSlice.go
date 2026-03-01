package main

import (
	"fmt"
)

func printBanner() {
	fmt.Println("-----------------------------------")
	fmt.Println("	Server Details			")
	fmt.Println("-----------------------------------")

}

type Server struct {
	Hostname string
	CPU      float64
	Memory   float64
	Disk     float64
	running  bool
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

func (s *Server) GetCPUStatus() string {
	return getStatus(s.CPU)
}

func (s *Server) GetMemoryStatus() string {
	return getStatus(s.Memory)
}

func (s *Server) GetDiskStatus() string {
	return getStatus(s.Disk)
}

func (s Server) IsRunning() bool {
	return s.running
}

func (s *Server) IsHealthy() bool {
	return s.running &&
		s.CPU <= 80 &&
		s.Memory <= 80 &&
		s.Disk <= 80
}

func (s Server) Status() string {
	// status := "Offline"
	if s.running {
		return "Online"
	}
	return "Offline"
}
func (s *Server) Start() bool {
	if s.running {
		return false // Already running
	}
	s.running = true
	return true
}

func (s *Server) Stop() bool {
	if !s.running {
		return false // Already stopped
	}
	s.running = false
	return true
}

func (s *Server) PrintReport() {
	fmt.Printf("%-18s : %s\n", "Hostname", s.Hostname)
	fmt.Printf("%-18s : %1.f%% [%s]\n", "CPU Load", s.CPU, s.GetCPUStatus())
	fmt.Printf("%-18s : %1.f%% [%s]\n", "Memory Usage", s.Memory, s.GetMemoryStatus())
	fmt.Printf("%-18s : %1.f%% [%s]\n", "Disk Usage", s.Disk, s.GetDiskStatus())
	fmt.Printf("%-18s : %s\n", "RunningStatus", s.Status())
	// fmt.Printf("%-18s : %t\n", "Node-Health Status", s.IsHealthy())
	// Fixed: Now using two verbs (%t for bool, %s for string)
	fmt.Printf("%-18s : %t (Overall: %s)\n", "Node Health", s.IsHealthy(), s.Status())
	fmt.Println("-----------------------------------")
}

func main() {
	1
	// 1. Group your servers into a slice
	inventory := []*Server{
		{"Node-A", 85.5, 16.0, 65.0, true},
		{"Node-B", 55.0, 8.0, 59.0, true},
		{"Node-C", 0, 0, 0, false},
	}

	printBanner()

	// 2. Use a 'for range' loop to print all reports
	for _, s := range inventory {
		s.PrintReport()
	}

	// 3. Automation: Start all Offline servers
	fmt.Println("\n>>> SYSTEM SCAN: Checking for Offline Nodes...")
	for _, s := range inventory {
		if !s.IsRunning() {
			fmt.Printf("RECOVERING: %s is down. Sending Start command...\n", s.Hostname)
			s.Start()
		}
	}

	fmt.Println("\n>>> FINAL STATUS REPORT:")
	for _, s := range inventory {
		fmt.Printf("%-18s : %s\n", s.Hostname, s.Status())
	}
}
