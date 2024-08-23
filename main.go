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

		fmt.Println("Introdusca su consulta: ")
		// Lectura de inputs
		c, err := cli.Read()
		if err != nil {
			panic(err)
		}

		// In: String output Out: String slice
		instructions := strings.Split(c, " ")
		switch instructions[0] {
		case "info":
			commands()
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

// func initJSON
// inicializa la carpeta y el directorio en el que se trabajara.
func initJSON(list *task.Tasks) {
	var directorio = "./Data"
	var file = "/task.json"

	_, err := os.Stat(directorio)
	if os.IsNotExist(err) {
		fmt.Println("No se encontró directorio, se creara")
		err := os.Mkdir(directorio, 0755)
		if err != nil {
			fmt.Println("Error al crear el directorio:", err)
			return
		} else {
			fmt.Println("Se ha creado el directorio: %s \n", directorio)
		}
	}

	_, err = os.Stat(directorio + file)
	if os.IsNotExist(err) {
		file, err := os.Create(directorio + file)
		if err != nil {
			fmt.Println("Error al crear el archivo:", err)
			return
		} else {
			defer file.Close()
			list.WriteInJSON()
			fmt.Println("No se encontró task, se creara un JSON")
		}
	}
}

// func commands
// Da una lista de los comandos disponibles en el programa.
func commands() {
	fmt.Println("Lista de comandos:")
	lista := []string{"task-cli", "exit", "info"}

	for _, v := range lista {
		fmt.Printf("-%v", v)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}
