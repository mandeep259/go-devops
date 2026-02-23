
package main

import "fmt"
type Server struct {
	Hostname string
	CPUCore int
	RAM int
	IsRunning bool
}


func (s *Server) UpdateHostname(newHostname string) {
	s.Hostname = newHostname
}

func (s *Server) UpdateRAM(newRAM int) bool {	 
	if newRAM >0 && newRAM <= 1024 {
		s.RAM = newRAM
		return true
	}
	return false

	
}

func (s *Server) ToggleStatus() {
	s.IsRunning = !s.IsRunning
}

func main () {
	Server1 := Server{
		Hostname: "Server1",
		CPUCore: 4,
		RAM: 16,
		IsRunning: true,
	}

	fmt.Println("Before updating RAM:", Server1.RAM)
	// fmt.Println("Server Details:", server1)
	// fmt.Println(Server1)
	// fmt.Printf("Hostname: %s, CPU Cores: %d, RAM: %dGB, Is Running: %t\n", Server1.Hostname, Server1.CPUCore, Server1.RAM, Server1.IsRunning)

	// Server1.RAM = 32
	// fmt.Printf("Updated RAM for %s: %dGB\n", Server1.Hostname, Server1.RAM)
	// Server1.UpdateRAM(0)
	// fmt.Println("After updating RAM:", Server1.RAM)
	// fmt.Printf("Updated RAM for %s: %dGB\n", Server1.Hostname, Server1.RAM)
	// fmt.Println("Before toggling status:", Server1.IsRunning)
	// Server1.ToggleStatus()
	// fmt.Println("After toggling status:", Server1.IsRunning)
	
	// fmt.Println("Before updating hostname:", Server1.Hostname)
	// Server1.UpdateHostname("NewServer1")
	// fmt.Println("After updating hostname:", Server1.Hostname)
	
	result := Server1.UpdateRAM(64)
	fmt.Println("After updating RAM:", Server1.RAM)
	fmt.Println("Update result:", result)
	// fmt.Println("After updating RAM:", Server1.RAM)
}
