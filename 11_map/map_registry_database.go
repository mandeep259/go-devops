
package main

import "fmt"

type Server struct {
	Hostname string
	CPU		float64
	RAM		float64
	IsRunning bool
}

func printBanner () {
    fmt.Println("-----------------------------------------")
    fmt.Println("   Server Monitoring Details")
    // Added a header for the list
    fmt.Printf("%-15s %-8s %-8s %-10s\n", "Hostname", "CPU", "RAM", "Status")
    fmt.Println("-----------------------------------------")
}

func main () {

	serverOne := Server{"Node-A", 4, 16, true}
	serverTwo := Server{"Node-B", 8, 32, false}
	serverThree := Server{"Node-C", 16, 64, true}

	servers := []Server{serverOne, serverTwo, serverThree}

	printBanner()
// Everything related to printing individual servers 
// must stay INSIDE the loop braces.

	for _, s := range servers {
		status := "Offline" 
			if s.IsRunning {
				status = "Online"
			}
			fmt.Printf("%-15s %-8.0f %-8.0f %-10s \n", s.Hostname, s.CPU, s.RAM, status)
	}
	
	fmt.Println("Server Registry", servers)
	fmt.Println("-----------------------------------------")
	fmt.Println("Total Servers", len(servers))
	fmt.Println("Slice Capacity", cap(servers))



	
}
