# Review — Drill 019

## Veredito

Entregue no feriado, com o refazer e os três testes passando. O refazer mostra que a repetição está funcionando: o `SumNumbers`, com o `strings.Fields` e o `Atoi` + `err` que tinham sumido no drill 018, saiu inteiro de memória. O `FirstN` pediu as dicas 2 e 3, e o `RemoveDuplicates` travou num ponto específico: como dizer "ignorar" em Go.

## Processo

- **O refazer funcionou exatamente como deveria.** As duas ferramentas que você tinha esquecido no 018 voltaram no 019, sem consulta. É para isso que ele existe.
- **O `contains` também saiu de memória,** e o plano dele tem o padrão certo: "se encontrar, retornar true imediatamente; se terminar sem encontrar, retornar false".
- **A tabela do `RemoveDuplicates` foi feita sozinha, e está certa** (no Excel, inclusive). A lógica estava inteira antes de você escrever qualquer linha de Go.
- **O plano tem o método nas três funções e no ajudante.** No `RemoveDuplicates`, os dois padrões juntos (buscar + montar).
- **"Não posso dizer que resolvi sozinho"** sobre o `FirstN`. Esse tipo de honestidade deixa o diagnóstico preciso. E a sua alternativa (devolver a cópia inteira com retorno antecipado quando `n` passa do tamanho) também funciona.
- **Faltaram os comentários do pacote e das funções.** Nos drills 013 a 018 eles vinham certos.

## Erros e melhorias

**[lógica] Como dizer "ignorar" em Go: não diga nada.** No `RemoveDuplicates`, você queria ignorar os repetidos e tentou `return false` e depois `if/else`. A tabela estava certa; o que travou foi a tradução. A regra:

> **Cada `if` guarda uma ação. O que não tem ação não precisa de código.**

"Ignorar" não é uma ação, é a ausência dela. Então escreva o passo em português pela ação que **acontece**: não "se já está, ignoro; senão, adiciono", e sim **"se ainda não está, adiciono"**. Essa frase vira Go direto: `if !contains(result, n) { result = append(result, n) }`. O caso "já está" não aparece no código, e não precisa aparecer.

E sobre as tentativas:
- O `return` sai da **função inteira**, então não serve para pular um item.
- O `continue` pula para a próxima volta. Funciona, mas aqui nem é preciso: o `if` sozinho já resolve.

*Como pensar:* quando um passo do plano tiver "ignorar", "pular" ou "não fazer nada", reescreva-o como "só faço X quando Y".

**[lógica] Antes do padrão de loop, procure a ferramenta.** No `FirstN`, você pensou em buscar e em percorrer o slice. Mas pegar os primeiros `n` não é um loop: é uma **operação pronta**, o fatiamento (`nums[:n]`). Os padrões do cardápio são formas de loop; algumas tarefas não precisam de loop nenhum. Antes do passo 2 do método, faça uma pergunta:

> **Existe uma operação pronta que já faz isso?** Fatiar, `len`, `strings.Fields`, `strings.ToLower`, `slices.Clone`...

Se existir, use. Se não, escolha o padrão. É aqui que a cola `notes/toolbox.md` (sugestão do review 018) ajuda: ela é a lista das operações prontas que você já conhece.

**[sintaxe] `contains(result, num) == !true` é só `!contains(result, num)`.** O `contains` já devolve um `bool`, e o `if` aceita o `bool` direto. Comparar com `true` (ou com `!true`, que é `false`) não acrescenta nada. É a mesma lição do drill 004: uma comparação já **é** um `bool`. Em Go:
- `if x == true` se escreve `if x`;
- `if x == false` se escreve `if !x`.

**[idiomático] Comentários.** O pacote e as quatro funções ficaram sem. O formato de sempre: `// Package slices2 ...` e `// FirstN returns ...`.

## Solução de referência

```go
// Package slices2 slices, copies and filters slices without changing the input.
package slices2

import "slices"

// Evens returns the even numbers of nums, in order, in a new slice.
// Pattern: build with a filter.
func Evens(nums []int) []int {
	var result []int
	for _, n := range nums {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	return result
}

// FirstN returns an independent copy of the first n elements of nums (n >= 0).
// No pattern needed: slicing does the selection, slices.Clone does the copy.
func FirstN(nums []int, n int) []int {
	if n > len(nums) {
		n = len(nums) // never slice past the end
	}
	return slices.Clone(nums[:n])
}

// contains reports whether target is in nums.
func contains(nums []int, target int) bool {
	for _, n := range nums {
		if n == target {
			return true
		}
	}
	return false
}

// RemoveDuplicates returns the values of nums once each, in order of first appearance.
// Patterns: build + search.
func RemoveDuplicates(nums []int) []int {
	var result []int
	for _, n := range nums {
		if !contains(result, n) { // only act when there is something to do
			result = append(result, n)
		}
	}
	return result
}
```

Por que assim:

- **O comentário do `FirstN` diz que não há padrão de loop:** a ferramenta resolve. Saber quando **não** usar um loop vale tanto quanto saber montar um.
- **O `RemoveDuplicates` verifica o `result` a cada número.** Com listas grandes, isso fica lento (para cada número, ele percorre o resultado inteiro). No próximo drill, com maps, você vai escrever a versão rápida.

## Refazer de memória (SumNumbers)

Certo de memória: `strings.Fields`, `strconv.Atoi(word)` com o `word` (não o `s`, que era o bug do drill 016), o `err` conferido antes de somar e o retorno antecipado. Faltou só o comentário.

## Para levar

Cada `if` guarda uma ação. Para "ignorar", escreva "só faço X quando Y". E antes de escolher um padrão de loop, pergunte se já existe uma operação pronta que faz aquilo.
