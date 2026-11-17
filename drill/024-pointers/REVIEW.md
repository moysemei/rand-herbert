# Review — Drill 024

## Veredito

Entregue na segunda, com as quatro funções certas e o refazer certo de memória. As duas partes de pensar saíram como deviam. O que ficou foi uma dúvida de conceito, o `&` contra o `*`, e ela está respondida abaixo.

## Processo

- **No `AddBonus` e no `FindByOwner`, você desviou das duas armadilhas do drill:** percorreu pelos índices (não pela cópia do `range`) e devolveu `&accounts[i]`, o endereço do item de verdade. E o `return nil` ficou **depois** do loop, não num `else` dentro dele: o padrão buscar, certo.
- **O `Swap` saiu numa linha,** com o `*` nos quatro lugares. A troca do drill 017 reaproveitada.
- **As "Dificuldades" são suas,** com a tentativa errada copiada e a pergunta que ela gerou ("`&n` não continua sendo um int?"). É o registro mais útil que você podia me dar: mostra exatamente qual peça do modelo está faltando, e é por ela que eu começo.

## Erros e melhorias

**[ferramenta] `&` sobe um nível, `*` desce um nível.** No `Double`, você tentou `p := &n * 2` e `return p`. São duas dúvidas, e as duas perguntas que você escreveu são as certas.

*"`&n` não continua sendo um int?"* Não. Pense num papel com o endereço de uma casa escrito nele:

| Expressão | Tipo | O que é | No papel e na casa |
|---|---|---|---|
| `x` | `int` | o número | a casa |
| `&x` | `*int` | o endereço de `x` | o papel com o endereço da casa |
| `n`, dentro do `Double(&x)` | `*int` | uma cópia desse endereço | outro papel, com o mesmo endereço |
| `*n` | `int` | o número que está no endereço (o próprio `x`) | ir até a casa |
| `&n` | `**int` | o endereço da variável `n` | onde o papel está guardado |

**O `&` sobe um nível** (do valor para o endereço) e **o `*` desce** (do endereço para o valor). Dentro do `Double`, o `n` já é o endereço. Para chegar no número, você precisava descer (`*n`), mas o `&n` subiu mais um nível, para um `**int`. Por isso o "mismatched types": não existe conta entre um endereço de endereço e o número 2.

*"O `Double` não promete devolver nada; preciso de `return`?"* Não precisa, e nem pode devolver um valor: `return p` não compila numa função sem tipo de retorno. Uma função entrega o resultado de um destes dois jeitos:

| | Devolve | Escreve no endereço |
|---|---|---|
| Assinatura | `func Doubled(n int) int` | `func Double(n *int)` |
| Como entrega | `return n * 2` | `*n = *n * 2` |
| Quem chama | `y := Doubled(x)` (o `x` não muda) | `Double(&x)` (o `x` muda) |
| Tem `return`? | sim, com o valor | não; só um `return` vazio, se quiser sair antes do fim |

O `Double` é do segundo jeito: quem chamou vê o resultado na própria variável. E o "tá distante ainda" é normal: o Python esconde os endereços, então não existe nenhum hábito antigo em que se apoiar. O próximo drill começa exatamente por esse contraste, antes de avançarmos.

**[idiomático] Os comentários sumiram.** No 022 e no 023 estavam todos certos; aqui não tem o do pacote nem os das funções. Cada função exportada leva um comentário que começa pelo nome dela.

## Solução de referência

```go
// Package pointers practices addresses, nil, and changing values through pointers.
package pointers

// Account is a bank account; Balance is in reais.
type Account struct {
	Owner   string
	Balance int
}

// Double doubles the value at n.
// n is an address: *n reads and writes the int that lives there.
func Double(n *int) {
	*n *= 2
}

// Swap swaps the values at a and b.
// The * goes on both sides: swap what is at the addresses, not the addresses.
func Swap(a, b *int) {
	*a, *b = *b, *a
}

// AddBonus adds bonus to the balance of every account in the slice.
// Loop by index: accounts[i] is the real item; the range value would be a copy.
func AddBonus(accounts []Account, bonus int) {
	for i := range accounts {
		accounts[i].Balance += bonus
	}
}

// FindByOwner returns a pointer to owner's account inside accounts, or nil.
// Pattern: search. Return as soon as it is found; "not found" only after the loop.
func FindByOwner(accounts []Account, owner string) *Account {
	for i := range accounts {
		if accounts[i].Owner == owner {
			return &accounts[i]
		}
	}
	return nil
}
```

Por que assim: é a sua solução, mais os comentários. A única troca é o `*n *= 2`, que é o atalho do `*n = *n * 2`, como o `+=` que você já usa.

## Refazer de memória (IsAnagram)

Certo de memória, com uma função ajudante de contagem e o `maps.Equal`. No 021, essa estratégia veio da mentoria; agora as duas perguntas do método ("que representação deixa isso fácil?" e "já tenho uma função que faz parte disso?") saíram sozinhas. Só o nome da ajudante: em Go, palavras juntas usam maiúscula no meio, então `charFrequency`, não `charfrequency`.

## Para levar

O `&` sobe um nível (do valor para o endereço) e o `*` desce (do endereço para o valor). E uma função entrega o resultado de dois jeitos: com `return`, ou escrevendo no endereço que recebeu.
