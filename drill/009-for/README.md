# Drill 009 — for

**Hoje:** repetir uma conta com o `for` clássico e acumular o resultado numa variável.

## O que é

O `for` clássico tem três partes, separadas por `;`:

```go
for i := 1; i <= 3; i++ {
	// runs with i = 1, then 2, then 3
}
//  ↑ start   ↑ condition  ↑ step
```

1. **Início** (`i := 1`): roda uma vez, antes de tudo.
2. **Condição** (`i <= 3`): testada antes de cada volta. Se der `false`, o loop acaba.
3. **Passo** (`i++`): roda no fim de cada volta. `i++` é o mesmo que `i = i + 1`.

O padrão que você vai usar o tempo todo é o **acumulador**: uma variável criada **antes** do loop, atualizada **dentro** dele e devolvida **depois**.

```go
total := 0              // before: the starting value
for i := 1; i <= 3; i++ {
	total += i          // inside: total = total + i
}
return total            // after: 1 + 2 + 3 = 6
```

| volta | `i` | `total` depois da volta |
|---|---|---|
| 1 | 1 | 1 |
| 2 | 2 | 3 |
| 3 | 3 | 6 |
| — | 4 | a condição `4 <= 3` é falsa, e o loop acaba |

## Por que é assim

- **O Go só tem `for`.** Não existe `while` nem `do-while`: o mesmo `for` faz os três papéis (o "while" é o drill que vem). Uma palavra só, menos coisa para lembrar.
- **Comparado com Python e JS:** o Python faz `for i in range(1, n + 1):`. O JS faz `for (let i = 1; i <= n; i++)`. O Go é o JS sem os parênteses e com chaves obrigatórias.
- **`i++` é um comando, não um valor.** Em JS dá para escrever `x = i++`. Em Go, não: `i++` fica sozinho na linha. Isso evita linhas espertas e difíceis de ler.
- **Desde o Go 1.22 existe `for i := range n`,** que vai de 0 a n-1. Você vai ver bastante em código novo. Hoje, treine a forma clássica, porque ela mostra as três partes e deixa você escolher onde começar e onde parar.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package text

// Repeat returns word repeated times times: Repeat("ab", 3) is "ababab".
func Repeat(word string, times int) string {
	result := "" // start: an empty string
	for i := 0; i < times; i++ {
		result += word // add one more copy on each round
	}
	return result
}
```

Repare no valor inicial: para juntar texto, você começa do vazio. Sempre se pergunte qual é o valor inicial certo para a operação que o loop repete.

## Armadilhas

- **Erro de um a mais ou um a menos.** `i < n` ou `i <= n`? Começa em 0 ou em 1? Teste na cabeça a **primeira** e a **última** volta.
- **Acumulador dentro do loop.** Se o `total := 0` ficar dentro das chaves do `for`, ele volta a zero em toda volta.
- **`:=` dentro do loop cria outra variável.** `total := total + i` dentro do loop não atualiza o `total` de fora: cria um novo, que some no fim da volta. Dentro do loop é `=` ou `+=`.
- **O valor inicial depende da operação.** Para somar, começar em 0 não muda nada. Para multiplicar, 0 estraga tudo.
- **`int` tem limite.** No Python os inteiros crescem sem fim. Em Go, um `int` (64 bits) passa do limite em 21! e "dá a volta" sem erro nenhum, virando um número errado. Os testes de hoje ficam bem abaixo disso.

## O problema

Crie o arquivo **`loops.go`** nesta pasta, com **`package loops`**, e implemente usando `for`:

1. **`func SumTo(n int) int`**: a soma de 1 até `n`.
   - `SumTo(5)` → `15` (1 + 2 + 3 + 4 + 5)
   - `SumTo(1)` → `1`
   - `SumTo(0)` → `0`
   - `SumTo(100)` → `5050`
2. **`func Factorial(n int) int`**: o fatorial, `n × (n-1) × ... × 1`.
   - `Factorial(5)` → `120`
   - `Factorial(1)` → `1`
   - `Factorial(0)` → `1` (por definição, 0! = 1)
   - `Factorial(10)` → `3628800`
3. **`func SumOfMultiples(limit int) int`**: a soma de todos os números **menores** que `limit` que são múltiplos de 3 **ou** de 5.
   - `SumOfMultiples(10)` → `23` (3 + 5 + 6 + 9)
   - `SumOfMultiples(16)` → `60` (3 + 5 + 6 + 9 + 10 + 12 + 15; o 15 entra uma vez só)
   - `SumOfMultiples(1)` → `0`
   - `SumOfMultiples(1000)` → `233168`

O `Factorial` e o `SumOfMultiples` são as partes de pensar. **Antes de codar**, faça no `// Plano:`:
- **`Factorial(4)`:** uma tabelinha como a do README (volta, `i`, resultado). Com que valor o resultado começa?
- **`SumOfMultiples(10)`:** qual é o primeiro e o último `i` do loop? Como você decide se um `i` entra na soma?

Se travar, as dicas estão em `HINTS.md`.

## Checklist

- [ ] `// Plano:` com a tabelinha do `Factorial(4)` e o primeiro e o último `i` do `SumOfMultiples(10)`.
- [ ] Comentário do pacote e de cada função, começando pelo nome, com `returns`.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Na raiz do repo:

  ```
  git add .
  git commit -m "drill 009: for"
  git push
  ```
