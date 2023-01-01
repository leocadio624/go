package expenses

/*
Autor : Eduardo Bernal
Fecha : 31/12/22
Descp : Calcula el promedio de lista
*/
func Average(expns ...float32) float32 {

	return Sum(expns...) / float32(len(expns))

}

/*
Autor : Eduardo Bernal
Fecha : 31/12/22
Descp : Realiza sumatoria de lista
*/
func Sum(expns ...float32) float32 {
	var sum float32
	for _, i := range expns {
		sum += i
	}
	return sum
}

/*
Autor : Eduardo Bernal
Fecha : 31/12/22
Descp : Calcula valor maximo de lista
*/
func Max(expns ...float32) float32 {
	var max float32
	for _, i := range expns {
		if i > max {
			max = i
		}

	}
	return max
}

/*
Autor : Eduardo Bernal
Fecha : 31/12/22
Descp : Calcula valor minimo de lista
*/
func Min(expns ...float32) float32 {

	if len(expns) == 0 {
		return 0
	}
	var min float32 = expns[0]
	for _, i := range expns {
		if i < min {
			min = i
		}

	}
	return min
}
