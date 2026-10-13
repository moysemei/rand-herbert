# Review — Drill 004

## Veredito

Entregue, os três testes passam. `IsEven` e `IsLeapYear` já estavam verdes no primeiro teste que você rodou. O `SecondsToClock` levou quatro rodadas, e o motivo é um só: está em "Erros e melhorias".

## Processo

- **O ponto de atenção do 003 virou hábito.** Os três comentários começam pelo nome da função.
- **A seção "Dificuldades" no fim do arquivo** é exatamente o que um revisor precisa: você escreveu o que não entrou na cabeça (de onde vêm as horas, os minutos e os segundos que vão no `Sprintf`). Ela decidiu o drill de amanhã.
- **Você pediu ajuda com evidência**: o erro do nvim e a saída do `go test`, não só "não funciona". É assim que se pede ajuda num time.
- **No `IsEven`, você comparou com `0`**, e por isso o caso do `-3` não te pegou.
- **O plano ficou sem o cálculo à mão** de 3665 e de 1900 que o checklist pedia. No `IsLeapYear`, o plano dizia `year % 4 == 0 && year % 100 == 0 && year % 400 == 0`. Com 2024: `2024 % 100` é `24`, então essa expressão dá `false`, e 2024 é bissexto. Fazer um exemplo no papel teria mostrado isso antes de codar. O código final ficou certo, então algo te fez mudar de ideia. Vale anotar no plano o que foi.

## Erros e melhorias

**[sintaxe] Faltou o `func`.** A mensagem "expected declaration, found IsEven" apareceu porque, provavelmente, você copiou a assinatura da seção "O problema" do README, que lista só `IsEven(n int) bool`. No nível de cima do arquivo, tudo em Go começa por uma palavra reservada: `func`, `const`, `var`, `type` ou `import`. Quando aparecer "expected declaration" (no `go build`, a mensagem é "non-declaration statement outside function body"), procure a palavra que falta no começo da linha.

**[lógica] Valores no `Sprintf` antes de existirem.** Na primeira versão, os slots de horas e de segundos receberam o `total`. Isso responde à sua dúvida das "Dificuldades": o `Sprintf` **não calcula nada**, só coloca valores prontos nos slots, na ordem. Então o trabalho se divide em duas etapas:
1. **Calcular:** uma variável para cada pedaço, nas linhas de cima.
2. **Formatar:** o `Sprintf` recebe essas variáveis na ordem dos `%d`.

Se o pedaço não existe numa variável (ou numa expressão), não há o que passar para o slot.

**[lógica] Linha copiada três vezes e segundos com `% oneHour`.** Primeiro as três variáveis tinham a mesma conta, por isso todo teste mostrava os três números iguais (`1h 1m 1s`). Depois, os segundos ficaram com `total % oneHour`, que dá `65s` para 3665. Provavelmente você tratou a conta dos minutos como "a fórmula do relógio" e foi adaptando: copiou e apagou um pedaço. *Como pensar:* cada valor da saída responde a **uma pergunta própria**, e a pergunta escolhe o operador.

| Valor | Pergunta | Conta |
|---|---|---|
| horas | Quantas horas inteiras cabem no total? | `total / 3600` |
| minutos | No que sobrou das horas, quantos minutos inteiros cabem? | `total % 3600 / 60` |
| segundos | O que sobra depois de tirar todos os minutos inteiros? | `total % 60` |

"Quantos inteiros cabem?" é `/`. "O que sobra?" é `%`.

E um teste rápido para achar a linha errada: **confira o intervalo do resultado.** Minutos e segundos ficam sempre entre 0 e 59. Se aparece `65s`, a linha dos segundos deixou passar alguma coisa. Pensar no intervalo possível de um valor também é um hábito que entrevistador gosta de ver.

**[idiomático] Nome das constantes.** `oneHour = 3600` deixa a pergunta "3600 o quê?". `secondsPerHour` e `secondsPerMinute` levam a unidade no nome. E escrever `secondsPerHour = 60 * secondsPerMinute` mostra de onde vem o 3600, em vez de pedir que o leitor confie.

**[idiomático] Ordem das variáveis.** Você calculou `minutes` antes de `hours`. Funciona, mas na ordem da saída (horas, minutos, segundos) o código se lê de cima para baixo igual ao texto que ele produz.

**[idiomático] Comentários.** O nome no começo já está certo. Faltam três detalhes:
- **Comentário do pacote:** `// Package numbers ...`, logo acima do `package`.
- **Função que devolve `bool`:** a convenção da biblioteca padrão é "reports whether" (`// IsEven reports whether n is even.`). O "or not" no fim fica sobrando.
- **Inglês:** "**an** total" → "**a** total"; "converts ... **in** hours" → "converts ... **into** hours".

## Solução de referência

```go
// Package numbers has small functions built with arithmetic,
// comparison and logical operators.
package numbers

import "fmt"

const (
	secondsPerMinute = 60
	secondsPerHour   = 60 * secondsPerMinute
)

// IsEven reports whether n is even.
// Comparing with 0 also works for negatives: in Go, -3 % 2 is -1.
func IsEven(n int) bool {
	return n%2 == 0
}

// SecondsToClock converts a total of seconds into "Xh Ym Zs".
func SecondsToClock(total int) string {
	hours := total / secondsPerHour                      // how many whole hours fit?
	minutes := total % secondsPerHour / secondsPerMinute // whole minutes in what is left
	seconds := total % secondsPerMinute                  // what is left after the whole minutes
	return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
}

// IsLeapYear reports whether year is a leap year:
// divisible by 4, except centuries, unless divisible by 400.
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
```

Por que assim:

- **Constantes com unidade no nome**, e o 3600 derivado do 60.
- **Uma pergunta por linha, na ordem da saída**, com um comentário curto dizendo qual é a pergunta.
- **O `Sprintf` só formata.** Toda a conta acontece antes.
- **Parênteses no `IsLeapYear`**, mesmo o `&&` já tendo prioridade sobre o `||`: quem lê vê os dois casos sem precisar lembrar a regra.

## Para levar

Quebrar um número é fazer perguntas: "quantos inteiros cabem?" é `/`, "o que sobra?" é `%`. Calcule cada pedaço numa variável e só depois monte o texto, porque o `Sprintf` só arruma o que já está pronto.
