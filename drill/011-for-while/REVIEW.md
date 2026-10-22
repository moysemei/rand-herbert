# Review — Drill 011

## Veredito

Entregue no mesmo dia, os três testes passam, e o arquivo já veio formatado pelo `go fmt`. Percorrer os dígitos com `% 10` e `/ 10` saiu. O ponto difícil foi remontar o número no `ReverseNumber`, e é nele que este review se concentra.

## Processo

- **Você caiu no loop infinito e saiu sozinho.** "Em mudar o valor do n, se não caía em loop infinito." É a armadilha número 1 do `for` com condição, e você a reconheceu pelo sintoma.
- **No `ReverseNumber`, você deu nome ao dígito:** `digit := n % 10`, e depois `rev = rev*10 + digit`. Com o nome, a fórmula se lê sozinha. É o jeito que eu escolheria.
- **O `go fmt` funcionou:** os comentários vieram com espaço depois do `//`, começando pelo nome.
- **O plano ficou parcial.** A primeira função tem a linha combinada; na segunda você foi direto ao código, e na terceira faltou a tabela. Justamente na terceira você travou na fórmula, e é para esse momento que a tabela existe.

## Erros e melhorias

**[lógica] Remontar o número: "abrir espaço" a cada volta.** Você disse que a dificuldade foi montar o número de novo depois de pegar o último dígito. Esta é a tabela do `ReverseNumber(472)`:

| volta | `n` no começo | `digit` | `rev*10` | `rev` depois |
|---|---|---|---|---|
| 1 | 472 | 2 | 0 | 2 |
| 2 | 47 | 7 | 20 | 27 |
| 3 | 4 | 4 | 270 | 274 |

O `rev*10` empurra tudo uma casa para a esquerda e deixa um zero no fim; o `+ digit` ocupa esse zero. E repare que você já tinha feito essa conta antes, no drill 005: `ones*100 + tens*10 + hundreds`. Abrindo o loop volta por volta, sai a mesma coisa:

```
((0*10 + 2)*10 + 7)*10 + 4  =  2*100 + 7*10 + 4  =  274
```

O loop é a fórmula do 005 construída um pedaço por volta, e por isso funciona para qualquer número de dígitos.

**[idiomático] O `CountDigits` "ajambrado".** Você achou que ficou improvisado. Ele está **certo**, mas entendo a sensação: para ver que `CountDigits(0)` dá 1, quem lê precisa simular o código (o `count++` do `if`, e depois o loop que não roda). Com retorno antecipado, o caso especial fica explícito:

```go
if n == 0 {
	return 1 // the loop below would not run for 0
}
```

Existe ainda um terceiro jeito, que usa o `break` do README: o loop que roda **pelo menos uma vez**. Assim o zero nem precisa de tratamento especial.

```go
count := 0
for {
	count++
	n /= 10
	if n == 0 {
		break
	}
}
```

Os três funcionam. O retorno antecipado é o mais comum em Go.

**[idiomático] Nome de variável.** No `SumDigits`, o acumulador se chama `f`. Funciona, mas não diz nada; `total` ou `sum` dizem o que ele guarda. O `digit` do `ReverseNumber` mostra que você sabe nomear bem.

**[idiomático] Comentários.** A formatação está certa agora. Falta o inglês:
- `SumDigits explit the n in section of each number and sum them` → `SumDigits returns the sum of the digits of n.` ("explit" não existe; seria "split", mas o comentário fica melhor dizendo o que a função devolve.)
- `CountDigits reports how many digits n have` → `CountDigits returns how many digits n has`: "returns", porque devolve um `int`, e "has" na terceira pessoa.
- `Package digitloop for-while study purpose` está sem verbo → `Package digitloop walks through the digits of a number.`

## Solução de referência

```go
// Package digitloop walks through the digits of a number.
package digitloop

// SumDigits returns the sum of the digits of n (n >= 0).
func SumDigits(n int) int {
	total := 0
	for n > 0 {
		total += n % 10 // take the last digit
		n /= 10         // drop it
	}
	return total
}

// CountDigits returns how many digits n has (n >= 0). Zero has one digit.
func CountDigits(n int) int {
	if n == 0 {
		return 1 // the loop below would not run for 0
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

// ReverseNumber returns n with its digits in reverse order: 472 -> 274.
func ReverseNumber(n int) int {
	rev := 0
	for n > 0 {
		digit := n % 10
		rev = rev*10 + digit // shift rev one place left, put digit at the end
		n /= 10
	}
	return rev
}
```

Por que assim:

- **As três funções têm o mesmo esqueleto:** `for n > 0 { ...; n /= 10 }`. Muda só o que se faz com o dígito: somar, contar ou encaixar no fim. Quando você reconhece o esqueleto, metade do problema já está resolvida.
- **O comentário do `CountDigits` documenta o caso especial** ("Zero has one digit"). Quem chama a função não precisa ler o código para saber.

## Para levar

Percorrer um número é sempre "pega o último, tira o último", até sobrar 0. Para remontar, cada volta abre espaço (`*10`) e encaixa o dígito (`+ digit`).
