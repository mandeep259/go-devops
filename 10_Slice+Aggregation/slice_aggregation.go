package main

import "fmt"

type Server struct {
	Hostname  string
	CPUCore   int
	RAM       int
	IsRunning bool
}

func CountRunningServers(servers []Server) int {
	count := 0
	for _, s := range servers {
		if s.IsRunning {
			count++
		}
	}
	return count
}

func CountDownServers(servers []Server) int {
	count := 0
	for _, s := range servers {
		if !s.IsRunning {
			count++
		}
	}
	return count
}

func TotalRAM(servers []Server) int {
	total := 0
	for _, s := range servers {
		total += s.RAM
		// fmt.Printf("%+v\n", s)
	}
	return total
}

func TotalCPUCore(servers []Server) int {
	total := 0
	for _, s := range servers {
		total += s.CPUCore
	}
	return total
}

func HighestRAM(servers []Server) int {
	 highest := 0

	for _, s := range servers {
		if s.RAM > highest {
			highest = s.RAM
		}
	}
	return highest

}

func HighestCPU(servers []Server) int {
	highestCpu := 0

	for _, s := range servers {
		if s.CPUCore > highestCpu {
			highestCpu = s.CPUCore
		}
	}
	return highestCpu
}

func (s *Server) Start () {
	if !s.IsRunning {
		s.IsRunning = true
	}
}

func (s *Server) Stop () {
	if s.IsRunning {
		s.IsRunning = false
	}
}

func main() {

	Server1 := Server{"Node-A", 4, 16, true}
	Server2 := Server{"Node-B", 8, 32, false}
	Server3 := Server{"Node-C", 16, 64, true}

	Servers := []Server{Server1, Server2, Server3}

	// for _, s := range Servers {
	// 	fmt.Println("Checking Hostname:", s.Hostname)
	// 	fmt.Println("is running:", s.IsRunning)
	// 	fmt.Println("------------------------------")
		
	// }

	// for i := range Servers {
	// 	fmt.Println("Checking Hostname:", Servers[i].Hostname)
	// 	fmt.Println("is running:", Servers[i].IsRunning)
	// 	fmt.Println("------------------------------")
	// }
	runningCount := CountRunningServers(Servers)
	// fmt.Printf("Number of running servers: %d\n", runningCount)
	// fmt.Println("------------------------------")

	downCount := CountDownServers(Servers)
	// fmt.Printf("Number of down servers: %d\n", downCount)
	// fmt.Println("------------------------------")

	totalRAM := TotalRAM(Servers)
	// fmt.Printf("Total RAM of all servers: %d GB\n", totalRAM)
	// fmt.Println("------------------------------")
	

	totalCPUCore := TotalCPUCore(Servers)
	// fmt.Printf("Total CPU cores of all servers: %d\n", totalCPUCore)
	// fmt.Println("------------------------------")

	// fmt.Println("Highest RAM", HighestRAM(Servers))
	// fmt.Println("Finished checking all servers.")
// Instead of 4 separate Println calls, you could do something like:
	// fmt.Printf("Stats: [Running: %d | Down: %d | Total CPUCore: %dGB | Total RAM: %dGB | Max RAM: %dGB]\n", 
    // runningCount, downCount, totalCPUCore, totalRAM, HighestCPU(Servers), HighestRAM(Servers))

	// Added one more %d for Max CPU and aligned the labels
	fmt.Printf("Stats: [Running: %d | Down: %d | Total CPU: %d | Total RAM: %dGB | Max CPU: %d | Max RAM: %dGB]\n", 
    	runningCount, 
    	downCount, 
    	totalCPUCore, 
    	totalRAM, 
    	HighestCPU(Servers), 
    	HighestRAM(Servers),
	)
		Servers[2].Start()
		fmt.Println("Running Server",runningCount() )
				// Stop()
	fmt.Printf("Stats: [Running: %d | Down: %d | Total CPU: %d | Total RAM: %dGB | Max CPU: %d | Max RAM: %dGB]\n", 
    	runningCount, 
    	downCount, 
    	totalCPUCore, 
    	totalRAM, 
    	HighestCPU(Servers), 
    	HighestRAM(Servers),
	)
}
