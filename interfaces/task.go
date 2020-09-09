package main

import "fmt"

type task struct {
	nombre string
	descripcion string
	completado bool
}

type taskList struct {
	tasks []*task
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

func (taskList *taskList) agregarTarea(task *task) {
	taskList.tasks = append(taskList.tasks, task)
}

func (taskList *taskList) eliminarTarea(index int) {
	taskList.tasks = append(taskList.tasks[:index], taskList.tasks[index + 1:]...)
}

func (taskList *taskList) imprimirTarea() {
	for _, tarea := range taskList.tasks {
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripción:", tarea.descripcion)
	}
}

func (taskList *taskList) imprimirTareaCompleta() {
	for _, tarea := range taskList.tasks {
		if tarea.completado {
			fmt.Println("Nombre:", tarea.nombre)
			fmt.Println("Descripción:", tarea.descripcion)
		}
	}
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

	taskCuatro := &task {
		nombre: "Completar curso de Java.",
		descripcion: "Completar curso de Java ésta semana.",
	}

	taskCinco := &task {
		nombre: "Completar curso de C#.",
		descripcion: "Completar curso de C# ésta semana.",
	}

	taskListUno := &taskList {
		tasks: []*task {
			taskUno,
			taskDos,
		},
	}

	taskListDos := &taskList {
		tasks: []*task {
			taskCuatro,
			taskCinco,
		},
	}

	mapTasks := make(map[string]*taskList) // Se puede usar cualquier tipo de dato.

	mapTasks["Nicole"] = taskListUno
	mapTasks["Cirilla"] = taskListDos

	/* fmt.Println(taskListUno.tasks[0])

	taskListUno.agregarTarea(taskTres)
	fmt.Println(len(taskListUno.tasks))

	taskListUno.eliminarTarea(1)
	fmt.Println(len(taskListUno.tasks))

	taskListUno.imprimirTarea()

	fmt.Printf("%+v\n", taskUno)

	taskUno.marcarCompleta()
	fmt.Printf("%+v\n", taskUno) */
	taskListUno.agregarTarea(taskTres)

	fmt.Sprintln("Tareas de Nicole")
	mapTasks["Nicole"].imprimirTarea()

	fmt.Sprintln("Tareas de Cirilla")
	mapTasks["Cirilla"].imprimirTarea()
}