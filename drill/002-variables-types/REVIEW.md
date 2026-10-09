# Review — Drill 002

## Veredito

Entregue, os três testes passaram, sem abrir nenhuma dica. Faltou só a constante pedida nas regras do dia.

## Processo

- **Você passou sozinho pela armadilha da divisão inteira.** Você tentou `float64(minutes / 60)`, viu o teste falhar e entendeu que precisava converter **antes** de dividir. Era a parte do drill feita pra exigir raciocínio.
- **Você pesquisou o `strconv.Itoa` na documentação.** É exatamente o que um dev faz no dia a dia.
- **Você foi honesto sobre ter consultado o review do 001** para lembrar do `fmt.Sprintf`. Isso não é colar: consultar código que você já viu e a documentação faz parte do trabalho. Colar seria copiar uma solução pronta deste drill.
- **Você aplicou o review anterior:** comentário em todas as funções exportadas, nome de arquivo certo, e um plano que descreve o que cada função faz, não o que ela "imprime".
- **Você anotou as dificuldades de novo.** Foi por essas anotações que eu soube exatamente onde estava a dúvida sobre `const`.

## Erros e melhorias

**[sintaxe] `const temp := c * 9 / 5 + 32`.** Você anotou que tentou isso e não deu. São dois problemas:
- *Sintaxe:* constante usa `=`, nunca `:=`. O compilador diz `syntax error: unexpected :=, expected =`.
- *Conceito:* mesmo com `=`, não funcionaria. Uma constante precisa ter o valor conhecido **quando você escreve o código**, e `c` só chega quando alguém chama a função. O erro seria `c * 9 / 5 (value of type float64) is not constant`.
- *Por que você provavelmente pensou assim:* em outras linguagens (como o `const` do JavaScript), "constante" quer dizer "variável que não muda depois de criada", e aí pode receber qualquer valor calculado. Em Go, constante é um valor fixo do programa: `32`, `60`, `"Hello"`.
- *Como pensar:* pergunte "esse valor depende da entrada?". Se depende, é variável (`:=`). Se é sempre o mesmo, como os 60 minutos de uma hora ou os 32 da fórmula, é constante. Esses dois números eram as constantes do drill.

**[idiomático] `Sprintf` e `Itoa` juntos:** `fmt.Sprintf("%s is %v years old", name, strconv.Itoa(age))` funciona, mas faz a conversão duas vezes. Escolha um dos caminhos:
- `fmt.Sprintf("%s is %d years old", name, age)`: o `Sprintf` já sabe formatar `int` com `%d`.
- `name + " is " + strconv.Itoa(age) + " years old"`: converte com `Itoa` e junta com `+`.

*Por que você provavelmente fez assim:* você tinha duas peças que funcionavam (uma do review, outra da pesquisa) e usou as duas, por garantia. Vale aprender os "verbos" mais comuns do `fmt`: `%s` para string, `%d` para inteiro, `%v` para qualquer coisa no formato padrão, `%q` para string entre aspas. `%v` funciona com string, mas `%s` diz melhor a intenção.

**[idiomático] Nomes.** `temp` em inglês costuma ser lido como "temporário", não "temperatura". `fahrenheit` ou só `f` seria mais claro. Na verdade nem precisava de variável: dá pra fazer `return` direto da conta.

**[idiomático] Inglês nos comentários.** "tempture" → *temperature*, "convertes" → *converts*, "persona" → *person*. Parece detalhe, mas você mira vagas lá fora, e comentário em PR é lido por gente que fala inglês. O comentário do pacote também ficaria melhor como frase completa: `// Package convert provides simple unit and text conversions.`

## Solução de referência

```go
// Package convert provides simple unit and text conversions.
package convert

import "strconv"

const (
	freezingPointF = 32 // water freezes at 32 °F
	minutesPerHour = 60
)

// CelsiusToFahrenheit converts a temperature from Celsius to Fahrenheit.
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + freezingPointF
}

// MinutesToHours converts minutes to hours, keeping the fraction (90 -> 1.5).
func MinutesToHours(minutes int) float64 {
	return float64(minutes) / minutesPerHour
}

// Describe returns a sentence like "Ada is 36 years old".
func Describe(name string, age int) string {
	return name + " is " + strconv.Itoa(age) + " years old"
}
```

Por que assim:

- **`const ( ... )`** agrupa constantes relacionadas, igual ao `import ( ... )`. Os nomes explicam o que os números são; um `60` solto no meio da conta não diz nada (é o que se chama de "número mágico").
- **`float64(minutes) / minutesPerHour`** funciona sem converter a constante, porque constante sem tipo se adapta ao `float64` do outro lado. Se `minutesPerHour` fosse uma variável `int`, daria erro de tipos misturados.
- **No `Describe`**, `+` com `Itoa` e `Sprintf` com `%d` são ambos idiomáticos. Para frases com várias partes, o `Sprintf` costuma ficar mais legível.

## Para levar

Constante é um valor fixo que você conhece ao escrever o código; se depende da entrada, é variável. E, para formatar texto, escolha uma ferramenta só: `Sprintf` com o verbo certo, ou `+` com `strconv`.
