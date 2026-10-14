# Drill 006 — if / else

**Hoje:** escolher entre caminhos com `if`, `else if` e `else`.

## O que é

O `if` executa um bloco só quando a condição é `true`. O `else` é o caminho para quando ela é `false`. O `else if` encadeia novas perguntas: o Go testa de cima para baixo e executa **só o primeiro** bloco cuja condição for verdadeira.

```go
if condition {
	// runs when condition is true
} else if otherCondition {
	// runs when condition is false and otherCondition is true
} else {
	// runs when nothing above was true
}
```

**`if` com inicialização:** dá para criar uma variável antes da condição, separada por `;`. Essa variável só existe dentro do `if` e dos `else` dele.

```go
if change := paid - price; change > 0 {
	return change
}
// change does not exist here
```

Você vai usar muito essa forma mais para frente, com erros: `if err := save(); err != nil { ... }`.

## Por que é assim

- **Sem parênteses, com chaves obrigatórias.** É `if x > 0 {`: nada de `if x > 0:` como no Python, e nada de `if (x > 0) doSomething();` sem chaves como no JS. As chaves são obrigatórias mesmo para uma linha só. Isso elimina um tipo de bug: o famoso "goto fail" da Apple, de 2014, foi uma linha a mais embaixo de um `if` sem chaves, que rodava sempre.
- **Não existe operador ternário.** O JS tem `a > b ? a : b` e o Python tem `a if a > b else b`. O Go escolheu não ter nenhum dos dois: um `if` de três linhas é mais fácil de ler do que ternários aninhados.
- **Retorno antecipado.** Quando um ramo termina em `return`, o Go prefere não usar `else`: você trata o caso especial, retorna, e o caminho principal continua embaixo, sem indentação.

```go
// Works, but Go code rarely looks like this.
func Greeting(hour int) string {
	if hour < 12 {
		return "good morning"
	} else {
		return "good afternoon"
	}
}

// Idiomatic: the if handles one case and returns; the rest goes below.
func Greeting(hour int) string {
	if hour < 12 {
		return "good morning"
	}
	return "good afternoon"
}
```

## Como se usa

Exemplos de outro assunto (não são o problema de hoje):

```go
package weather

// Label describes a temperature in Celsius.
func Label(celsius int) string {
	if celsius < 0 {
		return "freezing"
	}
	if celsius < 15 {
		return "cold" // only reached when celsius >= 0
	}
	if celsius < 25 {
		return "mild" // only reached when celsius >= 15
	}
	return "hot"
}

// Advice says what to wear. Here the branches don't return,
// so else if and else choose which value goes into clothes.
func Advice(celsius int, raining bool) string {
	var clothes string // zero value: ""
	if celsius < 15 {
		clothes = "a coat"
	} else if celsius < 25 {
		clothes = "a sweater"
	} else {
		clothes = "a t-shirt"
	}
	if raining { // a separate question, not part of the chain above
		clothes = clothes + " and an umbrella"
	}
	return clothes
}
```

Repare na ordem do `Label`: o `celsius < 15` só é testado quando o `celsius < 0` já deu `false`. Por isso ele significa "de 0 a 14". Se você invertesse as duas linhas, -5 viraria "cold".

## Armadilhas

- **O `else` fica na mesma linha do `}`:** `} else {`. Se ele for para a linha de baixo, dá erro de sintaxe, porque o Go coloca um `;` automático no fim da linha do `}`.
- **A ordem importa.** A primeira condição verdadeira vence, e as de baixo nem são testadas. Coloque os casos mais específicos primeiro.
- **`missing return`.** Se a função devolve algo, todo caminho precisa terminar em `return`. Um `if` com `return` dentro, sem nada depois dele, não basta.
- **Escopo.** Uma variável criada dentro de um bloco `{ }`, ou no `if` com inicialização, não existe depois do `}`.
- **Não use `if` para devolver `bool`.** `if n > 0 { return true }` seguido de `return false` é só `return n > 0` (drill 004).

## O problema

Crie o arquivo **`decide.go`** nesta pasta, com **`package decide`**, e implemente:

1. **`func Max(a, b int) int`**: o maior dos dois.
   - `Max(3, 7)` → `7`
   - `Max(7, 3)` → `7`
   - `Max(-2, -5)` → `-2`
   - `Max(4, 4)` → `4`
   - Sem usar o `max` que já vem no Go (ele existe desde a versão 1.21): hoje o objetivo é o `if`.
2. **`func Classify(n int) string`**:
   - negativo → `"negative"`
   - zero → `"zero"`
   - positivo e par → `"positive even"`
   - positivo e ímpar → `"positive odd"`
   - Exemplos: `Classify(-4)` → `"negative"`, `Classify(0)` → `"zero"`, `Classify(8)` → `"positive even"`, `Classify(7)` → `"positive odd"`
3. **`func RoundToTen(n int) int`**: arredonda `n` (com `n` >= 0) para o múltiplo de 10 mais próximo. Quando termina em 5, sobe.
   - `RoundToTen(47)` → `50`
   - `RoundToTen(43)` → `40`
   - `RoundToTen(45)` → `50`
   - `RoundToTen(50)` → `50`
   - `RoundToTen(994)` → `990`

O `Classify` e o `RoundToTen` são as partes de pensar. **Antes de codar**, faça no `// Plano:`:
- **`Classify(-4)`:** qual condição você testa primeiro? O que sairia se testasse "é par" antes de "é negativo"?
- **`RoundToTen`:** a conta com números, para 47 e para 43, passo a passo (por exemplo, 47 → 40 → 50). Depois, a técnica do review do 005: pergunte de onde veio cada número e troque pelo nome de uma variável.

Se travar, as dicas estão em `HINTS.md`.

## Checklist

- [ ] `// Plano:` com a ordem do `Classify` e as contas do `RoundToTen` com 47 e 43.
- [ ] Comentário do pacote e de cada função, começando pelo nome e com verbo (`// Max returns ...`).
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Na raiz do repo:

  ```
  git add .
  git commit -m "drill 006: if else"
  git push
  ```
