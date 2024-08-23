package cli

import (
	"TaskTracker/Utils"
	"TaskTracker/task"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// Read  Lee de consola /*
func Read() (string, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	str = strings.Replace(str, "\r\n", "", 1)
	return str, nil

}

func Commands() {
	fmt.Println("Lista de comandos:")
	lista := []string{"task-cli", "exit", "info"}

	for _, v := range lista {
		fmt.Printf("-%v", v)
		fmt.Print(" ")
		//if i == len(lista)-1 {
		//	break
		//}
	}
	fmt.Print("\n")
}

func TaskCli(command []string, t *task.Tasks) {

	// Example: add "Programar Modulo emergente"
	//			update 1 "NTX"
	instruction, args, id := getValues(command)

	if id <= 0 {
		fmt.Println("ID Invalida")
		return
	}
	//Instruction deberia tener add, list etc y args un string

	switch instruction {
	case "add":
		if args == "" {
			fmt.Println("Debe añadir una descripcion ala tarea")
			return
		}
		t.AddTask(args, 0)
	case "update":
		//t.UpdateTask(id, args)

	//case "delete":
	//case "mark-in-progress":
	//case "mark-done":
	case "list":
		t.ListTask()
	//case "info":
	default:
		fmt.Println("Comando no reconocido precisione help para saber mas...")
		return
	}

}

func getValues(consoleOutPut []string) (command, args string, id int) {

	// Ejem. [add "hola mundo"]

	//Verificamos que lleve algo
	if len(consoleOutPut) == 0 {
		fmt.Println("Necesita dar una instruccion")
		return
	}
	instruction := consoleOutPut[0]

	if len(consoleOutPut) > 1 {
		if Utils.EsNumero(consoleOutPut[1]) {
			id, _ = strconv.Atoi(consoleOutPut[1])
			args = strings.Join(consoleOutPut[2:], " ")
		} else {
			args = strings.Join(consoleOutPut[1:], " ")
		}
		fmt.Println(args)
		return instruction, args, id
	}
	return instruction, "", -1
}
