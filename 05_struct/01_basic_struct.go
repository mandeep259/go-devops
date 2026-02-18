package main

import "fmt"

func printBanner() {
	fmt.Println("====================================")
	fmt.Println("	Server Monitoring Report		 ")
	fmt.Println("====================================")
}

type server struct {
	Hostname string
	CPU      float64
	Memory   float64
	Disk     float64
}

type container struct {
	Name    string
	Image   string
	CPU     float64
	Memory  float64
	Running bool
}

type node struct {
	Name string
	IP   string
	CPU  float64
}

func main() {

	server1 := server{
		Hostname: "k8s-Node-01",
		CPU:      75.5,
		Memory:   85.0,
		Disk:     60.0,
	}

	server2 := server{
		Hostname: "k8s-Node-02",
		CPU:      60.0,
		Memory:   70.0,
		Disk:     50.0,
	}

	printBanner()

	fmt.Println(server1.Hostname)
	fmt.Println(server1.CPU)
	fmt.Println(server1.Memory)
	fmt.Println(server1.Disk)

	fmt.Println(server2.Hostname)
	fmt.Println(server2.CPU)
	fmt.Println(server2.Memory)
	fmt.Println(server2.Disk)

	// container1 := container{
	// 	Name:    "nginx-container",
	// 	Image:   "nginx:latest",
	// 	CPU:     0.5,
	// 	Memory:  256.0,
	// 	Running: true,
	// }
	// fmt.Println(container1.Name)
	// fmt.Println(container1.Image)
	// fmt.Println(container1.CPU)
	// fmt.Println(container1.Memory)
	// fmt.Println(container1.Running)

	// node1 := node{
	// 	Name: "k8s-Node-01",
	// 	IP:   "192.168.1.10",
	// 	CPU:  75.5,
	// }
	// fmt.Println(node1.Name)
	// fmt.Println(node1.CPU)

	// node2 := node{
	// 	Name: "k8s-Node-02",
	// 	IP:   "192.168.1.11",
	// 	CPU:  60.0,
	// }
	// fmt.Println(node2.Name)
	// fmt.Println(node2.CPU)
}
