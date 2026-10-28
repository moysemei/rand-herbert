# Review — Drill 015

## Veredito

Entregue no mesmo dia, às 12h, com o refazer, os três testes passando e as "Dificuldades" contando o caminho todo. Bytes contra runes ficou claro (o `Sizes` saiu certo). O `Reverse` pediu ajuda, e a explicação está logo abaixo.

## Processo

- **A tabela do `Reverse("ação")` está certa,** e é ela que explica a função. Mais abaixo eu volto a ela.
- **Você leu o erro e entendeu o que ele dizia.** No `Sizes`, o `for _, r := range s` deu "r declarado mas não usado", e você trocou por `for range s`, que é exatamente a forma para "só contar as voltas".
- **A sua primeira ideia para o `Reverse` não estava errada:** usar a posição para levar o último para o começo é como se inverte um **slice** (com dois índices trocando valores, das pontas para o meio). Faltava só a ferramenta, que chega no tópico 12. Guarde essa ideia, porque ela vai voltar.
- **"Precisei de ajuda para compreender"** e "sabe Deus..." no plano: honestidade que me ajuda a calibrar. Obrigado por escrever.
- **O refazer do `ReverseNumber` saiu de memória,** com o `digit` nomeado. A fórmula que travou no 011 agora veio sozinha.

## Erros e melhorias

**[sintaxe] `r == 'a' || 'e'` não existe.** Você tentou primeiro `r == 'a' ou 'e' ou 'i'`, como se fala em português. Mas o `||` liga duas **condições** completas, e `'e'` sozinho não é uma condição: é uma rune. O Go recusa com `mismatched types untyped bool and untyped rune`. (Em Python, `r == 'a' or 'e'` compila, mas está errado do mesmo jeito: o `'e'` sozinho conta como verdadeiro, e a condição passa para qualquer letra. É um bug clássico.) *Como pensar:* cada lado do `||` tem que fazer sentido sozinho como pergunta de sim ou não. Por isso é `r == 'a' || r == 'e'`.

**[idiomático] Cinco `||` comparando a mesma variável: use `switch`.** É o caso do drill 007, "um valor contra várias opções":

```go
switch r {
case 'a', 'e', 'i', 'o', 'u':
	count++
}
```

Sem pipes, sem repetir o `r`. (No próximo drill tem um terceiro jeito, com o pacote `strings`.)

**Sobre não achar o `|` no teclado:** no teclado brasileiro (ABNT2), ele fica na tecla entre o **Shift esquerdo e o Z**: **Shift + `\`**. No teclado americano, fica acima do Enter: também **Shift + `\`**. Como último recurso no Windows: **Alt + 124** no teclado numérico. Vale treinar, porque `||` vai aparecer em todo código Go que você escrever.

**[lógica] Por que "colocar na frente" inverte.** Esta é a explicação que te faltou, a partir da sua própria tabela:

| volta | `r` | `result` antes | `string(r) + result` |
|---|---|---|---|
| 1 | `a` | `""` | `"a"` |
| 2 | `ç` | `"a"` | `"ça"` |
| 3 | `ã` | `"ça"` | `"ãça"` |
| 4 | `o` | `"ãça"` | `"oãça"` |

O loop lê a string **do começo para o fim**, mas cada rune nova entra **na frente** de tudo que já está no `result`. Então a primeira rune lida, o `a`, vai sendo empurrada para trás a cada volta, até terminar no fim. É como empilhar pratos: o primeiro que você põe fica embaixo de todos.

E compare com o `ReverseNumber` que você acabou de refazer: lá, você lê os dígitos **do fim para o começo** (`% 10`) e cada um entra **atrás** (`*10 + digit`). São dois jeitos de inverter:
- ler do fim e colocar atrás;
- ler do começo e colocar na frente.

## Solução de referência

```go
// Package runes works with strings byte by byte and rune by rune.
package runes

// CountVowels returns how many lowercase unaccented vowels (a, e, i, o, u) s has.
func CountVowels(s string) int {
	count := 0
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

// Sizes returns how many bytes and how many runes s has.
func Sizes(s string) (bytes, runes int) {
	for range s { // one round per rune
		runes++
	}
	return len(s), runes
}

// Reverse returns s with its runes in reverse order.
func Reverse(s string) string {
	result := ""
	for _, r := range s {
		result = string(r) + result // each new rune goes to the front
	}
	return result
}
```

Por que assim:

- **`switch` no `CountVowels`:** as cinco vogais numa linha só, como uma lista.
- **No `Reverse`, cada `+` cria uma string nova,** porque strings são imutáveis. Para textos curtos, não faz diferença. Para textos enormes, o jeito eficiente é converter para `[]rune` e trocar as pontas, que é a sua primeira ideia. Ela volta no tópico de slices.

## Refazer de memória (ReverseNumber)

Certo de memória, com `digit := n % 10` e `reversed = reversed*10 + digit`. A fórmula que travou no drill 011 agora saiu sem ajuda. Faltou só o comentário da função, e as linhas em branco entre cada linha do loop podem sair: num loop de três linhas, elas mais separam do que organizam.

## Para levar

O `||` liga duas perguntas completas: `r == 'a' || r == 'e'`. E para inverter, ou você lê do fim e coloca atrás, ou lê do começo e coloca na frente.
