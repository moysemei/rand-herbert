//Package practice reinforces learning in for
package practice

//Plano:
//Primeira func, i muda e n fixo
//Segunda func não fiz o plano pois montei na minha cabeça.
//Terceira func, volta, i e product:

//SumEven sum all even numbers from 1 to n
func SumEven(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			total += i
		}
	}
	return total
}

//CountMultiples how many numbers from 1 to limit are multiples of k
func CountMultiples(limit, k int) int {
	result := 0
	for i := 1; i <= limit; i++ {
		if i%k == 0 {
			result++
		}
	}
	return result
}

//Power returns base to the power of exp without the math package
func Power(base, exp int) int {
	product := 1
	for i := 1; i <= exp; i++ {
		product *= base
	}
	return product
}

//Dificuldades:
//Primeira func resolvi sem hint e sem muita dificuldade.
//Segunda func fiz de primeira. (=O <- surpreso).
//Terceira func: dificuldade em montar a fórmula, abri a primeira hint
//não consegui formular ainda. Eu estava errando em, não colocar um valor que permaneça,
//tipo product + = base * base, tentei também com i := base. Irei abrir hint 2 do Power
//abri a hint 2 e fiz na hora puta merda kk.
//não montei os planos que foram pedidos, pq eu começo a montar e vou tendo ideia e passo
//pra prática de fato.
