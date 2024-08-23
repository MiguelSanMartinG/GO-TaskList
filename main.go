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
	initJSON(list)
	list.ReadJSON()
	for {
		c, err := cli.Read()
		if err != nil {
			panic(err)
		}

		// In: String output Out: String slice
		instructions := strings.Split(c, " ")
		switch instructions[0] {
		case "info":
			cli.Commands()
		case "task-cli": //Eliminamos el comando del programa ->
			cli.TaskCli(instructions[1:], list)

		case "exit":
			fmt.Println("Adios")
			os.Exit(0)

		default:
			println("Comando no reconocido")
		}
	}

}

func initJSON(list *task.Tasks) {
	var directorio = "./Data"
	var file = "/task.json"

	_, err := os.Stat(directorio)
	if os.IsNotExist(err) {
		fmt.Println("El directorio no existe.")
		err := os.Mkdir(directorio, 0755)
		if err != nil {
			fmt.Println("Error al crear el directorio:", err)
			return
		} else {
			fmt.Println("Directorio creado")
		}
	}
	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		fmt.Println("El Archivo no existe.")
		file, err := os.Create(directorio + file)
		if err != nil {
			fmt.Println("Error al crear el archivo:", err)
			return
		} else {
			defer file.Close()
			list.WriteInJSON()
			fmt.Println("Directorio creado")

		}
		return
	}
}
