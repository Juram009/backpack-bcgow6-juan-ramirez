package ejercicio_2

func burbuja(list []int) []int {
	var aux int
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if list[i] < list[j] {
				aux = list[i]
				list[i] = list[j]
				list[j] = aux
			}
		}
	}
	return list
}
