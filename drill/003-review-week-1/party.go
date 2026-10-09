package party

import "strconv"

/*
Plan:
 - Para a primeira func, irei tentar multiplicar a constante slicesForOnePizza pelo total de pizzas.
 - Na segunda func, pensei em criar uma nova variável atribuir a ela o valor da quantidade de slices utilizando a func TotalSlices
   e então em uma nova variável f atribuir s/people e dar return f. Mudei isso no meio do caminho, pensando que posso colocar tudo
   na mesma conta e só retornar s. (não sei se irá dar certo).
 - Para a terceira func, preciso descobrir o preço primeiro, ou seja fazer pizza * pricePerPizza, mas o pricePerPizza está em float
   então preciso converter antes de multiplicar, mas o que eu queria mesmo era converter na saída.
*/

const slicesForOnePizza = 8

// How many slices in total.
func TotalSlices(pizzas int) int {
	return slicesForOnePizza * pizzas
}

// How many  slices each people get it.
func SlicesEach(pizzas int, people int) int {
	s := TotalSlices(pizzas) / people
	return s
}

// How many each one will pay.
func CostEach(pizzas int, pricePerPizza float64, people int) float64 {
	p := float64(pizzas) * pricePerPizza / float64(people)
	return p
}

// Returns a phrase that resume the party.
func Summary(pizzas int, people int) string {
	return strconv.Itoa(people) + " people, " + strconv.Itoa(TotalSlices(pizzas)) + " slices, " + strconv.Itoa(SlicesEach(pizzas, people)) + " slices each"
}
