package main

import (
	"fmt"
)

type taskList struct {
	/* Slices: Estructuras de datos similares a los arrays, pero de longitud variable. */
	tasks []*task
}

type task struct {
	nombre string
	descripcion string
	completado bool
}

func (taskList *taskList) agregarTarea(task *task) {
	// Agregar elementos a un slices.
	taskList.tasks = append(taskList.tasks, task)
}

func (taskList *taskList) eliminarTarea(index int) {
	// Eliminar elementos a un slices.
	taskList.tasks = append(taskList.tasks[:index], taskList.tasks[index + 1:]...)
	/* Los ":" en taskList.tasks[index + 1:] aisla el elemento. Los "..." permite ejecutar
	la operación. Sino están marca error de sintaxis. */
}

func (taskList *taskList) imprimirTarea() {
	for _, tarea := range taskList.tasks {
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripción:",tarea.descripcion)
	}
}

func (taskList *taskList) imprimirTareaCompleta() {
	for _, tarea := range taskList.tasks {
		if tarea.completado {
			fmt.Println("Nombre:", tarea.nombre)
			fmt.Println("Descripción:",tarea.descripcion)
		}
	}
}

func (task *task) marcarCompleta() {
	task.completado = true
}

func (task *task) actualizarDescripcion(descripcion string) {
	task.descripcion = descripcion
}

func (task *task) actualizarNombre(nombre string) {
	task.nombre = nombre
}

func main() {
	taskUno := &task {
		nombre: "Completar curso de Go.",
		descripcion: "Completar curso de Go ésta semana.",
	}

	taskDos := &task {
		nombre: "Completar curso de Python.",
		descripcion: "Completar curso de Python ésta semana.",
	}

	taskTres := &task {
		nombre: "Completar curso de Node.",
		descripcion: "Completar curso de Node ésta semana.",
	}

	taskList := &taskList {
		tasks: []*task {
			taskUno,
			taskDos,
		}, // Cuando se abren corchetes, se indica elementos para preinicializar el slice.
	}

	fmt.Println(taskList.tasks[0])

	taskList.agregarTarea(taskTres)
	fmt.Println(len(taskList.tasks))

	taskList.eliminarTarea(1)
	fmt.Println(len(taskList.tasks))

	taskList.imprimirTarea()

	fmt.Printf("%+v\n", taskUno) // Permite dar un formato.
	// %+v indica que se va a mostar una interface, y seguir la estructura key: value.
	taskUno.marcarCompleta()
	fmt.Printf("%+v\n", taskUno)

	for i := 0; i < len(taskList.tasks); i++ {
		fmt.Println("Index:", i, "Nombre:", taskList.tasks[i].nombre)
	}

	for index, task := range taskList.tasks {
		fmt.Println("Index:", index, "Nombre:", task.nombre)
	} /* range crea un iterador */

	for in := 0; in < 10; in++ {
		if in == 5 {
			break // break rompe el ciclo.
		}

		fmt.Println(in)

		// continue rompe explícitamente en el index deseado, y continúa en la próxima posición.
	}

	for in := 0; in < 10; in++ {
		if in == 5 {
			continue // continue rompe explícitamente en el index deseado, y continúa en la próxima posición.
		}

		fmt.Println(in)
	}
}