# Review — Drill 017

## Veredito

Entregue, os quatro testes e o refazer passam. Mas o mais importante deste drill não está no código: está na sua seção "Geral" das "Dificuldades". Você descreveu, com precisão, o ponto que trava o seu aprendizado desde o drill 004. Este review é sobre ele, e o drill de hoje também.

## Processo

- **O seu diagnóstico está certo.** "Consigo entender exemplos e até responder perguntas sobre as operações, mas na hora de juntar tudo, escolher o for, definir as variáveis e montar a lógica, eu acabo travando." E também: "às vezes tento escrever o código antes de entender direito o que precisa acontecer". Esse é exatamente o problema, e ele tem método. Está logo abaixo.
- **A tabela do `ReverseInPlace` está certa,** e a frase "se continuasse, acabaria desfazendo as trocas" mostra que você entendeu por que parar no meio, não só onde.
- **No `MaxOf`, você achou o próprio bug:** o `return` dentro do `if`, que encerrava a função antes de ver todos os números. É a armadilha do padrão busca do drill 012, só que ao contrário.
- **No `Countdown`, você entendeu duas soluções:** `n - i` com o `i` subindo, ou o `i` descendo a partir de `n`.
- **Os comentários estão certos:** começam pelo nome, usam "returns" e dizem o que acontece com o slice vazio.
- **Você foi honesto sobre a ajuda e sobre o ChatGPT.** Sem isso, eu estaria criando drills para alguém que não existe.

## O método: da ideia ao código

Você mesmo escreveu a solução no fim das "Dificuldades": "decompor o problema, resolver manualmente, organizar os passos e só depois transformar isso em código Go". Aqui está o método, em quatro passos fixos:

1. **Exemplo à mão:** a entrada e a saída esperada, com números de verdade.
2. **Qual é o padrão?** Quase todo problema desta fase é um (ou dois) de um cardápio pequeno, e você já usou todos.
3. **Os passos em português,** como comentários, uma linha por passo.
4. **Cada linha vira Go.** Depois, a tabela de rastreio confere.

**O passo 2 é o que te falta.** Quando você diz "quando recebo a estrutura, consigo completar", a estrutura que o ChatGPT te dá é o padrão. A partir de agora, ela vai vir do cardápio, de você.

| Padrão | Pergunta que ele responde | Onde você já usou |
|---|---|---|
| **Acumular** | "Quanto dá somando (ou multiplicando)?" | `SumTo`, `Factorial`, `Sum` |
| **Contar** | "Quantos são?" | `CountMultiples`, `CountPrimes`, `CountVowels` |
| **Melhor até agora** | "Qual é o maior, o menor, o mais longo?" | `MinMaxDigit`, `MaxOf` |
| **Buscar** | "Existe algum?" / "Qual é o primeiro?" | `IsPrime`, `NextPrime`, `SumNumbers` |
| **Montar** | "Qual é a nova lista, o novo texto?" | `Countdown`, `Reverse` |
| **Percorrer dígitos** | "O que tem dentro do número?" | `SumDigits`, `ReverseNumber` |
| **Trocar pontas** | "Como inverter no lugar?" | `ReverseInPlace` |

Veja o `Countdown`, que te embananou, passando pelos quatro passos:

1. **Exemplo:** `Countdown(3)` → `[3 2 1]`.
2. **Padrão:** a pergunta é "qual é a nova lista?", então é **montar**. O esqueleto já vem pronto: slice vazio, um loop, `append` em cada volta, `return`.
3. **Passos:** "começo com um slice vazio"; "para cada valor que deve entrar, faço append"; "devolvo o slice".
4. **A pergunta que falta:** **quais valores entram, e em que ordem?** 3, depois 2, depois 1. Qual variável pode passar por esses valores? Um `i` que começa em `n` e desce até 1. É a técnica do drill 005 ("de onde vem cada número?") aplicada a uma lista.

As suas tentativas (`append(result, -1)` e `append(result, n-1)`) tentavam escrever a **operação** ("diminuir") em vez do **valor** que entra. *Como pensar:* no padrão montar, liste primeiro os valores que vão entrar, e só depois procure a variável que passa por eles.

## Erros e melhorias

**[idiomático] `max` como nome de variável, de novo.** Já apareceu no review do 014: desde o Go 1.21, `max` é uma função do próprio Go. Use `largest`.

**[idiomático] Comentário no refazer.** O `DigitalRoot` e o `sumDigits` do `redo/` saíram certos de memória, mas sem comentário.

## Sobre o ChatGPT

Nenhuma culpa: você usou, disse que usou e ainda explicou o que ficou faltando. Mas, pelo que você mesmo descreveu, quando o ChatGPT te dá a estrutura, ele faz o passo 2 por você, e é justamente o passo que você precisa treinar. A ordem quando travar passa a ser:

1. Qual é o padrão? (olhe a tabela acima)
2. A tabela de rastreio com o exemplo.
3. As dicas, um nível de cada vez.
4. Uma mensagem para mim no chat.
5. ChatGPT, se ainda precisar, perguntando só **"qual padrão é esse problema?"**, sem pedir código.

## Solução de referência

```go
// Package lists works with slices of ints.
package lists

// Sum returns the sum of the elements of nums.
// Pattern: accumulate.
func Sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// MaxOf returns the largest element of nums. ok is false when nums is empty.
// Pattern: best so far.
func MaxOf(nums []int) (largest int, ok bool) {
	if len(nums) == 0 {
		return 0, false
	}
	largest = nums[0] // a real element, never a guess like 0
	for _, v := range nums {
		if v > largest {
			largest = v
		}
	}
	return largest, true
}

// Countdown returns n, n-1, ..., 1.
// Pattern: build. The values that go in are n, n-1, ..., 1: exactly what i walks through.
func Countdown(n int) []int {
	var result []int
	for i := n; i >= 1; i-- {
		result = append(result, i)
	}
	return result
}

// ReverseInPlace reverses nums by swapping its elements.
// Pattern: swap the ends, moving towards the middle.
func ReverseInPlace(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		j := len(nums) - 1 - i // the matching index from the other end
		nums[i], nums[j] = nums[j], nums[i]
	}
}
```

Por que assim:

- **Cada função diz o seu padrão no comentário.** Num código de verdade não precisaria, mas enquanto o cardápio está sendo decorado, escrever o nome do padrão obriga você a fazer o passo 2.
- **O `Countdown` com o `i` descendo** usa o `i` como o próprio valor que entra. A sua versão (`n - i` com o `i` subindo) também está certa; ela só exige uma conta a mais para chegar no valor.

## Refazer de memória (DigitalRoot)

Certo de memória, com o ajudante `sumDigits`, o `n >= 10` e o `n = sumDigits(n)` guardando o resultado. Faltou só o comentário.

## Para levar

Antes do código, quatro passos: exemplo à mão, qual é o padrão, os passos em português e só então Go. A estrutura vem do padrão; o detalhe vem do exemplo.
