# Review — Drill 008

## Veredito

Entregue, os três testes passam, sem nenhum pedido de ajuda. Destaque: o `Price` saiu direto do seu plano. A técnica "do exemplo para o código", que travava no 005, funcionou aqui sem dica.

## Processo

- **O `Price` é o review do 005 aplicado.** No plano: "A primeira hora custa 8 reais e as outras 2 horas custam 5 cada". No código: `firstHourPriceReais + (AdditionalHourPriceReais * (hours - 1))`. As "outras 2 horas" viraram `hours - 1`, e cada número virou o nome de onde ele veio. É exatamente o passo que faltava duas semanas atrás.
- **Você usou retorno antecipado para a tolerância,** com `<=`. O limite do 15 estava certo de primeira.
- **No plano do `BilledHours`, você identificou o tipo de problema** ("precisa arredondar para cima") e testou os dois casos de fronteira, 60 e 61.
- **As constantes levam a unidade no nome** (`minutesPerHour`, `freeMinutes`, `...Reais`). Esse ponto de atenção do drill 004 está resolvido.
- **Você acertou a prioridade dos operadores** em `(hours - 1)`: sem esses parênteses, a multiplicação viria antes e o resultado seria outro.

## Erros e melhorias

**[lógica] Use um truque só depois de conseguir explicar ele.** O seu `BilledHours` é `(minutes + minutesPerHour - 1) / minutesPerHour`. Funciona, e é um truque conhecido (divisão arredondando para cima). Mas ele não estava no README nem nas dicas, e o plano não explica como ele funciona. Se você chegou nele sozinho, ótimo. Se pesquisou, também vale: pesquisar faz parte do trabalho. A regra é uma só: **num código seu, só entra o que você consegue explicar**. Numa entrevista, a pergunta seguinte seria "por que o `- 1`?".

A explicação: somar 59 empurra qualquer sobra para a próxima hora cheia, mas não empurra uma hora que já é exata.

| minutos | + 59 | ÷ 60 (inteira) |
|---|---|---|
| 60 | 119 | 1 |
| 61 | 120 | 2 |
| 1 | 60 | 1 |
| 0 | 59 | 0 |

Com `+ 60` no lugar de `+ 59`, o 60 viraria 120 e daria 2 horas, o que estaria errado. É por isso que existe o `- 1`.

A versão com candidatos (as horas inteiras ou uma a mais, decidido pelo `%`) faz a mesma coisa e se explica sozinha. Ela está na solução de referência. As duas estão certas; o importante é você saber escrever as duas.

**[idiomático] Constante exportada sem querer.** `AdditionalHourPriceReais` começa com maiúscula, então é exportada: quem importar o pacote enxerga ela. As outras três são minúsculas. É a regra do drill 001 (maiúscula = exportado), de novo no drill 007, valendo também para constantes. Aqui a constante é um detalhe interno, então deveria ser minúscula.

**[idiomático] Variável só para devolver.** `hours := ...` seguido de `return hours`, e o mesmo com `price`. É o mesmo ponto do review 003: se o nome não acrescenta nada à expressão, devolva a expressão direto.

**[idiomático] Comentários.** O nome no começo está certo em todos. Falta acertar o verbo e o conteúdo:
- **`FormatDuration shows ...`:** "shows" dá a ideia de imprimir. Use `returns`.
- **`BilledHours calculates minutes in hours`:** descreve uma conversão e não fala do arredondamento, que é o ponto da função. Melhor: `// BilledHours returns how many hours are charged; any started hour counts as a full one.`
- **`Price ... based in hours`:** "based **on**" (segunda vez, depois do 007).
- **Comentário do pacote:** diga o que o pacote faz, não o que uma função devolve. `// Package parking computes how long a car stayed and how much it pays.`

## Solução de referência

```go
// Package parking computes how long a car stayed and how much it pays.
package parking

import "fmt"

const (
	minutesPerHour = 60
	freeMinutes    = 15 // tolerance: stays up to this long are free
	firstHourPrice = 8  // reais
	extraHourPrice = 5  // reais for each started hour after the first
)

// FormatDuration returns minutes as text like "2h 5m".
func FormatDuration(minutes int) string {
	hours := minutes / minutesPerHour
	rest := minutes % minutesPerHour
	return fmt.Sprintf("%dh %dm", hours, rest)
}

// BilledHours returns how many hours are charged for a stay.
// Any started hour counts as a full hour: 60 minutes is 1, 61 is 2.
func BilledHours(minutes int) int {
	fullHours := minutes / minutesPerHour // candidate below
	if minutes%minutesPerHour > 0 {       // some minutes left over?
		return fullHours + 1 // candidate above
	}
	return fullHours
}

// Price returns how much a stay costs, in reais.
func Price(minutes int) int {
	if minutes <= freeMinutes {
		return 0
	}
	return firstHourPrice + (BilledHours(minutes)-1)*extraHourPrice
}
```

Por que assim:

- **A unidade foi para o comentário da constante** (`// reais`), e o nome ficou mais curto. As duas formas são aceitáveis. O seu `...Reais` no nome também é claro.
- **O `BilledHours` com candidatos** deixa a regra visível: quem lê vê "horas inteiras, mais uma se sobrou minuto". O truque do `+ 59` é mais curto, mas pede um comentário para ser entendido.
- **O `Price` chama o `BilledHours` direto dentro da conta.** Uma variável `hours` também seria boa aqui, porque o nome ajuda a ler. É questão de gosto, desde que você saiba por que escolheu.

## Para levar

O caminho do `Price` é o modelo: conta com números no plano, cada número vira um nome, e o código sai disso. E um truque só entra no seu código quando você consegue explicar por que ele funciona.
