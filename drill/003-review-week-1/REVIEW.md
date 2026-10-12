# Review — Drill 003

## Veredito

Entregue no mesmo dia, os quatro testes passaram. Os dois pontos do review anterior (usar `const`, um caminho só para montar texto) já apareceram corrigidos.

## Processo

- **Seu plano mostrou o raciocínio, não só a intenção.** No `SlicesEach` você começou com duas variáveis, percebeu que dava para fazer tudo numa conta só, mudou o plano e anotou que não tinha certeza se ia dar certo. Pensar, testar e ajustar: é esse o ciclo.
- **A constante veio certa**, com `=` e para um valor que não depende da entrada. No drill 002 isso tinha travado.
- **Você seguiu as duas regras de reaproveitamento:** o `SlicesEach` usa o `TotalSlices`, e o `Summary` não faz conta nenhuma sozinho.
- **Você escolheu um caminho só no `Summary`** (`+` com `strconv.Itoa`), sem misturar com `Sprintf`.
- **A divisão inteira foi usada de propósito** no `SlicesEach`, e no `CostEach` você converteu os dois `int` antes da conta.

## Erros e melhorias

**[lógica] "O que eu queria mesmo era converter na saída."** Você escreveu isso no plano do `CostEach`, e vale entender por que não dá:
- O Go confere os tipos **a cada operação**, não só no resultado final. Em `pizzas * pricePerPizza`, o `int` e o `float64` já se encontram na multiplicação, e o compilador recusa ali mesmo, antes de chegar na "saída".
- Mesmo quando os tipos batem, converter no fim pode ser tarde demais. Foi o que aconteceu no drill 002 com `float64(minutes / 60)`: a divisão entre `int` já tinha jogado as casas decimais fora.
- *Como pensar:* converta **no ponto em que os tipos se encontram** pela primeira vez. Foi exatamente o que você acabou fazendo, então agora você sabe o porquê.

**[idiomático] Comentários das funções.** O Go espera que o comentário comece pelo nome da função: `// TotalSlices returns ...`. Os seus começaram com "How many...". O checklist pedia isso, e o `gopls` avisa. Faltou também o comentário do pacote (`// Package party ...`).

**[idiomático] Inglês: "resume" é falso cognato.** Em inglês, *resume* significa "retomar" (ou currículo, *résumé*). "Resumir" é *summarize*. Também: "each **people** get it" → "each **person** gets". São detalhes, mas é o tipo de coisa que aparece em PR de vaga lá fora.

**[idiomático] Variável só para devolver.** `s := TotalSlices(pizzas) / people` seguido de `return s` funciona, mas a variável não acrescenta nada: devolva a expressão direto. Variável vale a pena quando o nome explica algo que a conta sozinha não explica.

**[idiomático] Frase com vários números.** Seu `Summary` ficou com três `strconv.Itoa` e seis `+` numa linha só. Funciona e é idiomático, mas com três números o `fmt.Sprintf` fica mais fácil de ler: você vê a frase inteira de uma vez. A regra prática: um ou dois pedaços, `+`; frase com vários valores, `Sprintf`.

## Solução de referência

```go
// Package party splits pizzas and costs between people.
package party

import "fmt"

const slicesPerPizza = 8

// TotalSlices returns how many slices the given pizzas have.
func TotalSlices(pizzas int) int {
	return pizzas * slicesPerPizza
}

// SlicesEach returns how many whole slices each person gets.
// Integer division on purpose: leftover slices stay in the box.
func SlicesEach(pizzas int, people int) int {
	return TotalSlices(pizzas) / people
}

// CostEach returns how much each person pays.
func CostEach(pizzas int, pricePerPizza float64, people int) float64 {
	return float64(pizzas) * pricePerPizza / float64(people)
}

// Summary returns a sentence like "5 people, 16 slices, 3 slices each".
func Summary(pizzas int, people int) string {
	return fmt.Sprintf("%d people, %d slices, %d slices each",
		people, TotalSlices(pizzas), SlicesEach(pizzas, people))
}
```

Duas coisas que você vai ver bastante:

- **Comentário explicando uma decisão** ("Integer division on purpose"). Bons comentários explicam o *porquê*; o *o quê* o código já mostra.
- **Parâmetros do mesmo tipo podem ser agrupados:** `func SlicesEach(pizzas, people int) int` é o mesmo que `func SlicesEach(pizzas int, people int) int`. As duas formas estão certas.

## Para levar

Converta no ponto em que os tipos se encontram, não no final. E escreva os comentários começando pelo nome da função, em inglês claro.
