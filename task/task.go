package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Es el singular de tareas
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	State       string `json:"estado"`
}

// Plural contiene un conjunto de tareas
type Tasks struct {
	Stack []Task
}

/*
*
AddTask Añade tareas ala lista.
Recibe: Nombre de la tarea.
*/
func (t *Tasks) AddTask(v string, id int) { // state, id int

	//Si la ID es igual
	if id < 1 {
		id = len(t.Stack) + 1

	}
	t.Stack = append(t.Stack, Task{
		ID:          id,
		Description: v,
		State:       "in-progress",
	})
	t.writeInJSON()
	fmt.Println("Añadido exitosamente id: ", id)
}

func (t *Tasks) ListTask() {
	fmt.Println("Listando: --------------")
	for _, v := range t.Stack {
		fmt.Printf("ID: %d\t %s\t %s\n", v.ID, v.Description, v.State)
	}
}

func (t *Tasks) writeInJSON() {
	data, err := json.MarshalIndent(t, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	//Write
	err = os.WriteFile("./JSON/output.json", data, 0644)
	if err != nil {
		log.Fatalf("Error al escribir el archivo: %v", err)
	}
	fmt.Println("Archivo JSON escrito con éxito")

}

func (t *Tasks) UpdateTask() {

}
