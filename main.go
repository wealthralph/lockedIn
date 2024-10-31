package main

func main() {

	todos := Todos{}
	todos.add("Pray for grace")
	todos.add("Ask the holy spirit for guidiance")
	todos.toggle(0)
	todos.print()

}
