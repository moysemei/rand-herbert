# Review — Drill 023

## Veredito

Entregue na sexta, com os quatro métodos certos e o refazer certo. O `Deposit` e o `Withdraw` saíram sozinhos, com o receiver certo. A trava foi no `Fits`, e não foi de Go: foi transformar a rotação em comparações.

## Processo

- **Você explicou a escolha do receiver com as suas palavras:** "o deposit tem que modificar o valor original. Area não precisa, ele só quer calcular". É exatamente a regra do drill, e dita por você, não copiada.
- **O `Withdraw` saiu sem orientação:** as duas condições ligadas por `&&`, o saldo só muda quando o saque é permitido, e cada `return` está no lugar certo.
- **No `Fits`, o caso em pé foi seu.** Você travou só na metade nova (a rotação), e a ajuda veio só nela.
- **Comentários em tudo, e o `Fits` começa com *reports whether*,** a convenção para quem devolve `bool`. Isso estava nos pontos de atenção desde o drill 007.
- **As "Dificuldades" desta vez foram só a avaliação do ChatGPT.** Ela foi útil e eu li tudo. Mas, como no 022, duas linhas com as suas palavras me dizem coisas que ela não diz. Mande os dois juntos.

## Para que servem métodos (a pergunta que o README não respondeu)

O README explicou o que é um método e como se escreve, mas não para que ele serve num sistema de verdade, e foi exatamente isso que você perguntou. Fica aqui:

- **Você já usa métodos todo dia.** Nos testes, `t.Errorf(...)`: o `t` é um `*testing.T`, e `Errorf` é um método dele com receiver por ponteiro.
- **Num backend, os dados carregam as perguntas e as ações deles:** `order.Total()`, `user.IsAdmin()` (só leem) e `cart.Add(item)` (muda o carrinho, então leva `*`). Quem lê `cart.Add(item)` entende o código sem procurar onde a função mora.
- **O motivo mais forte vem na fase 2: interfaces.** Em Go, um tipo "é" um `error` se tiver o método `Error() string`. O `err` que você confere depois do `Atoi` funciona assim. Métodos são o jeito de um tipo seu se encaixar em código que nunca ouviu falar dele.

A partir do próximo drill, toda ferramenta nova vem com essa parte no README.

## Erros e melhorias

**[lógica] Transformar uma situação em condições.** A sua estratégia no `Fits` estava certa: dois casos, cabe se qualquer um funcionar. O que travou foi escrever o caso deitado. O caminho é o mesmo que já destravou você no drill 008 (o `Price`) e no 022 (o `NewBook`): **comece pelos números do exemplo, e só depois troque cada número pela variável de onde ele veio.**

| Exemplo: `{2 5}` deitado no `{6 3}` | Com números | De onde vem cada número | Em Go |
|---|---|---|---|
| o lado 5 vai na largura da caixa | `5 <= 6` | 5 é a altura de `r`; 6 é a largura do `box` | `r.Height <= box.Width` |
| o lado 2 vai na altura da caixa | `2 <= 3` | 2 é a largura de `r`; 3 é a altura do `box` | `r.Width <= box.Height` |

A coluna do meio é o passo que faltou. Geometria, regras de negócio, horários: qualquer "situação" vira condição assim.

**[idiomático] `if condição { return true }; return false` é só `return condição`.** O seu `Fits` funciona, mas os dois `if` existem só para devolver o que a condição já é. Guarde cada caso numa variável `bool` com nome (`standing`, `turned`) e devolva `standing || turned`. O código fica do tamanho da ideia: "cabe em pé ou deitado".

## Solução de referência

```go
// Package methods practices value and pointer receivers.
package methods

// Rectangle is a rectangle with integer sides.
type Rectangle struct {
	Width  int
	Height int
}

// Account is a bank account; Balance is in reais.
type Account struct {
	Owner   string
	Balance int
}

// Area returns the area of r.
// Value receiver: it only reads.
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// Fits reports whether r fits inside box, standing or turned 90 degrees.
// Pattern: choose between candidates; each candidate is one full condition.
func (r Rectangle) Fits(box Rectangle) bool {
	standing := r.Width <= box.Width && r.Height <= box.Height
	turned := r.Height <= box.Width && r.Width <= box.Height
	return standing || turned
}

// Deposit adds amount to the balance; zero or negative amounts are ignored.
// Pointer receiver: it changes the caller's account.
func (a *Account) Deposit(amount int) {
	if amount > 0 {
		a.Balance += amount
	}
}

// Withdraw takes amount from the balance and reports whether it happened.
// Guard first: reject the cases that block the withdrawal, then do it.
func (a *Account) Withdraw(amount int) bool {
	if amount <= 0 || amount > a.Balance {
		return false
	}
	a.Balance -= amount
	return true
}
```

Por que assim:

- **`standing` e `turned` com nome:** cada candidato vira uma frase lida em voz alta, e o `return` diz a regra inteira.
- **O `Withdraw` com guarda:** o seu (condição boa dentro do `if`, `return false` no fim) está igualmente certo. A versão com guarda é a mais comum em Go: os casos que barram saem primeiro, e o caminho normal fica sem indentação. Quando os casos que barram forem três ou quatro, a diferença aparece.

## Refazer de memória (Unique)

Certo: um `seen` para "já vi?" e um `result` para montar a saída, na ordem certa. No 020, a parte central dessa função veio do ChatGPT, e agora ela saiu inteira.

## Para levar

Método que só lê leva receiver por valor; método que muda a struct leva `*`. E para transformar uma situação em condição, comece pelos números do exemplo e troque cada um pela variável de onde ele veio.
