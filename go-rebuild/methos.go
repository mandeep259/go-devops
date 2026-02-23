package main

import "fmt"

type Server struct {
	Hostname  string
	CPU       float64
	Memory    float64
	Disk      float64
	isRunning bool
}

func printBanner() {
	fmt.Println("====================================")
	fmt.Println("	Server Monitoring Report		 ")
	fmt.Println("====================================")
}

func (s *Server) isOperational() bool {
	return s.isRunning
}

func (s Server) IsDiskLow(limit float64) bool {
	return s.Disk < limit
}

func (s *Server) AddCPU(value float64) {
	s.CPU += value
}

func (s *Server) SetOperational(status bool) {
	s.isRunning = status
}

func main() {
	printBanner()
	Node1 := Server{
		Hostname:  "K8s-Node-01",
		CPU:       75.5,
		Memory:    80.0,
		Disk:      60.0,
		isRunning: true,
	}
	Node2 := Server{
		Hostname:  "K8s-Node-02",
		CPU:       85.0,
		Memory:    90.0,
		Disk:      70.0,
		isRunning: false,
	}

	fmt.Println("Server Name : ", Node1.Hostname)
	fmt.Println("CPU : ", Node1.CPU, "%")
	fmt.Println("Memory : ", Node1.Memory, "MB")
	fmt.Println("Disk :", Node1.Disk, "GB")
	fmt.Println("Is Operational? ", Node1.isOperational())
	fmt.Println("----------------------------")
	fmt.Println("Server Name : ", Node2.Hostname)
	fmt.Println("CPU : ", Node2.CPU, "%")
	fmt.Println("Memory : ", Node2.Memory, "MB")
	fmt.Println("Disk :", Node2.Disk, "GB")
	fmt.Println("Is Operational? ", Node2.isOperational())

	fmt.Println("----------------------------")
	fmt.Println("Increasing CPU usage for Node1...")
	// Node1.CPU += 10.0
	fmt.Println("After increase the CPU", Node1.CPU, "%")
	fmt.Println("----------------------------")
	DiskStatus := Node1.IsDiskLow(80.0)
	fmt.Println("Is Disk Low? ", DiskStatus)

	fmt.Println("----------------------------")
	fmt.Println("Increasing CPU usage for Node1 using method...")
	Node1.AddCPU(10.0)
	fmt.Println("After increase the CPU", Node1.CPU, "%")
	fmt.Println("----------------------------")
	fmt.Println("Setting Node2 as operational...")
	Node2.SetOperational(true)
	fmt.Println("Is Node2 Operational? ", Node2.isOperational())
}
