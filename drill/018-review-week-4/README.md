# Drill 018 — Revisão da semana: do problema ao código

## Antes de tudo: refazer de memória (~15 min)

Do drill 015: **`func Reverse(s string) string`**, a string de trás para frente, sem quebrar acentos nem kanji.
- `Reverse("hello")` → `"olleh"`
- `Reverse("ação")` → `"oãça"`
- `Reverse("日本語")` → `"語本日"`

Crie o arquivo **`redo/redo.go`**, com **`package redo`**. Não abra o `runes.go` nem o `REVIEW.md` do drill 015. Se travar por 10 minutos, anote no código o que tentou, e aí pode consultar. Para testar: `cd redo` e `go test -v`.

Num dia ruim, pule direto para o drill.

## Hoje

Sem conceito novo. Hoje o objetivo não é a sintaxe, é o caminho entre "entendi o problema" e "escrevi a função": quatro passos fixos e um cardápio de padrões, aplicados em quatro problemas pequenos.

## O que é

**O método, em quatro passos:**

1. **Exemplo à mão:** entrada e saída, com valores de verdade.
2. **Qual é o padrão?** (Às vezes são dois.)
3. **Passos em português,** como comentários, uma linha por passo.
4. **Cada linha vira Go.** Depois, a tabela de rastreio confere.

**O cardápio de padrões, com o esqueleto de cada um:**

| Padrão | Pergunta | Esqueleto |
|---|---|---|
| **Acumular** | "Quanto dá somando?" | `total := 0` · `for _, v := range xs { total += v }` · `return total` |
| **Contar** | "Quantos são?" | `count := 0` · `for _, v := range xs { if cond { count++ } }` · `return count` |
| **Melhor até agora** | "Qual é o maior, o mais longo?" | `best := ponto de partida` · `for _, v := range xs { if v é melhor { best = v } }` · `return best` |
| **Buscar** | "Existe algum?" | `for _, v := range xs { if cond { return true } }` · `return false` |
| **Montar** | "Qual é a nova lista?" | `var result []T` · `for _, v := range xs { result = append(result, ...) }` · `return result` |
| **Percorrer dígitos** | "O que tem no número?" | `for n > 0 { d := n % 10; ...; n /= 10 }` |
| **Trocar pontas** | "Inverter no lugar?" | `for i := 0; i < len(xs)/2; i++ { j := len(xs)-1-i; trocar xs[i] e xs[j] }` |

Os esqueletos são a estrutura. Os pedaços que faltam (`cond`, "é melhor", o que entra no `append`) saem do exemplo à mão.

## Por que é assim

- **Em entrevista, este é o roteiro esperado.** Os entrevistadores chamam de "think out loud": entender com um exemplo, dizer qual é a abordagem, planejar, codar, testar. Ninguém espera que você escreva a função de primeira. Esperam que você mostre o caminho, e o caminho é este.
- **Programador experiente não inventa cada função do zero:** ele reconhece o padrão e adapta. "Isso é contar com filtro" já entrega metade do código. O que separa júnior de pleno, na maior parte das vezes, é o tamanho desse cardápio.
- **Escrever os passos em português primeiro separa os dois problemas** que hoje você tenta resolver juntos: o que fazer (lógica) e como escrever (sintaxe).

## Como se usa

Exemplo completo do método, em outro problema (não é nenhum dos de hoje):

**Problema:** `Positives(nums []int) []int` devolve só os números positivos, na mesma ordem.

1. **Exemplo:** `[3 -1 0 7]` → `[3 7]`.
2. **Padrão:** a pergunta é "qual é a nova lista?", então é **montar**, com um filtro (um `if` antes do `append`).
3. **Passos em português:**

```go
// 1. start with an empty slice
// 2. for each number:
// 3.     if it is positive, add it to the result
// 4. return the result
```

4. **Cada linha vira Go:**

```go
func Positives(nums []int) []int {
	var result []int            // 1
	for _, v := range nums {    // 2
		if v > 0 {              // 3
			result = append(result, v)
		}
	}
	return result               // 4
}
```

**Rastreando** com `[3 -1 0 7]`: o `v` vale 3 (entra), −1 (não), 0 (não) e 7 (entra), e o `result` termina `[3 7]`.

## Armadilhas

- **Pular o passo 2** e começar pelo `for`. Se você não sabe qual é o padrão, ainda não sabe qual é o esqueleto, e o código vira tentativa.
- **O ponto de partida do "melhor até agora".** Às vezes o primeiro elemento (no `MaxOf`, porque 0 podia vencer errado); às vezes um valor neutro (no texto mais longo, `""` funciona, porque nenhuma palavra tem menos que 0 letras). Pergunte: esse ponto de partida pode vencer errado?
- **Dois padrões juntos.** Quando a função monta **e** pode falhar no meio, é **montar** + **buscar**: o primeiro erro encerra a função com retorno antecipado.
- **A variável do `for range` é o item da vez** (review do drill 016). Dentro do loop, trabalhe com ela, não com a coleção inteira.

## O problema

Crie o arquivo **`patterns.go`** nesta pasta, com **`package patterns`**, e implemente:

1. **`func CountChar(s string, target rune) int`**: quantas vezes a rune `target` aparece em `s`.
   - `CountChar("banana", 'a')` → `3`
   - `CountChar("ação", 'ç')` → `1`
   - `CountChar("go", 'x')` → `0`
2. **`func Contains(nums []int, target int) bool`**: o `target` está no slice?
   - `Contains([]int{4, 8, 15}, 8)` → `true`
   - `Contains([]int{4, 8, 15}, 7)` → `false`
   - `Contains([]int{}, 1)` → `false`
3. **`func LongestWord(s string) string`**: a palavra mais longa (as palavras são separadas por espaços). No empate, a primeira. Sem palavras, `""`. Os testes só têm letras sem acento, então `len(word)` serve.
   - `LongestWord("the quick brown fox")` → `"quick"` (`quick` e `brown` têm 5; vence a primeira)
   - `LongestWord("go is fun")` → `"fun"`
   - `LongestWord("  hello  ")` → `"hello"`
   - `LongestWord("")` → `""`
4. **`func ParseAll(words []string) ([]int, bool)`**: converte cada texto em número. Se algum não for número, devolva `nil, false`.
   - `ParseAll([]string{"1", "2", "3"})` → `[1 2 3], true`
   - `ParseAll([]string{"10", "-4"})` → `[10 -4], true`
   - `ParseAll([]string{"1", "x", "3"})` → `nil, false`
   - `ParseAll([]string{})` → `[], true` (lista vazia é válida)

O `LongestWord` e o `ParseAll` são as partes de pensar.

**Antes de codar, no `// Plano:`, os passos 1 a 3 para cada função:** o exemplo, o nome do padrão (ou dos dois) e os passos em português. Pode ser curto: três ou quatro linhas por função. Hoje esse é o exercício principal, e o código é a consequência.

E no fim do arquivo, as **"Dificuldades"** ("nenhuma", se for o caso). Em especial: em qual passo do método você travou?

## Checklist

- [ ] `// Plano:` com exemplo, padrão e passos em português para as quatro funções.
- [ ] Comentário do pacote e de cada função.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos quatro testes.
- [ ] Se fez o "refazer": `cd redo` e `go test -v` mostra `PASS`.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 018: week 4 review"
  git push
  ```
