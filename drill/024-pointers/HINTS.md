# Dicas — Drill 024

**Regra:** antes de abrir uma dica, escreva no código (num comentário) o que você tentou e onde travou. Abra um nível de cada vez.

## 1. Double

<details><summary>Nível 1 — direção</summary>

Olhe a tabela "o `*` tem dois sentidos". Dentro da função, você quer o sentido de **expressão**.
</details>

<details><summary>Nível 2 — estratégia</summary>

`n` é o endereço; o número está **no** endereço. Leia o valor que está lá, multiplique por 2 e escreva de volta no mesmo lugar. Sem `return`: quem chamou vê a mudança pela própria variável.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
o valor no endereço n = o valor no endereço n * 2
```
(ou o atalho `*=`, aplicado ao valor, não ao endereço)
</details>

## 2. Swap

<details><summary>Nível 1 — direção</summary>

Você já trocou dois valores numa linha só no `ReverseInPlace` (drill 017). Agora os dois lados são valores **em endereços**.
</details>

<details><summary>Nível 2 — estratégia</summary>

Troque o que está nos endereços, não os endereços. A troca em uma linha (`x, y = y, x`) funciona aqui também, com o `*` nos quatro lugares. Duas linhas separadas (`*a = *b` e depois `*b = *a`) não funcionam: na segunda, o valor antigo de `*a` já sumiu.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
valor em a, valor em b = valor em b, valor em a
```
</details>

## 3. AddBonus

<details><summary>Nível 1 — direção</summary>

Releia a armadilha "a variável do `range` é uma cópia". O padrão é **percorrer e mudar cada item**; o difícil é chegar no item de verdade.
</details>

<details><summary>Nível 2 — estratégia</summary>

O slice enxerga os itens de quem chamou, então mudar `accounts[i]` muda o original. O que não funciona é mudar a cópia que o `range` entrega. Percorra pelos **índices** e mude o saldo do item naquele índice.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
para cada índice i do slice:
    saldo de accounts[i] += bonus
```
(`for i := range accounts` entrega só os índices)
</details>

## 4. FindByOwner

<details><summary>Nível 1 — direção</summary>

Padrão do cardápio: **buscar**. Posso responder antes de o loop terminar? E o "não achei" fica onde?
</details>

<details><summary>Nível 2 — estratégia</summary>

Percorra pelos índices. Quando o dono do item for o procurado, devolva o **endereço do item dentro do slice** (não o endereço da variável do `range`, que é uma cópia). O `nil` só pode ser devolvido **depois** do loop: dentro dele, um `else` devolveria `nil` já no primeiro item diferente.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
para cada índice i:
    se o dono de accounts[i] == owner:
        devolver &accounts[i]
devolver nil
```
</details>

## Erros comuns do compilador (e do teste)

| Mensagem | O que significa |
|---|---|
| `invalid operation: n *= 2 (mismatched types *int and untyped int)` | Você fez conta com o **endereço**. A conta é com o valor: `*n`. |
| `*n * 2 (value of type int) is not used` | Calculou o dobro, mas não guardou em lugar nenhum. Falta escrever no endereço. |
| `after Swap(...): x = 1, y = 2, want x = 2, y = 1` | Trocou os endereços (cópias locais), não os valores. Falta o `*`. |
| `after Swap(...): x = 2, y = 2, ...` | Trocou em duas linhas: o valor antigo se perdeu na primeira. Use a troca numa linha só. |
| `after AddBonus(accs, 10), accs = [{... Balance:100} ...]` | Mudou a cópia do `range`. Mude pelo índice. |
| `after p.Balance += 5, accs[1].Balance = 50, want 55` | Devolveu o endereço da variável do `range` (uma cópia). Devolva `&accounts[i]`. |
| `FindByOwner(accs, "Caio") = {Owner: Balance:0}, want nil` | Devolveu uma conta vazia. "Não achei" é `nil`. |
| `FindByOwner(accs, "Bia") = nil, want a pointer...` | Um `else { return nil }` dentro do loop desiste no primeiro item diferente. O `nil` vai depois do loop. |
| `cannot use accounts[i] (variable of struct type Account) as *Account value` | Devolveu a conta, e a assinatura pede o endereço dela: falta o `&`. |
| `missing return` | Falta o `return nil` no fim, para quando o loop termina sem achar. |
| `panic: ... nil pointer dereference` | Algum ponteiro `nil` foi usado. Confira se a função devolve um endereço quando acha. |
