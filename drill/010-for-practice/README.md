# Drill 010 — Reforço de for: rastrear o loop

**Hoje:** sem conceito novo. Três loops curtos para firmar o `for` do drill 009, e uma ferramenta: a tabela de rastreio, feita no papel antes de codar.

## O que é

**Cada variável de um loop tem um papel:**

| Papel | Exemplo | Muda a cada volta? |
|---|---|---|
| **Contador**: o número da vez | `i` | sim |
| **Limite**: diz quando parar | `n`, `limit`, `exp` | não, nunca |
| **Acumulador**: guarda o resultado | `total`, `count`, `result` | sim |
| **Valor fixo** usado na conta (às vezes) | `base`, `word` | não |

Antes de escrever um loop, responda: **o que muda a cada volta, e o que fica fixo?**

**Somar e contar são acumuladores diferentes:**

```go
total += i // sum: adds the value of i
count++    // count: adds 1, whatever i is
```

**Às vezes o `i` não entra na conta.** Lembra do `Repeat` do README do 009? Ali o `i` só conta as voltas, e quem entra no resultado é o `word`. O loop serve para "fazer isso N vezes".

**A tabela de rastreio** é o loop feito à mão: uma linha por volta, uma coluna por variável que muda.

## Por que é assim

- **Em entrevista isso tem nome: *dry run*.** Depois de escrever a solução, o entrevistador espera que você rode o código na mão, com um exemplo pequeno, linha por linha. Quem faz a tabela encontra o próprio bug antes do entrevistador.
- **O computador não faz nada que você não consiga fazer no papel, só faz mais rápido.** Se você não consegue preencher a tabela, o problema ainda não está claro. Isso é sinal para voltar ao plano, não para pedir o código pronto.
- **A tabela ensina o valor inicial.** Na primeira linha você precisa escrever o valor do acumulador antes da primeira volta. É ali que aparece se ele começa em 0 ou em 1.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package squares

// SumOfSquares returns 1*1 + 2*2 + ... + n*n.
func SumOfSquares(n int) int {
	total := 0 // sum: starts at 0
	for i := 1; i <= n; i++ {
		total += i * i // i changes every round; n only says when to stop
	}
	return total
}
```

Rastreando `SumOfSquares(3)`:

| volta | `i` | `i * i` | `total` depois |
|---|---|---|---|
| — | | | 0 (antes do loop) |
| 1 | 1 | 1 | 1 |
| 2 | 2 | 4 | 5 |
| 3 | 3 | 9 | 14 |
| — | 4 | | `4 <= 3` é falso, e o loop acaba |

## Armadilhas

- **Somar quando era para contar.** `total += i` soma os valores; `count++` conta quantos passaram.
- **O limite dentro da conta.** O `n` (ou `limit`, ou `exp`) só diz quando parar. Se ele aparece dentro do `{ }` do loop, desconfie.
- **O valor inicial.** Soma e contagem começam em 0; multiplicação começa em 1.
- **Incluir ou não o limite.** "De 1 até `limit`" inclui o `limit`: `i <= limit`. "Menores que `limit`" não inclui: `i < limit`.
- **Loop que roda zero vezes.** Com `exp = 0`, um loop de `exp` voltas nem começa, e o resultado é o valor inicial. Confira se isso já dá a resposta certa.

## O problema

Crie o arquivo **`practice.go`** nesta pasta, com **`package practice`**, e implemente usando `for`:

1. **`func SumEven(n int) int`**: a soma dos números pares de 1 até `n`, incluindo o `n`.
   - `SumEven(6)` → `12` (2 + 4 + 6)
   - `SumEven(7)` → `12`
   - `SumEven(1)` → `0`
   - `SumEven(10)` → `30`
2. **`func CountMultiples(limit, k int) int`**: **quantos** números de 1 até `limit` (incluindo o `limit`) são múltiplos de `k`.
   - `CountMultiples(10, 3)` → `3` (3, 6 e 9)
   - `CountMultiples(10, 5)` → `2` (5 e 10)
   - `CountMultiples(20, 4)` → `5`
   - `CountMultiples(2, 3)` → `0`
3. **`func Power(base, exp int) int`**: `base` elevado a `exp` (com `exp` >= 0), sem usar o pacote `math`.
   - `Power(2, 10)` → `1024`
   - `Power(3, 4)` → `81`
   - `Power(7, 1)` → `7`
   - `Power(5, 0)` → `1`

**Antes de codar**, faça no `// Plano:`, para cada função, a resposta a "o que muda a cada volta e o que fica fixo?". E duas tabelas de rastreio:
- **`CountMultiples(10, 3)`:** colunas `i`, "é múltiplo?" e `count`.
- **`Power(2, 4)`:** colunas volta, `i` e `result`. O `i` entra na conta?

Se travar: primeiro as dicas, depois uma mensagem para mim no chat. Se for usar o GPT, peça só uma dica, sem código.

## Checklist

- [ ] `// Plano:` com "o que muda e o que fica fixo" das três funções e as duas tabelas.
- [ ] Comentário do pacote e de cada função, começando pelo nome, com `returns`.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Na raiz do repo:

  ```
  git add .
  git commit -m "drill 010: for practice"
  git push
  ```
