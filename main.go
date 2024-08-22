package main

import (
	"TaskTracker/cli"
	"TaskTracker/task"
	"fmt"
	"os"
	"strings"
)

func main() {
	print("Bienvenido a taskTracker")
	print("Escriba 'info' para ver la lista de programas")

	// Se declar la lista de tareas con un array de tareas de 0 elementos
	list := &task.Tasks{
		Stack: make([]task.Task, 0),
	}
	for {
		c, err := cli.Read()
		if err != nil {
			panic(err)
		}

		instructions := strings.Split(c, " ")
		switch instructions[0] {
		case "info":
			cli.Commands()
		case "task-cli":
			cli.TaskCli(instructions, list)

		case "exit":
			fmt.Println("Adios")
			os.Exit(0)

		default:
			println("Comando no reconocido")
		}
	}

}

func print(parameters any) {
	fmt.Println(parameters)
}
