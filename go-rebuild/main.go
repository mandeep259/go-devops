package main

import "fmt"

// fmt.Println("Lets go rebuild session 1")

//what is variable declaration and assignment
// A varibale is a container for storing data values. In Go, you can declare a variable using the var keyword followed by the variable name and its type. You can also assign a value to the variable at the time of declaration or later in the code.

//what is data types
//Data types specify the type of data that a variable can hold. In Go, there are several built-in data types, including:
// - int: for integer values
// - float64: for floating-point numbers
// - string: for text
// - bool: for boolean values (true or false)
// Data types help ensure that variables are used correctly and that operations on them are valid. They also allow the Go compiler to optimize memory usage and improve performance.
// In Go, you can declare a variable with a specific type and assign a value to it. For example:
// var age int = 30
// var name string = "Alice"
// var isStudent bool = true

// In Go, you can also use type inference to let the compiler determine the type of a variable based on the value assigned to it. For example:
// age := 30
// name := "Alice"
// isStudent := true

// In Go, you can also declare multiple variables of the same type in a single line. For example:
// var x, y, z int = 1, 2, 3
// var a, b, c string = "Hello", "World", "Go"
// Overall, understanding variable declaration, assignment, and data types is fundamental to programming in Go, as it allows you to write efficient and error-free code.

// Why does Go care about types?
// Go is a statically typed language, which means that every variable must have a specific type. This allows the Go compiler to catch type-related errors at compile time, improving code safety and performance. By knowing the types of variables, the compiler can optimize memory usage and ensure that operations on variables are valid.

// what is struct in Go?
// A struct in Go is a composite data type that groups together variables under a single name. It allows you to create complex data structures that can represent real-world entities. A struct can contain fields of different types, and you can define methods on structs to perform operations on the data they hold. Structs are commonly used to model objects, such as servers, users, or products, in Go programming.
// In Go, you can define a struct using the type keyword followed by the struct name and the fields it contains. For example:
// type ServerInfo struct {
//     ServerName string
//     Cpu        int
//	 Memory     float64
//	 Disk       float64
//	 isOperational bool
// }
// In this example, we have defined a struct called ServerInfo that has five fields: ServerName (a string), Cpu (an integer), Memory (a float64), Disk (a float64), and isOperational (a boolean). You can create instances of this struct and access its fields to store and manipulate server information in your Go programs.
// In Go, you can create an instance of a struct and assign values to its fields. For example:
// Node1 := ServerInfo{
//     ServerName:    "localhost",
//     Cpu:           4,
//     Memory:        1028.0,
//     Disk:          256.0,
//     isOperational: true,
// }
// In this example, we have created an instance of the ServerInfo struct called Node1 and assigned values to its fields. You can access the fields of the struct using dot notation, such as Node1.ServerName or Node1.Cpu, to retrieve or modify the values stored in the struct.
// Overall, structs are a powerful feature in Go that allow you to create complex data structures and organize related data together, making your code more modular and easier to maintain.

//what is method in Go?
// A method in Go is a function that is associated with a specific type, such as a struct. It allows you to define behavior related to that type and perform operations on the data it holds. Methods are defined using a receiver argument, which specifies the type that the method is associated with. This allows you to call the method on instances of that type and access its fields or perform actions based on its state.

// In Go, you can define methods with either value receivers or pointer receivers. Value receivers operate on a copy of the struct, while pointer receivers operate on the original struct, allowing you to modify its fields. Methods provide a way to encapsulate behavior related to a specific type and make your code more organized and reusable. They are an important aspect of Go's object-oriented programming capabilities and allow you to define actions that can be performed on instances of a struct.

//what is the pointer receiver in Go?
// A pointer receiver in Go is a method receiver that allows you to modify the fields of a struct. When you define a method with a pointer receiver, you use the * symbol before the type in the receiver argument. This means that the method operates on a pointer to the struct, allowing you to change the values of its fields. For example:
// func (s *ServerInfo) UpdateStatus(isOperational bool) {
//     s.isOperational = isOperational
// }
// In this example, we have defined a method called UpdateStatus for the ServerInfo struct that takes a pointer receiver (denoted by *ServerInfo). This method allows you to update the isOperational field of the struct by passing a boolean value as an argument. You can call this method on an instance of the ServerInfo struct, such as Node1.UpdateStatus(false), to change the operational status of the server. Using pointer receivers allows you to modify the state of the struct and is often used when you want to update thete that when you use a pointer receiver, you need to pass a pointer to the struct when calling the method, such as (&Node1).UpdateStatus(false), to ensure that the changes are reflected in the original struct instance.Overall, pointer receivers provide a way to modify the state of a struct and are an important aspect of Go's object-oriented programming capabilities, allowing you to define methods that can change the values of a struct's fields.

//Why not just use value receiver?
// Value receivers in Go are used when you want to perform read-only operations on a struct without modifying its state. When you define a method with a value receiver, the method operates on a copy of the struct, which means that any changes made to the fields within the method will not affect the original struct instance. For example:
// func (s ServerInfo) DisplayInfo() {
//     fmt.Printf("Server Name: %s\nCPU Cores: %d\nMemory: %.2f MB\nDisk: %.2f GB\nOperational: %t\n",
//         s.ServerName, s.Cpu, s.Memory, s.Disk, s.isOperational)
// }
// In this example, we have defined a method called DisplayInfo for the ServerInfo struct that takes a value receiver (denoted by ServerInfo). This method displays the information of the server in a formatted manner. You can call this method on an instance of the ServerInfo struct, such as Node1.DisplayInfo(), to print the server information to the console. Using value receivers is suitable when you want to perform read-only operations on the struct without modifying its state. It allows you to access and display the fields of the struct without changing their values, making it a good choice for methods that are meant to provide information about the struct without altering its data. Overall, value receivers are useful for methods that do not need to modify the state of the struct and are often used for displaying or retrieving information from a struct, while pointer receivers are used when you want to modify the state of the struct within a method.

// In Go, you can also define methods on structs to perform operations on the data they hold. A method is a function that is associated with a specific type, such as a struct. You can define a method for a struct by using a receiver argument, which specifies the type that the method is associated with. For example:
// func (s ServerInfo) IsOperational() bool {
//     return s.isOperational
// }
// In this example, we have defined a method called IsOperational for the ServerInfo struct. The method takes a receiver argument of type ServerInfo (denoted by s) and returns a boolean value indicating whether the server is operational or not. You can call this method on an instance of the ServerInfo struct, such as Node1.IsOperational(), to check if the server is operational. Methods allow you to encapsulate behavior related to a specific type and make your code more organized and reusable.
// In Go, you can also define methods with pointer receivers, which allow you to modify the fields of the struct. For example:
// func (s *ServerInfo) UpdateStatus(isOperational bool) {
//     s.isOperational = isOperational
// }
// In this example, we have defined a method called UpdateStatus for the ServerInfo struct that takes a pointer receiver (denoted by *ServerInfo). This method allows you to update the isOperational field of the struct by passing a boolean value as an argument. You can call this method on an instance of the ServerInfo struct, such as Node1.UpdateStatus(false), to change the operational status of the server. Using pointer receivers allows you to modify the state of the struct and is often used when you want to update the fields of a struct within a method.
// Overall, methods with pointer receivers provide a way to modify the state of a struct and are an important aspect of Go's object-oriented programming capabilities.
// In Go, you can also define methods with value receivers, which do not modify the fields of the struct. For example:
// func (s ServerInfo) DisplayInfo() {
//     fmt.Printf("Server Name: %s\nCPU Cores: %d\nMemory: %.2f MB\nDisk: %.2f GB\nOperational: %t\n",
//         s.ServerName, s.Cpu, s.Memory, s.Disk, s.isOperational)
// }
// In this example, we have defined a method called DisplayInfo for the ServerInfo struct that takes a value receiver (denoted by ServerInfo). This method displays the information of the server in a formatted manner. You can call this method on an instance of the ServerInfo struct, such as Node1.DisplayInfo(), to print the server information to the console. Using value receivers is suitable when you want to perform read-only operations on the struct without modifying its state.
// Overall, methods with value receivers allow you to perform operations on a struct without changing its state and are useful for displaying or retrieving information from a struct.

type ServerInfo struct {
	ServerName    string
	Cpu           float64
	Memory        float32
	Disk          float32
	isOperational bool
}

func printBanner() {
	fmt.Printf("====================================\n")
	fmt.Printf("	Server Monitoring Report		 \n")
	fmt.Printf("====================================\n")

}

func (s *ServerInfo) checkCPU(server string) string {
	s.Cpu += 10
}
func main() {

	printBanner()

	// var hostname string = "localhost"
	// var port int = 8080
	// cpuCoreCount := 4
	// var username string = "admin"
	// var isActive bool = true

	// fmt.Println("Server Hostname is ", hostname)
	// fmt.Println("Server Port is ", port)
	// fmt.Println("CPU Core", cpuCoreCount)
	// fmt.Println("Username is ", username)
	// fmt.Println("Is Active? ", isActive)

	Node1 := ServerInfo{
		ServerName:    "localhost",
		Cpu:           4,
		Memory:        1028.0,
		Disk:          256.0,
		isOperational: true,
	}
	fmt.Println("-----------------------------")
	fmt.Println("Server Information is ", Node1)
	fmt.Println("Server Name : ", Node1.ServerName)
	fmt.Println("CPU Core Count : ", Node1.Cpu, "cores")
	fmt.Println("Memory : ", Node1.Memory, "MB")
	fmt.Println("Disk :", Node1.Disk, "GB")
	fmt.Println("Is Operational? ", Node1.isOperational)
	fmt.Println("------------------------------")

	Node2 := ServerInfo{
		ServerName:    "remotehost",
		Cpu:           8,
		Memory:        2096.0,
		Disk:          512.0,
		isOperational: false,
	}
	fmt.Println("Server Information is ", Node2)
	fmt.Println("Server Name : ", Node2.ServerName)
	fmt.Println("CPU Core Count : ", Node2.Cpu, "cores")
	fmt.Println("Memory : ", Node2.Memory, "MB")
	fmt.Println("Disk :", Node2.Disk, "GB")
	fmt.Println("Is Operational? ", Node2.isOperational)

	Node1.checkCPU()
	Node1.checkCPU()
	fmt.Println("After increase the CPU", Node1.Cpu)
}
