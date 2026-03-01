package main

import "fmt"

type Server struct {
	Hostname string
	running bool
}


func NewServer(name string) *Server {
	if name == "" {
		panic("server name cannot be empty")
	}
	return &Server {
		Hostname : name,
	}
}

//Start returns true if it successfully started the server
func (s *Server) Start () bool {
	if s.running {
		return false // Already running
	}
	s.running = true
	return true
}

func (s Server) IsRunning () bool {
	return s.running

}

// Stop returns true if it successfully stopped the server

func (s *Server) Stop () bool {
	if !s.running {
		return false	// Already stopped
	}
	s.running = false
	return true
}

func (s Server) Status () string {
		// status := "Offline"
		if s.running {
			return "Online"
		}
		return "Offline"
}



func main () {

	// server1 := Server {
	// 	Hostname : "Node-A",
	// 	running : true,
	// }

	// server2 :=  Server {
	// 	Hostname : "Node-B",
	// 	running : false,
	// }

	server1 := NewServer("Node-A")
	server2 := NewServer("Node-B")
	server3 := NewServer("Node-C")
	server4 := NewServer("Node-D")


	fmt.Printf("%s current status: %s\n", server1.Hostname, server1.Status())
	fmt.Printf("%s current status: %s\n", server2.Hostname, server2.Status())
	fmt.Printf("%s Current status: %s\n", server3.Hostname, server3.Status())
	fmt.Printf("%s Current status: %s\n", server4.Hostname, server4.Status())
	fmt.Println("--------------------------------")

	// Start Node-B
	if server2.Start() {
		fmt.Println("Successfully started", server2.Hostname)
	}
	fmt.Println("Node-B now:", server2.Status())

	// Stop Node-A
	server1.Stop()
	fmt.Println("Node-A now:", server1.Status())

	
}
	// fmt.Println("Server Running State", server1.IsRunning)
	// fmt.Println("--------------------------------")
	// fmt.Println("Node-A Status", server1.Status())
	// fmt.Println("--------------------------------")
	// fmt.Println("Node-B Status", server2.Status())
	// server2.Start()
	// fmt.Println("After Node-B Status", server2.Status())

	// server1.Stop()
	// fmt.Println("After Node-A Status", server1.Status())


