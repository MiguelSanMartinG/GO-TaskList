package cli

import (
	"TaskTracker/task"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

/*
*
GetInput Lee la consola y regresa un string y un error
*/
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
	// command va a taer el comando en pedazos
	// Casos va a llegar TaskCli add "Argumentos"
	// O task-cli list sin argumtos

	//fmt.Println("47Actualmente", t, "Entrando", command)
	//Verificamos que lleve algo
	if len(command) < 2 {
		fmt.Println("Necesita dar una instruccion")
		return
	}

	instruction := command[1]
	// Si llegamos aqui el comando polomenos tiene una instruccion  ejem [add]

	// Lee la segunda linea
	// Lee los argumentos y los une
	args := ""
	if len(command) > 2 {
		body := command[2:]
		args = strings.Join(body, " ")
		fmt.Println(args)
		//fmt.Println("Linea 49", args)
	}

	//fmt.Println("L 66", t)

	switch instruction {
	case "add":
		t.AddTask(args)
	//case "delete":
	//case "mark-in-progress":
	//case "mark-done":
	case "list":
		t.ListTask()
	//case "info":
	default:
		fmt.Println("add, delete, list")
		return
	}

}
