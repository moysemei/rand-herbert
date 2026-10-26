# Drill 014 — Funções que devolvem mais de um valor

## Antes de tudo: refazer de memória (~15 min)

Do drill 010: **`func Power(base, exp int) int`**, `base` elevado a `exp` (com `exp` >= 0), sem o pacote `math`.
- `Power(2, 10)` → `1024`
- `Power(3, 4)` → `81`
- `Power(5, 0)` → `1`

Crie o arquivo **`redo/redo.go`**, com **`package redo`**. Não abra o `practice.go` nem o `REVIEW.md` do drill 010. Se travar por 10 minutos, anote no código o que tentou, e aí pode consultar. Para testar: `cd redo` e `go test -v`.

Num dia ruim, pule direto para o drill.

## Hoje

Escrever funções que devolvem dois ou três valores, e usar o "vírgula ok" para avisar quando a resposta não existe.

## O que é

Uma função Go pode devolver vários valores. Os tipos vão entre parênteses depois dos parâmetros, e quem chama recebe tudo separado por vírgula:

```go
func MinMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

small, big := MinMax(7, 3) // small = 3, big = 7
_, big2 := MinMax(1, 9)    // _ ignores the first value
```

**Retornos nomeados.** Você pode dar nome aos valores de saída: `func Split(total int) (meters, centimeters int)`. Os nomes servem de documentação (quem lê a assinatura sabe o que volta) e viram variáveis que já existem dentro da função, começando no valor zero.

**O "vírgula ok".** Quando uma resposta pode não existir, a função devolve o valor **e** um `bool` dizendo se ele vale:

```go
value, ok := ParseGrade("Z")
if !ok {
	// there is no answer: handle it
}
```

Quando o `ok` é `false`, o primeiro valor vem como o valor zero do tipo (0, `""`, `false`), só para ter alguma coisa ali.

**Função como valor (só para reconhecer).** Uma função pode ser guardada numa variável e passada como parâmetro: `double := func(x int) int { return x * 2 }`. Ela aparece de verdade na fase 2; hoje é só para você não estranhar quando vir.

## Por que é assim

- **O Go não tem exceções.** "Deu certo?" é um valor de retorno, e o "vírgula ok" de hoje é o irmão mais novo do `value, err := ...`, a linha mais comum em código Go de verdade (fase 2). Aprender a receber dois valores agora é aprender o formato de todo o resto.
- **Em Python, `return a, b` devolve uma tupla,** que você desempacota. Em Go não existe tupla: são valores separados, cada um com o seu tipo.
- **Em JS, a falha costuma vir como `null`, `undefined` ou uma exceção.** Em Go, ela está na assinatura: quem chama vê o `bool` e é obrigado a recebê-lo (ou a ignorar com `_`, o que fica visível no código).
- **Não dá para ignorar sem querer.** Se a função devolve dois valores, `x := f()` não compila.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package measure

// Split splits a length in centimeters into meters and centimeters.
func Split(total int) (meters, centimeters int) {
	meters = total / 100 // = (not :=): meters already exists
	centimeters = total % 100
	return meters, centimeters
}

// ParseGrade returns the points of a letter grade.
// ok is false when the letter is not a valid grade.
func ParseGrade(letter string) (points int, ok bool) {
	switch letter {
	case "A":
		return 4, true
	case "B":
		return 3, true
	case "C":
		return 2, true
	}
	return 0, false // zero value + false: "there is no answer"
}
```

E usando (dentro de alguma função):

```go
m, cm := Split(254)       // m = 2, cm = 54
p, ok := ParseGrade("Z") // p = 0, ok = false
```

## Armadilhas

- **Receber menos valores do que a função devolve.** `x := SplitDigits(472)` dá `assignment mismatch: 1 variable but SplitDigits returns 3 values`. Receba todos ou ignore com `_`.
- **A ordem importa.** `return max, min` numa função que promete `(min, max)` compila, e o teste falha. Quem chama confia na ordem da assinatura.
- **`:=` com retorno nomeado.** O nome já existe dentro da função. `meters := ...` dá `no new variables on left side of :=`. Use `=`.
- **O valor inicial de um "menor" ou "maior".** Começar com 0 ou 9 é um chute que pode dar errado. Comece com um candidato de verdade (um dos próprios valores) e pense no caso do 0.
- **`return` sem valores** (o *naked return*) funciona com retornos nomeados, mas em função de mais de duas linhas esconde o que está voltando. Aqui, escreva os valores.

## O problema

Crie o arquivo **`multi.go`** nesta pasta, com **`package multi`**, e implemente:

1. **`func SplitDigits(n int) (hundreds, tens, ones int)`**: os três dígitos de um `n` de 100 a 999. É a função de "devolver os três de uma vez" que o review do drill 005 prometeu.
   - `SplitDigits(472)` → `4, 7, 2`
   - `SplitDigits(905)` → `9, 0, 5`
   - `SplitDigits(100)` → `1, 0, 0`
2. **`func Divide(a, b int) (int, bool)`**: a divisão inteira de `a` por `b`. Se `b` for 0, não existe resposta: devolva `0, false`.
   - `Divide(7, 2)` → `3, true`
   - `Divide(6, 3)` → `2, true`
   - `Divide(0, 5)` → `0, true`
   - `Divide(5, 0)` → `0, false`
3. **`func MinMaxDigit(n int) (int, int)`**: o menor e o maior dígito de `n` (com `n` >= 0), nessa ordem.
   - `MinMaxDigit(4729)` → `2, 9`
   - `MinMaxDigit(1000)` → `0, 1`
   - `MinMaxDigit(5555)` → `5, 5`
   - `MinMaxDigit(7)` → `7, 7`
   - `MinMaxDigit(0)` → `0, 0`

O `MinMaxDigit` é a parte de pensar.

**Antes de codar**, no `// Plano:`:
- **Uma linha por função:** o que volta, em que ordem, e quando é o caso "sem resposta".
- **A tabela do `MinMaxDigit(4729)`,** com as colunas volta, `digit`, `min` e `max`. Com que valores `min` e `max` começam? E o que acontece com `n = 0`?
- Se ainda não escreveu: uma frase explicando o `d <= n/d` (pedido do review 013).

E no fim do arquivo, as **"Dificuldades"**. Se não houve nenhuma, escreva "nenhuma".

Se travar, faça a tabela antes de abrir a dica. Depois as dicas, depois o chat.

## Checklist

- [ ] `// Plano:` com uma linha por função e a tabela do `MinMaxDigit(4729)`.
- [ ] Comentário do pacote e de cada função.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Se fez o "refazer": `cd redo` e `go test -v` mostra `PASS`.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 014: multiple returns"
  git push
  ```
