# Review — Drill 012

## Veredito

Entregue, os três testes passam, e é o código mais limpo do tema até aqui: a ordem certa no `FizzBuzz`, o "não achei" depois do loop no `IsPrime` e um bom nome para o candidato no `NextPrime`. Ficaram de fora os comentários e as contas à mão do plano.

## Processo

- **O plano descreve o padrão busca exatamente:** "retornar true se nenhum divisor for encontrado" e "parar quando IsPrime retornar true". As duas armadilhas principais do README (desistir cedo e começar no próprio `n`) não apareceram.
- **O `FizzBuzz` testa primeiro o caso que junta as duas regras,** com retorno antecipado em cada `if`. A ordem da escada, dos drills 006 e 007, já está automática.
- **`candidate := n + 1` deixa o código ler como a regra:** o candidato começa depois do `n` e sobe até ser primo. Bom nome.
- **O `NextPrime` reaproveita o `IsPrime`** em vez de repetir a busca de divisores.
- **Faltaram as contas à mão do `IsPrime(21)` e do `IsPrime(7)`** que o checklist pedia, e também a seção "Dificuldades". Sem elas, eu não sei se foi fácil ou se você travou em algum ponto. É ela que me deixa calibrar o próximo drill.

## Erros e melhorias

**[lógica] O `d <= n/d`: explique para mim no próximo plano.** O README e as dicas usavam `d*d <= n`, e você escreveu `d <= n/d`. As duas formas são equivalentes para números positivos. A sua é a versão que não estoura o `int`: com um `n` gigante, `d*d` pode passar do limite (a armadilha do drill 009), e `n/d` nunca passa. É um detalhe de quem já viu muito código, e por isso vale a regra do review 008: **só entra no seu código o que você consegue explicar.** Se você chegou nisso sozinho, ótimo; se pesquisou ou pediu ajuda, também vale. Só me conte nas "Dificuldades", porque é isso que me diz onde você está.

A ideia por trás da raiz quadrada, para garantir:
- Os divisores vêm em pares: 36 = 2 × 18 = 3 × 12 = 4 × 9 = 6 × 6.
- Em cada par, um dos dois é no máximo 6, que é a raiz de 36.
- Então, se nenhum número até a raiz divide `n`, nenhum depois dela divide também.

**[idiomático] Nenhum comentário.** O pacote e as três funções ficaram sem. É o checklist de todo drill, e ele vai pesar mais a partir de agora: o `go vet` não reclama, mas linters como o `revive` (e qualquer revisor de PR numa vaga) reclamam de função exportada sem comentário. O formato:
- `// Package rules ...`
- `// FizzBuzz returns ...`
- `// IsPrime reports whether n is a prime number.` ("reports whether", porque devolve `bool`)
- `// NextPrime returns the smallest prime greater than n.`

## Solução de referência

```go
// Package rules combines conditions and stops loops as soon as the answer is known.
package rules

import "strconv"

// FizzBuzz returns "FizzBuzz" for multiples of 3 and 5, "Fizz" for multiples of 3,
// "Buzz" for multiples of 5, and n as text otherwise.
func FizzBuzz(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0: // the most specific case first
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

// IsPrime reports whether n is a prime number.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	// Divisors come in pairs (d and n/d), so checking up to the square root is enough.
	for d := 2; d*d <= n; d++ {
		if n%d == 0 {
			return false // found a divisor: stop looking
		}
	}
	return true // the loop ended without finding a divisor
}

// NextPrime returns the smallest prime greater than n.
func NextPrime(n int) int {
	for candidate := n + 1; ; candidate++ {
		if IsPrime(candidate) {
			return candidate
		}
	}
}
```

Por que assim:

- **O `FizzBuzz` com `switch` sem expressão** é a escada de `if` do seu código em outro formato. As duas são idiomáticas; o `switch` deixa as quatro regras alinhadas, uma por linha.
- **O `NextPrime` põe o candidato na própria linha do `for`** (`for candidate := n + 1; ; candidate++`), com a condição vazia no meio. Assim o `candidate` só existe dentro do loop. A sua versão, com a variável antes do loop, faz a mesma coisa.
- **`d*d <= n` na referência** porque se lê mais fácil. A sua `d <= n/d` é a mais segura. Escolher entre as duas é exatamente o tipo de troca que vale explicar numa entrevista.

## Para levar

No padrão busca, o "achei" pode sair de dentro do loop, mas o "não achei" só depois dele. E todo detalhe esperto do seu código precisa de uma frase sua explicando por que ele funciona.
