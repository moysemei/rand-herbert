# Review — Drill 005

## Veredito

Entregue no mesmo dia, os três testes passam. Desta vez você resolveu sem pedir ajuda no chat: usou as dicas até o nível 2 nas dezenas e o nível 3 no inverso. Era exatamente para isso que as dicas foram feitas.

## Processo

- **A sua pergunta sobre o README: pode, e deve.** O README é material de consulta, como a documentação que você vai abrir todo dia no trabalho. A regra vale só para ver uma solução pronta antes de tentar. Usar o exemplo do "Como se usa" para entender a sintaxe é o uso certo.
- **Você remontou o número à mão no plano, e certo:** `2*100 + 7*10 + 4*1 = 274`. Essa linha já é o raciocínio do `ReverseThreeDigits` inteiro (mais abaixo, em "Erros", explico o passo que faltou).
- **Você atualizou o plano quando entendeu algo.** "Atualização, acho que entendi essa primeira parte da sobra da divisão." O plano mudar durante o trabalho é sinal de que ele está sendo usado de verdade.
- **Você usou as dicas do jeito certo:** um nível de cada vez. Nas dezenas, parou no 2, quando já dava para seguir sozinho. E anotou exatamente quais níveis usou, o que me mostra onde está a dificuldade.
- **Você usou parênteses em `(n % 100) / 10`,** como a armadilha do README sugeria: não muda o resultado, mas deixa claro para quem lê qual conta vem primeiro.
- **Uma observação, não uma crítica:** você pensou em resolver com slices. É o reflexo do Python (`str(n)[-1]`). Funciona, mas a conta com `%` e `/` é mais barata e é o que as entrevistas esperam.

## Erros e melhorias

**[lógica] O dígito do meio.** As centenas e as unidades saíram de primeira, mas as dezenas pediram a dica 2. O plano tinha a remontagem, mas faltou a outra metade, a desmontagem. Nisso a culpa é minha: o README dizia "desmontar 472 à mão" sem mostrar o que isso significava. A ideia era escrever as contas com os números, uma pergunta por passo:

```
472 / 100 = 4     how many whole hundreds fit?    -> hundreds
472 % 100 = 72    what is left without them?
 72 / 10  = 7     how many whole tens fit in 72?  -> tens
472 % 10  = 2     what is left after the tens?    -> ones
```

Com isso no papel, o código é só trocar 472 por `n`. Existe ainda um segundo caminho para as dezenas, que vale conhecer: `n / 10 % 10` (tire o último dígito, `472 / 10 = 47`, e pegue o novo último, `47 % 10 = 7`). Esse vai aparecer de novo quando chegar o `for`: repetir "pega o último, tira o último" é como se percorre um número de qualquer tamanho.

**[lógica → sintaxe] O inverso estava no seu plano.** Você escreveu `2*100 + 7*10 + 4*1 = 274` e mesmo assim precisou da dica 3. Faltou um passo só: **perguntar de onde veio cada número do exemplo** e trocar pelo nome da variável que guarda ele.
- O `2` são as unidades de 472, `ones`.
- O `7` são as dezenas, `tens`.
- O `4` são as centenas, `hundreds`.

Então `2*100 + 7*10 + 4*1` vira `ones*100 + tens*10 + hundreds*1`. Essa é a passagem lógica → sintaxe que te trava: a lógica você já tinha. *Como pensar da próxima vez:* com o exemplo resolvido à mão, circule cada número e escreva do lado de onde ele veio. O código sai disso.

**[sintaxe] Constante e variável não são a mesma coisa.** No plano você escreveu "criar três constantes" e, nas dificuldades, "constante tens". No código você usou `:=`, que está certo. Mas vale acertar a palavra, porque ela muda o que o Go aceita:
- **Constante** tem valor conhecido antes de o programa rodar e nunca muda (`const secondsPerMinute = 60`).
- **`tens`** depende do `n`, que só existe quando a função é chamada, então é uma **variável**.

Se você escrever `const tens = n % 100 / 10`, o compilador recusa com `n % 100 / 10 (value of type int) is not constant`.

**[idiomático] Comentários.** O comentário do pacote apareceu, e isso é novo. Faltam ajustes de inglês:
- `LastDigit shows the last number of n` → **returns** the last **digit**. "Shows" dá a ideia de imprimir, a confusão do drill 001. E *number* é o número inteiro; *digit* é o algarismo.
- `DescribeDigits for n between 100 and 999 how many...` está sem verbo → `DescribeDigits returns how many hundreds, tens and ones n has.`
- `ReverseThreeDigits ... return the number` → **returns** (terceira pessoa, com "s").
- `Package digits reports tools to tell us about numbers` → `Package digits splits numbers into their decimal digits.` O comentário do pacote diz o que ele faz, numa frase.

## Solução de referência

```go
// Package digits splits numbers into their decimal digits.
package digits

import "fmt"

// LastDigit returns the last digit of n (n >= 0).
func LastDigit(n int) int {
	return n % 10
}

// DescribeDigits returns how many hundreds, tens and ones n has.
// n must be between 100 and 999.
func DescribeDigits(n int) string {
	hundreds := n / 100  // how many whole hundreds fit?
	tens := n % 100 / 10 // drop the hundreds, then how many whole tens fit?
	ones := n % 10       // what is left after the whole tens?
	return fmt.Sprintf("hundreds: %d, tens: %d, ones: %d", hundreds, tens, ones)
}

// ReverseThreeDigits returns n with its digits in reverse order: 472 -> 274.
// n must be between 100 and 999.
func ReverseThreeDigits(n int) int {
	hundreds := n / 100
	tens := n % 100 / 10
	ones := n % 10
	return ones*100 + tens*10 + hundreds // each digit goes to its new position
}
```

Por que assim:

- **O limite de entrada vai no comentário** ("n must be between 100 and 999"). Quem chama a função precisa saber, e o tipo `int` não diz isso.
- **Os três dígitos aparecem repetidos nas duas funções.** Por enquanto é o certo. No tópico de funções você vai escrever uma função que devolve os três de uma vez, e as duas passam a usar ela.

## Para levar

Do exemplo para o código: faça a conta com números, pergunte de onde veio cada número e troque pelo nome da variável que guarda ele. A lógica você já costuma ter; esse é o passo que transforma ela em Go.
