package main

import "fmt"

func main() {

	todos := Todos{}
	todos.add("Pray for grace")
	todos.add("Ask the holy spirit for guidiance")
	fmt.Printf("%+v\n\n\n", todos)
	todos.delete(0)
	fmt.Printf("%+v\n", todos)

}
