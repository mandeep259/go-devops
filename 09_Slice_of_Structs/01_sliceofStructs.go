package main 

import "fmt"

type Server struct {
	Hostname string
	CPUCore int
	RAM int
	IsRunning bool
}

func (s *Server) ToggleStatus() {
	s.IsRunning = !s.IsRunning
}

func main () {

	// servers := []Server{
	// 	{Hostname: "Server1", CPUCore: 4, RAM: 16, IsRunning: true},
	// 	{Hostname: "Server2", CPUCore: 8, RAM: 32, IsRunning: false},
	// 	{Hostname: "Server3", CPUCore: 16, RAM: 64, IsRunning: true},
	// }
	// 3 servers in the Rack 

	// what is slice ?
	// slice is a dynamic array that can grow and shrink in size. It is built on top of an array and provides a more flexible way to work with collections of data. A slice has three components: a pointer to the underlying array, the length of the slice (the number of elements it contains), and the capacity of the slice (the maximum number of elements it can hold before needing to resize).
	// In Go, slices are commonly used to work with collections of data because they provide a convenient and efficient way to manage dynamic arrays. They allow you to easily add, remove, and manipulate elements without worrying about the underlying array's size or memory management.
	// In this example, we have defined a struct called Server with fields for Hostname, CPUCore, RAM, and IsRunning. We then create three instances of the Server struct (Server1, Server2, and Server3) and store them in a slice called servers. We can then iterate over the slice to perform operations on each server, such as toggling their status or printing their details.
	// A slice of structs is a collection of struct instances that are stored in a slice. It allows you to manage and manipulate multiple instances of a struct in a dynamic and flexible way. In this example, we have a slice of Server structs, which allows us to easily manage and perform operations on multiple servers without needing to worry about the underlying array's size or memory management.
	
	Server1 := Server {"Node-A", 4, 16, true}
	Server2 := Server {"Node-B", 8, 32, false}
	Server3 := Server {"Node-C", 16, 64, true}

	//put them in to a Rack (slice of servers) 

	servers := [] Server{Server1, Server2, Server3}

	// Walk through the Rack and print the details of each server.

	for _, s := range servers {
		s.ToggleStatus()
		fmt.Println("Checking Hostname:", s.Hostname)
		fmt.Println("is running:", s.IsRunning)
	}

	for i := range servers {
		servers[i].ToggleStatus()
		fmt.Println("Checking Hostname:", servers[i].Hostname)
		fmt.Println("is running:", servers[i].IsRunning)
	}
	// downCount := 0

	// for _, s := range servers {	
	// 	if !s.IsRunning {
	// 		downCount++
	// 	}
		

	// for i := range servers {
	// 	servers[i].ToggleStatus()
	// }
	// s.ToggleStatus()

		// fmt.Println("Server:", s.Hostname)
		// fmt.Println("is running:", s.IsRunning)
		// fmt.Println("Total down servers:", downCount)
		// fmt.Println("-------------")
	}


