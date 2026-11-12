# Dicas — Drill 023

**Regra:** antes de abrir uma dica, escreva no código (num comentário) o que você tentou e onde travou. Abra um nível de cada vez.

## 1. Area

<details><summary>Nível 1 — direção</summary>

Olhe a linha "declarar um método que lê" na tabela de operações. A conta muda; o formato é o mesmo.
</details>

<details><summary>Nível 2 — estratégia</summary>

Sem loop e sem `if`: a área é largura vezes altura, e as duas estão nos campos do receiver `r`.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
devolver <largura de r> * <altura de r>
```
</details>

## 2. Fits

<details><summary>Nível 1 — direção</summary>

É o "escolher entre candidatos" do drill 006 e as condições combinadas do drill 012. Desenhe no papel o `{2 5}` entrando no `{6 3}`.
</details>

<details><summary>Nível 2 — estratégia</summary>

Há dois candidatos: **em pé** (largura com largura, altura com altura) e **deitado** (a altura de `r` vai na largura do `box`, e a largura de `r` na altura). Cada candidato é uma condição com duas comparações ligadas por `&&`. O retângulo cabe se **qualquer um** dos candidatos funcionar. Para ficar legível, guarde cada candidato numa variável `bool` com nome.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
standing = (largura de r <= largura do box) && (altura de r <= altura do box)
turned   = (altura de r <= largura do box) && (largura de r <= altura do box)
devolver standing || turned
```
</details>

## 3. Deposit

<details><summary>Nível 1 — direção</summary>

Duas perguntas: o receiver tem `*`? E o `if` guarda que ação? (Lembre do drill 019: "ignorar" não é ação.)
</details>

<details><summary>Nível 2 — estratégia</summary>

O método só **faz** algo quando o valor é positivo: somar ao saldo. Nos outros casos, não há nada a escrever. Não tem `return`, porque a assinatura não devolve nada.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
se amount > 0:
    saldo de a += amount
```
</details>

## 4. Withdraw

<details><summary>Nível 1 — direção</summary>

Liste os casos em que o saque **não** acontece. Depois, o que acontece no resto.
</details>

<details><summary>Nível 2 — estratégia</summary>

Dois casos barram o saque: o valor não é positivo, ou é maior que o saldo. Nesses casos, devolva `false` logo, sem tocar no saldo (retorno antecipado, do drill 006). Se passou pelos dois, o saque pode acontecer: tire do saldo e devolva `true`. Confira o caso da borda: sacar **exatamente** o saldo vale.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
se amount <= 0 || amount > saldo de a:
    devolver false
saldo de a -= amount
devolver true
```
</details>

## Erros comuns do compilador (e do teste)

| Mensagem | O que significa |
|---|---|
| `undefined: Rectangle` | Os tipos não estão no seu arquivo, ou estão com outro nome. Copie do README. |
| `r.Area undefined (type Rectangle has no field or method Area)` | Você escreveu uma **função** (`func Area(r Rectangle) int`), não um método. O receiver vai entre o `func` e o nome: `func (r Rectangle) Area() int`. |
| `after Deposit(50) on 100, Balance = 100, want 150` | O saldo não mudou: falta o `*` no receiver. O método mexeu numa cópia. |
| `Withdraw(30) on 100 = true with Balance 100, ...` | Devolveu `true`, mas o saldo não mudou: falta o `*` no receiver. |
| `after Deposit(-20) on 150, Balance = 130, want 150` | O valor negativo foi somado: falta o `if` que guarda a soma. |
| `{2 5}.Fits({6 3}) = false, want true (turned...)` | Só o candidato "em pé" foi testado. Falta o "deitado". |
| `{4 4}.Fits({3 6}) = true, want false` | Comparou as áreas. Área menor não garante que cabe: o 4 não entra no 3. |
| `{3 6}.Fits({3 6}) = false, want true (touching the border counts)` | Usou `<`; encostar na borda vale, então é `<=`. |
| `Withdraw(-5) on 70 = true with Balance 75, ...` | Faltou barrar o valor negativo: tirar -5 **soma** 5. |
| `Withdraw(70) on 70 = false with Balance 70, ...` | Usou `>=` no lugar de `>`: sacar exatamente o saldo vale. |
| `a.Balance + amount (value of type int) is not used` | Fez a conta mas não guardou. É `a.Balance += amount`. |
| `too many return values` | `return` com valor num método sem tipo de retorno (o `Deposit`). |
| `missing return` / `not enough return values` | No `Withdraw`, algum caminho termina sem `return true` ou `return false`. |
| `invalid operation: ... && r.Width (mismatched types untyped bool and int)` | Uma comparação pela metade dentro do `&&` ou do `\|\|`. Cada lado precisa ser uma comparação completa. |
| `declared and not used: turned` | Você guardou um candidato numa variável, mas não usou no `return`. |
