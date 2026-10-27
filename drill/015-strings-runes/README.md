# Drill 015 — Strings: bytes e runes

## Antes de tudo: refazer de memória (~15 min)

Do drill 011: **`func ReverseNumber(n int) int`**, o número com os dígitos invertidos (com `n` >= 0).
- `ReverseNumber(472)` → `274`
- `ReverseNumber(1200)` → `21`
- `ReverseNumber(5)` → `5`

Crie o arquivo **`redo/redo.go`**, com **`package redo`**. Não abra o `digitloop.go` nem o `REVIEW.md` do drill 011. Se travar por 10 minutos, anote no código o que tentou, e aí pode consultar. Para testar: `cd redo` e `go test -v`.

Num dia ruim, pule direto para o drill.

## Hoje

Entender o que é uma string em Go por dentro (bytes), e percorrê-la caractere por caractere (runes), inclusive com acentos e japonês.

## O que é

**Uma string em Go é uma sequência de bytes,** em UTF-8. Letras sem acento ocupam 1 byte; `ç`, `ã` e `é` ocupam 2; um kanji como `日` ocupa 3.

| Expressão | Resultado | Por quê |
|---|---|---|
| `len("go")` | `2` | 2 bytes |
| `len("ação")` | `6` | `a` (1) + `ç` (2) + `ã` (2) + `o` (1) |
| `len("日本")` | `6` | 3 bytes por kanji |
| `"go"[0]` | `103` | indexar devolve um **byte** (o código do `g`), não uma letra |

**Uma `rune` é um caractere** (um *code point* do Unicode). Por baixo, é um número `int32`. Rune se escreve com **aspas simples**: `'a'`, `'ç'`, `'日'`. String se escreve com **aspas duplas**: `"a"`.

**O `for range` percorre runes, não bytes:**

```go
for i, r := range "ação" {
	// i = byte position where the rune starts: 0, 1, 3, 5
	// r = the rune: 'a', 'ç', 'ã', 'o'
}
```

Se você só precisa da rune, ignore a posição: `for _, r := range s`.

**Para transformar uma rune de volta em string:** `string(r)`. (No drill 012, `string(7)` com um `int` estava errado. Com uma `rune`, é exatamente o que você quer: `string('ç')` é `"ç"`.)

## Por que é assim

- **Em Python, `len("ação")` é 4,** porque o Python conta caracteres. Em Go é 6, porque o Go conta bytes, que é o que realmente está guardado. O Go escolheu não esconder isso: `len` é sempre barato (o tamanho já está guardado) e nunca precisa percorrer a string.
- **Em JS, `"ação".length` é 4,** mas um emoji conta 2, porque o JS conta unidades de UTF-16. Cada linguagem conta de um jeito; o Go deixa explícito o que é byte e o que é rune.
- **Strings são imutáveis.** `s[0] = 'x'` não compila. Para "mudar" uma string, você monta uma nova.
- **Isso importa no backend:** nomes de usuário, senhas, mensagens em japonês ou com acento. Cortar uma string no meio de uma rune gera texto quebrado (o famoso `�`).

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package text

// CountSpaces returns how many space characters s has.
func CountSpaces(s string) int {
	count := 0
	for _, r := range s { // r is a rune
		if r == ' ' { // single quotes: a rune literal
			count++
		}
	}
	return count
}

// Initials returns the first rune of each word in s, separated by dots: "Ana Maria" -> "A.M.".
func Initials(s string) string {
	result := ""
	newWord := true
	for _, r := range s {
		if r == ' ' {
			newWord = true
			continue
		}
		if newWord {
			result += string(r) + "." // string(r) turns a rune into a string
			newWord = false
		}
	}
	return result
}
```

## Armadilhas

- **`s[i]` é byte, não letra.** Com acento ou kanji, indexar por posição corta o caractere no meio. Para caracteres, use `for range`.
- **`len(s)` conta bytes.** Para contar caracteres, conte as voltas de um `for range`.
- **Aspas simples e duplas não se misturam.** `r == "a"` não compila (rune contra string). É `r == 'a'`.
- **`string(r)` com `rune` é certo; `string(n)` com `int` é errado.** Para número virar texto, continua sendo `strconv.Itoa`.
- **Montar uma string ao contrário:** `result = string(r) + result` coloca a rune **na frente**; `result += string(r)` coloca **atrás**.

## O problema

Crie o arquivo **`runes.go`** nesta pasta, com **`package runes`**, e implemente:

1. **`func CountVowels(s string) int`**: quantas vogais minúsculas sem acento (`a`, `e`, `i`, `o`, `u`) a string tem.
   - `CountVowels("banana")` → `3`
   - `CountVowels("hello")` → `2`
   - `CountVowels("rhythm")` → `0`
   - `CountVowels("")` → `0`
2. **`func Sizes(s string) (bytes, runes int)`**: quantos bytes e quantos caracteres (runes) a string tem.
   - `Sizes("go")` → `2, 2`
   - `Sizes("ação")` → `6, 4`
   - `Sizes("日本")` → `6, 2`
   - `Sizes("")` → `0, 0`
3. **`func Reverse(s string) string`**: a string de trás para frente, sem quebrar acentos nem kanji.
   - `Reverse("hello")` → `"olleh"`
   - `Reverse("ação")` → `"oãça"`
   - `Reverse("日本語")` → `"語本日"`
   - `Reverse("")` → `""`

O `Reverse` é a parte de pensar.

**Antes de codar**, no `// Plano:`:
- **Uma linha por função:** conta, soma ou monta? Percorre bytes ou runes?
- **A tabela do `Reverse("ação")`,** com as colunas volta, `r` e `result`. A coluna `result` tem que terminar em `"oãça"`.

E no fim do arquivo, as **"Dificuldades"** ("nenhuma", se for o caso).

Se travar, faça a tabela antes de abrir a dica. Depois as dicas, depois o chat.

## Checklist

- [ ] `// Plano:` com uma linha por função e a tabela do `Reverse("ação")`.
- [ ] Comentário do pacote e de cada função.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Se fez o "refazer": `cd redo` e `go test -v` mostra `PASS`.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 015: strings and runes"
  git push
  ```
