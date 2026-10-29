# Review — Drill 016

## Veredito

Entregue no mesmo dia, com o refazer e os três testes passando. O mais valioso deste drill são as suas "Dificuldades": elas mostram exatamente onde o raciocínio travou no `SumNumbers`, e é um ponto que já apareceu antes.

## Processo

- **As "Dificuldades" são um diário de depuração de verdade:** a tentativa errada inteira, copiada como comentário, e por que você fez daquele jeito ("tirei o word, pois pensei que não saberia o que fazer com duas variáveis"). Com isso, não preciso adivinhar onde você travou.
- **Você achou o próprio bug:** "o que me destruiu foi não ter colocado `strconv.Atoi(word)`, mas sim `s`".
- **O `IsPalindrome` foi corrigido pelo raciocínio certo:** "eu precisava normalizar o `s` primeiro e daí sim comparar".
- **O `Normalize` em dois passos,** cada linha com uma transformação, se lê muito bem.
- **Você pôs o `if err != nil` antes do `sum += numb`** na versão final. Na tentativa, o número era somado antes de conferir o erro, e você inverteu a ordem.
- **No plano, você trocou o `TrimSpace` pelo `ReplaceAll`** ("na verdade strings.ReplaceAll"), percebendo que só as pontas não bastavam.
- **"Esse não entendi como faz"** no plano do `SumNumbers`: honesto, e o ponto exato que eu explico abaixo.

## Erros e melhorias

**[lógica] A variável do `for range` é "o item da vez".** Na sua tentativa:

```go
for range strings.Fields(s) {
	numb, err := strconv.Atoi(s) // s is the whole text, every round
```

Você tirou o `word` por não saber o que fazer com ele, e aí o `Atoi` recebeu o texto inteiro em toda volta. É a mesma confusão do drill 009, com outra roupa: lá, o `n` entrava na conta no lugar do `i`; aqui, o `s` entra no lugar do `word`.

| Variável | Papel | Muda a cada volta? |
|---|---|---|
| `s` | o texto inteiro, a coleção | não |
| `word` | o pedaço **desta volta** | sim: `"4"`, depois `"x"`, depois `"5"` |
| `sum` | o acumulador | sim |

*Como pensar:* quando o `for range` te dá uma variável, ela é o item da volta, e é **com ela** que você trabalha dentro do loop. A coleção (`s`, ou `strings.Fields(s)`) só diz o que percorrer. É o mesmo papel do `i` do 009 e da `r` do 015.

E o `SumNumbers("4 x 5")` que você não sabia fazer à mão, rastreado:

| volta | `word` | `Atoi(word)` | o que acontece |
|---|---|---|---|
| 1 | `"4"` | 4, sem erro | `sum` = 4 |
| 2 | `"x"` | erro | `return 0, false`, e a função para aqui |

O "5" nem chega a ser lido. É o padrão busca do drill 012: o primeiro pedaço inválido já decide a resposta.

**[idiomático] Nomes.** `newString` não diz o que a variável guarda; `reversed` diz. `numb` funciona, mas em Go se usa `n`, curto e convencional.

**[idiomático] Comentários.** O `reverse` ficou sem, e os outros precisam de ajuste:
- **`IsPalindrome compare two strings and return if they are palindrome`:** a função não compara duas strings (recebe uma) e devolve `bool`. Fica `IsPalindrome reports whether s reads the same backwards, ignoring case and spaces.`
- **`SumNumbers return the sum of numbers wrote as strings`:** "returns", com "s", e "written", não "wrote". Melhor ainda, dizendo o que é o `bool`: `SumNumbers returns the sum of the space-separated numbers in s. ok is false if any part is not a number.`

## Solução de referência

```go
// Package text cleans and checks text with the strings and strconv packages.
package text

import (
	"strconv"
	"strings"
)

// Normalize returns s in lowercase and without spaces.
func Normalize(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "")
}

// reverse returns s with its runes in reverse order.
func reverse(s string) string {
	result := ""
	for _, r := range s {
		result = string(r) + result
	}
	return result
}

// IsPalindrome reports whether s reads the same backwards, ignoring case and spaces.
func IsPalindrome(s string) bool {
	clean := Normalize(s)
	return clean == reverse(clean)
}

// SumNumbers returns the sum of the space-separated numbers in s.
// ok is false if any part of s is not a number.
func SumNumbers(s string) (sum int, ok bool) {
	for _, word := range strings.Fields(s) { // word is one piece per round
		n, err := strconv.Atoi(word)
		if err != nil {
			return 0, false // check the error before using n
		}
		sum += n
	}
	return sum, true
}
```

Por que assim:

- **O `Normalize` numa linha só** encadeia as duas funções. A sua versão em dois passos é igualmente boa, e talvez mais fácil de ler. As duas são idiomáticas.
- **`clean` em vez de reaproveitar o `s`.** Mudar o parâmetro funciona, mas com um nome novo dá para ver que o texto agora está limpo.
- **O retorno nomeado `(sum int, ok bool)`** deixa o `sum` já existindo em zero, e a assinatura diz o que é cada valor.

## Refazer de memória (FizzBuzz)

Certo de memória, com a ordem certa (o caso "3 e 5" primeiro) e o `strconv.Itoa` no fim, sem cair no `string(n)`. Faltou só o comentário da função.

## Para levar

Num `for range`, a variável que ele te dá é o item da vez, e é com ela que você trabalha dentro do loop. E com `value, err :=`, confira o `err` antes de usar o `value`.
