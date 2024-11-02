package main

func main() {

	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	todos.add("Pray for grace")
	todos.add("Ask the holy spirit for guidiance")
	todos.toggle(0)
	todos.print()

}
