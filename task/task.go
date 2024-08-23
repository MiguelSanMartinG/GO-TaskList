package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Task Es el singular de tareas
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	State       string `json:"estado"`
}

// Tasks contiene un conjunto de tareas
type Tasks struct {
	Stack []Task
}

// AddTask Añade tareas ala lista.
//
//	Recibe: Nombre de la tarea.
func (t *Tasks) AddTask(v string, id int) { // state, id int

	if len(t.Stack) == 0 {
		id = 1
	} else {
		lastSliceElement := len(t.Stack) - 1
		id = t.Stack[lastSliceElement].ID + 1
	}
	//Si la ID es igual

	t.Stack = append(t.Stack, Task{
		ID:          id,
		Description: v,
		State:       "in-progress",
	})
	t.WriteInJSON()
	fmt.Println("Añadido exitosamente id: ", id)
}
func (t *Tasks) ListTask(args string) {

	if args == "done" || args == "in-progress" || args == "todo" {
		fmt.Println("->", args)
		t.ListWithState(args)
	} else {
		t.listAll()
	}

}

func (t Tasks) listAll() {
	fmt.Println("Listando: --------------")
	for _, v := range t.Stack {
		fmt.Printf("ID: %d\t %s\t %s\n", v.ID, v.Description, v.State)
	}
}

func (t Tasks) ListWithState(state string) {
	fmt.Println("Listando: --------------")
	for _, v := range t.Stack {
		if v.State == state {
			fmt.Printf("ID: %d\t %s\t %s\n", v.ID, v.Description, v.State)
		}
	}
}

func (t *Tasks) UpdateTask(id int, description string) {
	if description == "" {
		fmt.Println("Es necesario una descripción de la tarea")
		return
	}
	for i, v := range t.Stack {
		if v.ID == id {
			t.Stack[i].Description = description
			fmt.Println("Descripción cambiada")
			t.WriteInJSON()
			t.ListTask("")
			return
		}
	}
	fmt.Println("ID No encontrada")
}

func (t *Tasks) UpdateTaskState(id int, state string) {
	if state == "" {
		return
	}
	for i, v := range t.Stack {
		if v.ID == id {
			t.Stack[i].State = state
			fmt.Println("", t.Stack)
			t.WriteInJSON()
			t.ListTask("")
			return
		}
	}
	fmt.Println("ID No encontrada")
}

func (t *Tasks) DeleteTask(id int) {

	for i, v := range t.Stack {
		if v.ID == id {
			t.Stack = append(t.Stack[:i], t.Stack[i+1:]...)
			fmt.Println("ID borrada ")
			t.WriteInJSON()
			t.ListTask("")
			return
		}
	}
	fmt.Println("ID No encontrada")

}
func (t *Tasks) ReadJSON() {
	t.ReadJSONWithPath("./Data/task.json")
}
func (t *Tasks) ReadJSONWithPath(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Directorio no encontrado")
		log.Fatalf("Error al leer el archivo: %v", err)
	}
	err = json.Unmarshal(data, &t)
	if err := json.Unmarshal(data, &t); err != nil {
		log.Fatalf("Error al deserializar JSON: No hay coincidencia %v", err)
	}
	fmt.Println("JSON leido de forma correcta")
}
func (t *Tasks) WriteInJSON() {
	data, err := json.MarshalIndent(t, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	//Write
	err = os.WriteFile("./Data/task.json", data, 0644)
	if err != nil {
		log.Fatalf("Error al escribir el archivo: %v", err)
	}
	fmt.Println("Archivo JSON escrito con éxito")

}
