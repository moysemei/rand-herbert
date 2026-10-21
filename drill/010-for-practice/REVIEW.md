# Review — Drill 010

## Veredito

Entregue, os três testes passam, e a diferença para o 009 é grande: `SumEven` sem dica, `CountMultiples` de primeira, `Power` com a dica 2. O padrão "loop gera, `if` filtra, acumulador guarda" firmou.

## Processo

- **O `CountMultiples` saiu de primeira, e com `result++`, não `result += i`.** Você separou "contar" de "somar" sem precisar de dica. No 009, o `SumOfMultiples`, com o mesmo formato, tinha vindo do GPT. Agora o formato é seu.
- **No `Power`, você registrou exatamente onde errou:** "não colocar um valor que permaneça", com `product += base * base` e `i := base`. Achar a frase certa para o erro é metade do aprendizado. E o "valor que permanece" é o papel da tabela do README: a `base` é o **valor fixo**, e o `i` só conta as voltas.
- **O `Power` com `i := 1; i <= exp` dá exatamente `exp` voltas.** É uma forma tão certa quanto `i := 0; i < exp`. O importante é o número de voltas, e você acertou.
- **Os comentários voltaram,** começando pelo nome.
- **Sobre o plano, você foi honesto:** "eu começo a montar e vou tendo ideia e passo pra prática". Está ok: nas partes fáceis, ir direto funciona, e funcionou. Então a proposta muda. A partir do 011, o plano obrigatório é **uma linha por função**: o que muda, o que fica fixo, o valor inicial e quando o loop para. A tabela fica como **primeiro passo quando travar**, antes de abrir uma dica. No `Power`, ela teria mostrado o caminho: volta 1 → 2, volta 2 → 4, volta 3 → 8. É sempre "o anterior vezes 2".

## Erros e melhorias

**[lógica] O valor que permanece.** Suas tentativas no `Power` eram `product += base * base` e `i := base`. Provavelmente você estava tentando escrever a conta inteira numa linha só (2 × 2 × 2 × 2), como no fatorial do 009. A virada é a mesma do review passado: **o loop faz um pedaço por volta**. Em `Power(2, 4)`, o pedaço é sempre "multiplicar por 2":

| volta | `i` | `product` depois |
|---|---|---|
| — | | 1 (antes do loop) |
| 1 | 1 | 1 × 2 = 2 |
| 2 | 2 | 2 × 2 = 4 |
| 3 | 3 | 4 × 2 = 8 |
| 4 | 4 | 8 × 2 = 16 |

O `i` aparece na tabela só para contar as voltas, e não entra em nenhuma conta. A `base` entra em todas.

**[idiomático] Nome do acumulador.** No `CountMultiples`, `result` funciona, mas `count` diz o que a variável guarda. No `Power`, `product` (ou `result`) está bom.

**[idiomático] Espaço depois do `//` e verbo nos comentários.** O `gofmt` reescreve `//SumEven ...` como `// SumEven ...`, com espaço. Para não depender de lembrar, rode **`go fmt ./...`** na raiz do repo antes do commit. Ele formata todos os arquivos no padrão do Go. (Se quiser, depois configuramos o nvim para fazer isso ao salvar.) E dois comentários ficaram sem verbo:
- `SumEven sum all even numbers ...` → `SumEven returns the sum of the even numbers from 1 to n.`
- `CountMultiples how many numbers ...` → `CountMultiples returns how many numbers ...`

## Solução de referência

```go
// Package practice has small loops to practice for.
package practice

// SumEven returns the sum of the even numbers from 1 to n.
func SumEven(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			total += i
		}
	}
	return total
}

// CountMultiples returns how many numbers from 1 to limit are multiples of k.
func CountMultiples(limit, k int) int {
	count := 0
	for i := 1; i <= limit; i++ {
		if i%k == 0 {
			count++ // count, don't sum: the value of i doesn't matter here
		}
	}
	return count
}

// Power returns base raised to exp (exp >= 0), without the math package.
func Power(base, exp int) int {
	result := 1                // multiplying by 1 changes nothing; also the answer for exp = 0
	for i := 0; i < exp; i++ { // i only counts the rounds
		result *= base
	}
	return result
}
```

Por que assim:

- **O `count` deixa claro, pelo nome,** que a função conta e não soma.
- **O comentário do `Power` explica o 1 inicial duas vezes:** ele é neutro na multiplicação e já é a resposta quando o loop roda zero vezes.
- **O seu código e o de referência são praticamente iguais.** As diferenças são de nome e de comentário, não de lógica.

## Para levar

Antes de um loop, pergunte: o que muda a cada volta, e o que fica fixo? No `Power`, a `base` fica fixa e entra em toda volta; o `i` muda, mas só conta.
