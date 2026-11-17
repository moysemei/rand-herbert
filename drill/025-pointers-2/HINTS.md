# Dicas — Drill 025

**Regra:** antes de abrir uma dica, escreva no código (num comentário) o que você tentou e onde travou. Abra um nível de cada vez.

## 1. Abs

<details><summary>Nível 1 — direção</summary>

É uma função que **devolve** (coluna da esquerda da tabela "dois jeitos"). E é um "escolher entre candidatos": quais são as respostas possíveis?
</details>

<details><summary>Nível 2 — estratégia</summary>

Dois candidatos: se o número é negativo, a resposta é ele com o sinal trocado; senão, é ele mesmo. Trocar o sinal é colocar um `-` na frente.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
se n < 0: devolver -n
devolver n
```
</details>

## 2. MakeAbs

<details><summary>Nível 1 — direção</summary>

É uma função que **escreve** (coluna da direita). E a pergunta do problema: já existe uma função sua que faz a conta?
</details>

<details><summary>Nível 2 — estratégia</summary>

Desça até o número (`*n`), passe esse número para o `Abs`, e escreva o resultado de volta no mesmo endereço. Sem `return`. (Também vale fazer sem o `Abs`, com um `if` sobre `*n`; mas aí você está escrevendo a mesma conta duas vezes.)
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
o valor no endereço n = Abs(o valor no endereço n)
```
</details>

## 3. Push

<details><summary>Nível 1 — direção</summary>

O topo é o fim do slice `items`. Que operação de slice põe um item no fim?
</details>

<details><summary>Nível 2 — estratégia</summary>

Acrescente `r` no fim de `s.items` e guarde o resultado de volta em `s.items` (o `append` devolve o slice novo). O receiver `*` é o que faz a mudança ficar.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
s.items = append(s.items, r)
```
</details>

## 4. Pop

<details><summary>Nível 1 — direção</summary>

Três perguntas, em ordem: está vazia? Qual é o topo? Como fica o slice sem ele? A terceira é o fatiar do drill 019.
</details>

<details><summary>Nível 2 — estratégia</summary>

Se `items` está vazio, devolva `0, false` logo (guarda). Senão: o índice do topo é `len(items) - 1`; guarde o valor desse índice numa variável; troque `items` pela janela que vai do começo até **antes** do topo; devolva o valor guardado com `true`. A ordem importa: guarde o topo **antes** de encolher o slice.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
se len(s.items) == 0: devolver 0, false
last = len(s.items) - 1
top = s.items[last]
s.items = s.items[:last]
devolver top, true
```
</details>

## Extra. IsBalanced

<details><summary>Nível 1 — direção</summary>

Cada rune é de um de três tipos: abre, fecha, ou outra coisa. O que a pilha guarda: quem abre ou quem fecha?
</details>

<details><summary>Nível 2 — estratégia</summary>

Quem abre vai para a pilha. Quem fecha precisa encontrar no **topo** o seu par de abertura: tire o topo com `Pop`; se a pilha estava vazia, ou se o topo não é o par certo, já é `false`. Um map `fecha → abre` (`')'` → `'('`) diz qual é o par. No fim do texto, a pilha tem que estar **vazia**: se sobrou alguém, ficou aberto.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
pairs = map: ')' → '(', ']' → '[', '}' → '{'
open = pilha vazia
para cada rune r de s:
    se r abre: open.Push(r)
    se r fecha:
        top, ok = open.Pop()
        se !ok ou top != pairs[r]: devolver false
no fim: devolver "a pilha está vazia?" (um Pop que devolve ok == false)
```
(um `switch r` com `case '(', '[', '{':` e `case ')', ']', '}':` deixa os três tipos bem separados)
</details>

## Erros comuns do compilador (e do teste)

| Mensagem | O que significa |
|---|---|
| `cannot use n (variable of type *int) as int value in argument to Abs` | Passou o endereço para quem quer o número. Desça com `*n`. |
| `cannot use Abs(*n) (value of type int) as *int value in assignment` | Escreveu `n = ...`, que troca o endereço. Para escrever no endereço, é `*n = ...`. |
| `invalid operation: operator - not defined on n (variable of type *int)` | `-n` com `n` endereço. O sinal vai no número: `-*n`. |
| `too many return values` | `return` com valor no `MakeAbs`, que não devolve nada. |
| `after Push('a'), Push('b'), items = [], want ['a' 'b']` | O `Push` mexeu numa cópia: falta o `*` no receiver. |
| `2nd Pop() = '[', true, want '(', true` | O topo não saiu: falta o `*` no receiver do `Pop`, ou falta encolher o `items`. |
| `1st Pop() ... = '(', true, want '[', true` | Tirou do começo (fila). Numa pilha, sai o último que entrou: o fim do slice. |
| `panic: runtime error: index out of range [-1]` | Leu o topo de uma pilha vazia. A guarda "está vazia?" vem antes. |
| `invalid argument: index -1 ... must not be negative` | `items[-1]` é Python. Em Go, o último é `items[len(items)-1]`. |
| `cannot use nil as rune value in return statement` | Rune não tem `nil`. O valor zero de rune é `0`. |
| `IsBalanced("((") = true, want false (left open)` | Faltou conferir, no fim, se a pilha ficou vazia. |
