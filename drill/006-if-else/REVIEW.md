# Review — Drill 006

## Veredito

Entregue no mesmo dia, os três testes passam. O `if` em si não te travou: `Max` e `Classify` saíram sozinhos. O que pediu a dica 3 foi de novo a conta do `RoundToTen`, e por isso ela ganha a explicação mais longa aqui.

## Processo

- **O `Max` já saiu no jeito idiomático,** com retorno antecipado e sem `else`. Você leu o "Por que é assim" e aplicou.
- **Você escolheu a ordem do `Classify` antes de codar,** e escolheu uma ordem diferente da que as dicas sugerem (zero, depois negativo). Está igualmente certa: zero e negativo nunca acontecem juntos, então tanto faz qual vem primeiro, desde que os dois venham antes do "é par".
- **O `n > 0` do plano sumiu no código, e com razão:** depois dos dois `return`, só sobra número positivo, então perguntar de novo seria redundante.
- **Você achou o próprio bug.** Nas "Dificuldades": "era pq eu tava colocando `n%2 = 0` e não `n%2 == 0`". Isolar o erro sozinho é a habilidade mais importante desta fase.
- **No plano do `RoundToTen` você desmontou o número sozinho:** `47 = 4*10 + 7*1`. É o review do 005 virando hábito.
- **`lastDigit` é um bom nome:** diz o que a variável guarda, sem precisar de comentário.
- **Desta vez não teve nenhum comentário,** nem no pacote nem nas funções. O checklist pedia. Fica registrado, sem drama: no 007 eles voltam.

## Erros e melhorias

**[sintaxe] `=` no lugar de `==`.** Você já corrigiu, então fica só o registro para reconhecer na próxima vez. Dentro de um `if`, o Go mostra `syntax error: cannot use assignment (n % 2) = 0 as value`. A palavra **assignment** (atribuição) na mensagem é a pista: em algum lugar tem um `=` que deveria ser `==`.

**[lógica] O `RoundToTen`.** Você disse que mesmo com a dica 3 não tinha certeza de ter entendido, então vamos devagar, partindo do seu próprio plano: `47 = 4*10 + 7*1`, ou seja, **47 = 40 + 7**.

Arredondar é **escolher entre dois candidatos**, os múltiplos de 10 vizinhos:

```
40 ----------------- 47 ----- 50
↑ down                        ↑ up
```

1. **Quais são os candidatos?** Para 47: 40 (o de baixo) e 50 (o de cima).
2. **Como chegar em cada um a partir de 47?**
   - **40:** tire o 7. É o `47 - 7`.
   - **50:** o de baixo mais 10. É o `40 + 10`.
3. **O que escolhe entre eles?** O último dígito. 7 é 5 ou mais, então 50.

Agora a técnica do 005: de onde veio cada número?
- O **7** é o último dígito: `n % 10`, o seu `lastDigit`.
- O **40** é `n - lastDigit`.
- O **50** é `n - lastDigit + 10`.

O seu código é exatamente isso. Com 43: o último dígito é 3, os candidatos são 40 e 50, e como 3 é menor que 5, sai 40.

*Por que provavelmente travou:* você desmontou o número (passo certo), mas não tinha a pergunta seguinte, "quais são as respostas possíveis?". *Como pensar da próxima vez:* quando o problema é "escolher", liste os candidatos primeiro, descubra a conta de cada um, e só no fim escreva o `if` que escolhe.

Um bônus para ligar com o que você já sabe: o múltiplo de baixo também sai com divisão inteira, `n / 10 * 10`. O `47 / 10` dá `4` (a parte decimal vai embora) e `4 * 10` dá `40`.

**[idiomático] Comentários.** O pacote e as três funções ficaram sem comentário. O formato é o mesmo de sempre: `// Max returns the larger of a and b.`

## Solução de referência

```go
// Package decide chooses between paths with if and else.
package decide

// Max returns the larger of a and b.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Classify returns "negative", "zero", "positive even" or "positive odd".
func Classify(n int) string {
	if n < 0 {
		return "negative"
	}
	if n == 0 {
		return "zero"
	}
	// From here on, n is positive.
	if n%2 == 0 {
		return "positive even"
	}
	return "positive odd"
}

// RoundToTen rounds n (n >= 0) to the nearest multiple of 10.
// A last digit of 5 rounds up.
func RoundToTen(n int) int {
	lastDigit := n % 10
	down := n - lastDigit // the multiple of 10 below: 47 -> 40
	up := down + 10       // the multiple of 10 above: 47 -> 50
	if lastDigit >= 5 {
		return up
	}
	return down
}
```

Por que assim:

- **`down` e `up` dão nome aos dois candidatos,** e com isso o `if` se lê como a regra em português: "se o último dígito for 5 ou mais, sobe; senão, desce". A sua versão, com as contas direto no `return`, está certa e também é comum. Os nomes valem a pena quando a conta sozinha não explica o que ela é.
- **O comentário "From here on, n is positive"** registra o que os `return` de cima garantem. Isso ajuda quem lê o código depois, inclusive você daqui a um mês.

## Para levar

Muito problema de lógica é escolher entre candidatos. Liste as respostas possíveis, descubra a conta de cada uma a partir do exemplo e, por último, escreva o `if` que escolhe.
