package ejercicios

var billetes = []int{10000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}

func Cambiar(cantidad int) map[int]int {
	if cantidad == 0 {
		return make(map[int]int)
	}

	billete := mayorCambio(cantidad, billetes)
	resultado := Cambiar(cantidad - billete)
	resultado[billete]++

	return resultado
}

func mayorCambio(cantidad int, billetes []int) int {
	var valor int
	for _, billete := range billetes {
		if billete <= cantidad {
			valor = billete
			break
		}
	}
	return valor
}
