package ejercicios

type Tarea struct {
	Nombre string
	Tiempo int
}

func Ejercicio2(tareas []Tarea) {
	tareasOrdenadas := ordenar(tareas)
	for i := range tareas {
		tareas[i] = tareasOrdenadas[i]
	}
}

func ordenar(tareas []Tarea) []Tarea {
	if len(tareas) == 0 {
		return make([]Tarea, 0)
	}
	corta, pos := buscarMasCorta(tareas)
	restantes := rearmar(tareas, pos)
	return append([]Tarea{corta}, ordenar(restantes)...)
}

func rearmar(tareas []Tarea, pos int) []Tarea {
	slice1 := tareas[:pos]
	slice2 := tareas[pos+1:]
	slice1 = append(slice1, slice2...)
	return slice1
}
func buscarMasCorta(tareas []Tarea) (Tarea, int) {
	var masCorta Tarea
	if len(tareas) == 0 {
		return masCorta, 0
	}
	masCorta = tareas[0]
	indice := 0
	for i, tarea := range tareas {
		if tarea.Tiempo < masCorta.Tiempo {
			masCorta = tarea
			indice = i
		}
	}
	return masCorta, indice
}
