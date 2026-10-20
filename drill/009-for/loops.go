package loops

// Plan: Factorial(4) - 4 * 3 * 2 * 1, valor começa com o próprio n
// Para o SumOfMultiples: o primeiro

func SumTo(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

func Factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func SumOfMultiples(limit int) int {
	total := 0
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}

// Dificuldades: Travei legal na primeira func...
// Era pra ser fácil mas.
// Abri a hint 1 da func 1 e resolvi.
// Basicamente eu estava fazendo isso mas meio que total += i + n mas o n era só o meu limitador digamos assim, o teto.
// Abri hint 1 da segunda func mas não resolvi ainda.
// Travado na formúla, primeira que tentei: total *= n * (n-1) depois tentei total = n*(n-1) ambas não deram certo.
// é, tive que abrir a terceira hint da segunda func pra resolver. Assim, o que me ferrou foi que eu tinha certeza que tinha que usar o n novamente, mas não kk.
// não consegui fazer a func 3 nem a pau kk
// usei o gpt :(
