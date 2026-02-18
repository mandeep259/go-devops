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

func PrintServer(s Server) {
	fmt.Printf("Server Hostname: %s\n", s.Hostname)
	fmt.Printf("Server CPU: %.2f%%\n", s.CPU)
	fmt.Printf("Server Memory: %.2f%%\n", s.Memory)
	fmt.Printf("Server Disk: %.2f%%\n", s.Disk)
	// fmt.Println("Server Hostname:", s.Hostname)
	// fmt.Println("Server CPU:", s.CPU)
	// fmt.Println("Server Memory:", s.Memory)
	// fmt.Println("Server Disk:", s.Disk)
}

func main() {
	Server1 := Server{
		Hostname: "k8s-Node-01",
		CPU:      75.5,
		Memory:   85.0,
		Disk:     60.0,
	}
	Server2 := Server{
		Hostname: "k8s-Node-02",
		CPU:      60.0,
		Memory:   70.0,
		Disk:     50.0,
	}
	printBanner()
	PrintServer(Server1)
	fmt.Println("--------------------")
	PrintServer(Server2)
	fmt.Println("-------------------")
	// fmt.Println("Server 1 Hostname:", Server1.Hostname)
	// fmt.Println("Server 1 CPU:", Server1.CPU)
	// fmt.Println("Server 1 Memory:", Server1.Memory)
	// fmt.Println("Server 1 Disk:", Server1.Disk)
	// fmt.Println("--------------------")
	// fmt.Println("Server 2 Hostname:", Server2.Hostname)
	// fmt.Println("Server 2 CPU:", Server2.CPU)
	// fmt.Println("Server 2 Memory:", Server2.Memory)
	// fmt.Println("Server 2 Disk:", Server2.Disk)
}
