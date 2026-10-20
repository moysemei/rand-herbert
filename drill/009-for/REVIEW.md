# Review — Drill 009

## Veredito

Entregue, os três testes passam, mas o tema ainda não firmou: dica 1 no `SumTo`, dica 3 no `Factorial`, e o `SumOfMultiples` veio do GPT, como você mesmo anotou. Não é motivo para culpa: loop é onde quase todo iniciante tropeça. Mas o drill 010 vai ser reforço de `for` antes de avançar.

## Processo

- **Você achou a raiz do problema sozinho.** "O n era só o meu limitador, o teto" e "eu tinha certeza que tinha que usar o n novamente". É exatamente esse o conceito que faltou, e está explicado abaixo. Quem sabe nomear a própria confusão já está na metade do caminho.
- **As "Dificuldades" foram honestas e precisas:** o que você tentou (`total += i + n`, `total *= n * (n-1)`), quais dicas abriu e onde usou ajuda de fora. É isso que me deixa calibrar o próximo drill. Continue assim.
- **O plano ficou pela metade** ("Para o SumOfMultiples: o primeiro") e sem a tabelinha que o checklist pedia. Neste tema, a tabela é a ferramenta principal, como você vai ver abaixo.
- **Nenhum comentário de novo** (como no 006). Fica registrado.

## Sobre o GPT

Usar não é proibido, e você ter anotado vale muito. O problema é outro: a resposta pronta pula exatamente a parte que te treina, a passagem de "entendo o problema" para "escrevo o loop", que é a que você quer dominar para as entrevistas. Da próxima vez que travar, a ordem é:

1. Dicas, um nível de cada vez.
2. Uma mensagem para mim aqui no chat. Eu respondo com pergunta ou dica, nunca com a resposta.
3. Se for usar o GPT, peça "me dê uma dica sem código" em vez da solução.

## Erros e melhorias

**[lógica] Cada variável do loop tem um papel.** As suas tentativas (`total += i + n`, `total *= n * (n-1)`) colocavam o `n` dentro da conta. Provavelmente porque, no papel, o fatorial de 4 se escreve "4 × 3 × 2 × 1", e você enxergou ali "n × (n-1) × ...". Mas o loop não faz a conta inteira de uma vez: ele faz **um pedaço por volta**. Cada variável tem um papel diferente:

| Variável | Papel | Muda a cada volta? |
|---|---|---|
| `i` | o número **da vez** | sim: 1, 2, 3... |
| `n` | o teto: diz quando parar | não, nunca |
| `result` / `total` | o acumulador: guarda o que já foi feito | sim, cresce a cada volta |

A técnica do 005 vale aqui também: no exemplo "4 × 3 × 2 × 1", pergunte de onde vem cada número. Nenhum deles é "o `n` de novo": cada um é o `i` de uma volta. Veja a tabela de rastreio do `Factorial(4)`:

| volta | `i` | `result` antes | `result *= i` |
|---|---|---|---|
| 1 | 2 | 1 | 2 |
| 2 | 3 | 2 | 6 |
| 3 | 4 | 6 | 24 |
| — | 5 | | `5 <= 4` é falso, e o loop acaba |

E a sua intuição do plano, "começa com o próprio n", também dá um loop válido: é só contar **para baixo**, com `i--`. O `i` passa por 4, 3, 2, exatamente como você escreveu no papel. Está na solução de referência.

**[lógica] O `SumOfMultiples` tem três camadas.** Como ele veio do GPT, vale entender cada linha, porque o drill 010 volta nesse formato.
1. **O `for` passa por todos os candidatos:** de 1 até `limit - 1` (`i < limit`, porque são os **menores** que `limit`).
2. **O `if` filtra:** só deixa passar quem é múltiplo de 3 ou de 5.
3. **O acumulador soma** quem passou.

Rastreando `SumOfMultiples(10)`:

| `i` | `i%3 == 0 \|\| i%5 == 0` | `total` |
|---|---|---|
| 1 | false | 0 |
| 2 | false | 0 |
| 3 | true | 3 |
| 4 | false | 3 |
| 5 | true | 8 |
| 6 | true | 14 |
| 7 | false | 14 |
| 8 | false | 14 |
| 9 | true | 23 |

O `if` dentro do loop é o mesmo `if` do drill 006, e o `%` é o do 004. A peça nova é só juntar tudo: **o loop gera os números, o `if` escolhe, o acumulador guarda.**

**[idiomático] Comentários.** O pacote e as três funções ficaram sem. O formato é o de sempre: `// SumTo returns 1 + 2 + ... + n.`

## Solução de referência

```go
// Package loops repeats calculations with for.
package loops

// SumTo returns 1 + 2 + ... + n.
func SumTo(n int) int {
	total := 0
	for i := 1; i <= n; i++ { // i is the number of this round: 1, 2, ..., n
		total += i
	}
	return total
}

// Factorial returns n * (n-1) * ... * 1. Factorial(0) is 1.
func Factorial(n int) int {
	result := 1              // multiplying by 1 changes nothing
	for i := n; i > 1; i-- { // counts down: n, n-1, ..., 2
		result *= i
	}
	return result
}

// SumOfMultiples returns the sum of the numbers below limit
// that are multiples of 3 or 5.
func SumOfMultiples(limit int) int {
	total := 0
	for i := 1; i < limit; i++ { // every candidate below limit
		if i%3 == 0 || i%5 == 0 { // keep only the multiples
			total += i
		}
	}
	return total
}
```

Por que assim:

- **O `Factorial` contando para baixo** segue a sua intuição do plano. Contando para cima, como no seu código, também está certo. Com `n = 0` ou `n = 1`, o loop não roda nenhuma vez e o 1 inicial já é a resposta.
- **Os comentários nas linhas do `for` dizem o papel do `i`.** Num código de verdade não precisaria de tanto, mas enquanto o conceito está firmando, escrever o papel de cada variável ajuda.

## Para levar

Num loop, cada variável tem um papel: o `i` é o número da vez (muda a cada volta), o `n` é só o teto (não muda) e o acumulador guarda o resultado (nasce antes, cresce dentro, é devolvido depois). Quando travar, faça a tabela de rastreio: uma linha por volta.
