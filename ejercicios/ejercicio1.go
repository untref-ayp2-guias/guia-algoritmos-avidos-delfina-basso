package ejercicios

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

func SelectorNoRecursivo(actividades []Actividad) []Actividad {
	actividades_seleccionadas := make([]Actividad, 0)
	actividad_actual := actividades[0]
	actividades_seleccionadas = append(actividades_seleccionadas, actividad_actual)
	for _, actividad := range actividades {
		if actividad.Inicio >= actividad_actual.Fin {
			actividades_seleccionadas = append(actividades_seleccionadas, actividad)
			actividad_actual = actividad
		}
	}
	return actividades_seleccionadas
}

func SelectorRecursivo(actividades []Actividad) []Actividad {

	actividadesElegidas := make([]Actividad, 0)
	actividadesElegidas = append(actividadesElegidas, actividades[0])

	aux(actividades[1:], &actividadesElegidas)

	return actividadesElegidas
}

func aux(actividades []Actividad, actsSeleccionadas *[]Actividad) {
	if len(actividades) == 0 {
		return
	}

	ultima := (*actsSeleccionadas)[len(*actsSeleccionadas)-1]
	if actividades[0].Inicio >= ultima.Fin {
		*actsSeleccionadas = append(*actsSeleccionadas, actividades[0])
	}
	aux(actividades[1:], actsSeleccionadas)
}
