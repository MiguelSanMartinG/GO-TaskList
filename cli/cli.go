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

// TaskCli funcion que evalua la instruccion y hace llos procedimientos.
func TaskCli(command []string, t *task.Tasks) {

	// instruction: add, update, delete, mark-in-progress, mark-done, list
	// id: La id de la tarea es usado solo en update, delete, mak, y list
	// args: suelen ser el nombre de la tarea.
	instruction, args, id := getValues(command)

	//Instruction debería tener add, list etc y args un string
	switch instruction {
	case "add":
		if args == "" {
			fmt.Println("Debe añadir una descripcion ala tarea")
			return
		}
		t.AddTask(args, -1)
	case "update":
		if id <= 0 {
			fmt.Println("ID Invalida")
			return
		}
		t.UpdateTask(id, args)
	case "delete":
		if id <= 0 {
			fmt.Println("ID Invalida")
			return
		}
		t.DeleteTask(id)
	case "mark-in-progress":
		t.UpdateTaskState(id, "in-progress")
	case "mark-done":
		t.UpdateTaskState(id, "done")
	case "list":
		t.ListTask(args)
	default:
		fmt.Println("Comando no reconocido")
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
