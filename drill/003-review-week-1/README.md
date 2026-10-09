# Drill 003 — Revisão da semana: a conta da pizza

**Hoje:** sem conceito novo. Você vai juntar tudo o que viu nos drills 001 e 002 num problema só, e uma função vai usar a outra.

## O que é

Revisão do que você já sabe:

- **Função devolve, não imprime.** `return` entrega o valor para quem chamou. (drill 001)
- **Parâmetro é a lacuna** que quem chama preenche. (drill 001)
- **Uma função pode chamar outra**, como o `Shout` usava o `Greet`. (drill 001)
- **Tipos não se misturam sozinhos.** `int` com `float64` só depois de converter. (drill 002)
- **Divisão entre dois `int` descarta a parte decimal.** (drill 002)
- **Constante é um valor fixo**, conhecido quando você escreve o código. (drill 002)
- **`fmt.Sprintf`** monta texto com verbos: `%d` para inteiro, `%s` para string. (drill 002)

## Por que é assim

No drill 002, a divisão inteira era a armadilha. Hoje ela é a **ferramenta**: quando você divide fatias de pizza entre pessoas, ninguém ganha "3,2 fatias". Cada um ganha 3, e o resto sobra na caixa. Escolher entre `int` e `float64` não é só fazer o compilador aceitar; é decidir o que o número representa. Fatias são inteiras. Dinheiro tem centavos.

## Como se usa

Exemplo de outro assunto, só para lembrar a ideia de uma função usar outra:

```go
package garden

const seedsPerPack = 20

// TotalSeeds returns how many seeds come in the given packs.
func TotalSeeds(packs int) int {
	return packs * seedsPerPack
}

// SeedsPerPot splits the seeds evenly between pots.
// Integer division on purpose: half a seed is not a seed.
func SeedsPerPot(packs int, pots int) int {
	return TotalSeeds(packs) / pots
}
```

## Armadilhas

Os "pontos de atenção" que apareceram nos reviews:

- **Não imprima.** Se aparecer um `fmt.Println` dentro de uma dessas funções, alguma coisa está errada.
- **Antes de rodar o teste, faça cada função na cabeça com duas entradas diferentes.**
- **Constante usa `=`, nunca `:=`,** e só para valores que não dependem da entrada.
- **Um caminho só para montar texto:** `Sprintf` com `%d`, **ou** `+` com `strconv.Itoa`. Não os dois juntos.
- **Divisão por zero** (0 pessoas) faria o programa quebrar. Não precisa tratar isso hoje: é assunto do `if`, que vem na semana que vem.

## O problema

Crie o arquivo **`party.go`** nesta pasta, com **`package party`**. Uma pizza tem **8 fatias**: use uma constante.

1. **`TotalSlices(pizzas int) int`**: quantas fatias existem no total.
   - `TotalSlices(3)` → `24`
2. **`SlicesEach(pizzas int, people int) int`**: quantas fatias **inteiras** cada pessoa ganha. O que sobra fica na caixa.
   - `SlicesEach(2, 4)` → `4`
   - `SlicesEach(2, 5)` → `3` (16 fatias, 5 pessoas: 3 para cada, sobra 1)
   - **Regra:** use o `TotalSlices` aqui dentro. Não repita a conta das 8 fatias.
3. **`CostEach(pizzas int, pricePerPizza float64, people int) float64`**: quanto cada pessoa paga.
   - `CostEach(2, 45.0, 4)` → `22.5`
   - `CostEach(3, 40.0, 5)` → `24`
4. **`Summary(pizzas int, people int) string`**: uma frase resumindo a festa.
   - `Summary(2, 5)` → `"5 people, 16 slices, 3 slices each"`
   - **Regra:** use as suas funções das partes 1 e 2. O `Summary` não faz conta nenhuma sozinho.

Repare na diferença entre a parte 2 e a parte 3: as duas dividem por `people`, mas uma devolve `int` e a outra `float64`. Pense no porquê antes de escrever.

Se travar, as dicas estão em `HINTS.md`.

## Checklist

- [ ] `// Plano:` dizendo o que cada função **devolve** e quais funções usam outras.
- [ ] Comentário em inglês acima de cada função, começando pelo nome dela.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos quatro testes.
- [ ] Na raiz do repo:

  ```
  git add .
  git commit -m "drill 003: week 1 review"
  git push
  ```
