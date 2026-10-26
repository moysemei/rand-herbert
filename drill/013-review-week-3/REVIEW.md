# Review — Drill 013

## Veredito

Entregue, os três testes passam, e quatro pontos de atenção das últimas semanas apareceram resolvidos de uma vez: comentários em todas as funções, nomes que dizem o que guardam, cópia do original antes do loop e as contas à mão no plano.

## Processo

- **As contas à mão voltaram, e certas:** "10 -> 1; 10 != 1, false" e "9875 -> 29 -> 11 -> 2". O caso do 10 é justamente o que derruba quem esquece que o invertido "perde" o zero.
- **Comentários em todas as funções, inclusive nos ajudantes,** com o verbo certo: `reports whether` nas que devolvem `bool`, `returns` nas outras. Isso estava nos pontos de atenção desde o drill 003.
- **Os nomes contam a história:** `original`, `reversed`, `digit`, `count`, `total`. Nada de `f`.
- **`original := n` antes do loop.** A armadilha do drill 011 ("depois do loop, o `n` vale 0") não te pegou.
- **O `CountPrimes` conta (`count++`) em vez de somar,** e começa em 2, que é onde os primos começam.
- **O `DigitalRoot` com `for n >= 10`** e `n = sumDigits(n)` guardando o resultado: as duas armadilhas desse problema (o `>` no lugar do `>=` e o resultado jogado fora) passaram longe.

## Erros e melhorias

Nenhum erro de lógica ou de sintaxe. Dois pedidos, os dois sobre o processo:

**O `d <= n/d` continua sem uma frase sua.** O pedido do review 012 continua valendo: escreva, no plano do drill de hoje, uma frase explicando por que `d <= n/d` funciona e por que você prefere essa forma ao `d*d <= n`. Se você explicar, ele passa a ser seu de verdade.

**As "Dificuldades" sumiram nos dois últimos drills.** Pode ser que não tenha havido nenhuma, e tudo bem. Nesse caso, escreva "nenhuma". Quando a seção não aparece, eu não sei se foi fácil ou se você travou, e é ela que decide se o próximo drill avança ou reforça.

## Solução de referência

```go
// Package puzzles has three classic interview warm-ups with loops.
package puzzles

// isPrime reports whether n is a prime number.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for d := 2; d*d <= n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

// sumDigits returns the sum of the digits of n (n >= 0).
func sumDigits(n int) int {
	total := 0
	for n > 0 {
		total += n % 10
		n /= 10
	}
	return total
}

// CountPrimes returns how many primes there are from 1 to limit, inclusive.
func CountPrimes(limit int) int {
	count := 0
	for i := 2; i <= limit; i++ { // 0 and 1 are never prime
		if isPrime(i) {
			count++
		}
	}
	return count
}

// IsPalindromeNumber reports whether n (n >= 0) reads the same backwards.
func IsPalindromeNumber(n int) bool {
	original := n // n becomes 0 in the loop below
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed == original
}

// DigitalRoot returns the single digit left after summing the digits of n repeatedly.
func DigitalRoot(n int) int {
	for n >= 10 {
		n = sumDigits(n)
	}
	return n
}
```

Por que assim:

- **É praticamente o seu código.** As diferenças são dois comentários dentro das funções e o `d*d <= n`.
- **Para ir além (não precisa decorar):** o `DigitalRoot` tem uma fórmula sem loop. Para `n > 0`, o resultado é `1 + (n-1)%9`; com 9875, `9874 % 9` dá 1, e a resposta é 2. "Dá para fazer sem loop?" é uma pergunta clássica de entrevista depois que você resolve com loop. Ninguém espera que você saiba a fórmula de cabeça, mas vale saber que a pergunta existe.

## Para levar

A semana de loops fechou bem: contador, acumulador, cópia do original, padrão busca e ajudantes reaproveitados, tudo no lugar. Agora o que falta é deixar o seu raciocínio por escrito: as "Dificuldades" e a explicação dos truques.
