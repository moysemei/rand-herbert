# Review — Drill 014

## Veredito

Entregue completo no primeiro dia de 45 minutos: refazer de memória, os três testes, plano com tabela, a explicação do `d <= n/d` e as "Dificuldades". É a entrega com o processo mais completo até aqui.

## Processo

- **A tabela do `MinMaxDigit(4729)` está certa linha por linha,** e começa no lugar certo: a primeira volta mostra `min = 9` e `max = 9`, porque os dois começam com o último dígito. Logo abaixo, você escreveu o que acontece com `n = 0` antes de codar. A armadilha do valor inicial não te pegou.
- **A explicação do `d <= n/d` está correta e é sua:** equivalente a `d*d <= n` para `d` positivo, e evita o overflow. Ponto de atenção fechado.
- **As "Dificuldades" contam o caminho do raciocínio,** não só o resultado: "pensei que precisaria armazenar todos os dígitos. Depois entendi que bastava inicializar min e max com um dígito real". Esse é exatamente o tipo de registro que me deixa calibrar.
- **No `SplitDigits`, você usou o segundo caminho para as dezenas,** `(n / 10) % 10`, aquele que o review do 005 apresentou ("tire o último, pegue o novo último"). Ele é o que generaliza para os loops.
- **Os comentários documentam o limite de entrada:** "It expects n to be between 100 and 999." Quem chama a função precisa saber disso, e o tipo `int` não diz.

## Erros e melhorias

**[idiomático] `min` e `max` são nomes de funções do Go.** Desde o Go 1.21, `min` e `max` são funções embutidas (`max(3, 7)` devolve 7). Dar esses nomes a variáveis compila, mas esconde as funções dentro do escopo, e linters como o `revive` avisam ("redefinition of the built-in function"). Use `lo`/`hi` ou `smallest`/`largest`. É o mesmo cuidado de não chamar uma variável de `len` ou `string`.

**[lógica] "Armazenar todos os dígitos" contra "dois acumuladores".** A sua primeira ideia, guardar os dígitos e depois achar o menor e o maior, é o reflexo do Python (`min(digits)`, `max(digits)`). Funciona, mas exige guardar tudo e percorrer duas vezes. A solução que você escreveu passa uma vez só e guarda só dois números. Em entrevista, isso tem nome: **single pass** com memória constante. Quando der para responder "o que eu preciso lembrar do que já vi?" com uma ou duas variáveis, não guarde a lista inteira.

**[sintaxe] O `=` no lugar de `==` de novo.** Você anotou que corrigiu, e isso basta. É a segunda vez que ele aparece (drills 006 e 014): quando o compilador falar em *assignment* dentro de um `if`, é ele.

## Solução de referência

```go
// Package multi has functions that return more than one value.
package multi

// SplitDigits returns the hundreds, tens and ones digits of n.
// n must be between 100 and 999.
func SplitDigits(n int) (hundreds, tens, ones int) {
	hundreds = n / 100
	tens = n / 10 % 10 // drop the last digit, then take the new last one
	ones = n % 10
	return hundreds, tens, ones
}

// Divide returns the integer division a / b.
// ok is false when b is zero, because there is no answer.
func Divide(a, b int) (quotient int, ok bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// MinMaxDigit returns the smallest and the largest digit of n (n >= 0).
func MinMaxDigit(n int) (smallest, largest int) {
	smallest, largest = n%10, n%10 // start with a real digit: works for n = 0 too
	for n > 0 {
		digit := n % 10
		if digit < smallest {
			smallest = digit
		}
		if digit > largest {
			largest = digit
		}
		n /= 10
	}
	return smallest, largest
}
```

Por que assim:

- **Retornos nomeados em `Divide` e `MinMaxDigit`:** `(quotient int, ok bool)` e `(smallest, largest int)` documentam a ordem dos valores na própria assinatura. Quem chama sabe o que é o primeiro e o que é o segundo sem abrir a função.
- **Sem linha em branco logo depois do `{`.** No `SplitDigits` você abriu a função com uma linha vazia. O `gofmt` aceita, mas não é comum; as linhas em branco servem para separar blocos de ideia dentro da função.

## Refazer de memória (Power)

Saiu certo de memória: `product := 1`, `exp` voltas com `i := 1; i <= exp` e a `base` multiplicada em cada uma. O "valor que permanece", que pediu a dica 2 no drill 010, agora veio sozinho. Faltou só o comentário da função.

## Para levar

Uma função Go pode devolver vários valores, e a ordem da assinatura é um contrato. E quando você só precisa lembrar do menor e do maior, dois acumuladores bastam: não guarde a lista inteira.
