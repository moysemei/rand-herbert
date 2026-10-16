# Progresso

O tutor atualiza este arquivo. O campo `mode` você pode trocar à mão.

```
mode: normal              # normal | maintenance (semana de prova: só revisão de ~10 min)
dose: 20min               # fase de hábito; subir quando eu pedir
phase: 1 - Fundamentos
current_drill: 008-review-week-2
status: pending           # pending | done
```

## Histórico

| Drill | Tema | Criado | Entregue | Nota do tutor |
|---|---|---|---|---|
| 001-hello-world | Anatomia de um arquivo Go | 2026-10-06 | 2026-10-07 | Confundiu imprimir com devolver e fixou um valor no lugar do parâmetro; resolveu com perguntas guiadas. `Shout` reaproveitando `Greet` certo de primeira. |
| 002-variables-types | Variáveis, tipos, conversão, constantes | 2026-10-07 | 2026-10-08 | Sem dicas. Achou sozinho a armadilha da divisão inteira. Não usou `const` (tentou com valor calculado e `:=`); misturou `Sprintf` com `Itoa`. |
| 003-review-week-1 | Revisão de sexta: funções que usam funções, int vs float, const, Sprintf | 2026-10-09 | 2026-10-09 | Entregue no mesmo dia. Corrigiu `const` e "um caminho só". Plano com raciocínio e ajuste no meio. Comentários fora da convenção e falso cognato ("resume"). |
| 004-operators | `%`, comparação, `&&` `\|\|` `!` | 2026-10-12 | 2026-10-13 | `IsEven` e `IsLeapYear` verdes de primeira; comentários já começam pelo nome. Travou no `SecondsToClock` (4 rodadas de dicas): `total` nos slots do `Sprintf`, linha copiada três vezes, segundos com `% oneHour`. Faltou `func` (copiou a assinatura do README). Plano sem cálculo à mão. Anotou as dificuldades no arquivo. |
| 005-digits | Reforço de `/` e `%`: separar e remontar dígitos, `Sprintf` | 2026-10-13 | 2026-10-13 | Sem pedir ajuda no chat: dica 2 nas dezenas, dica 3 no inverso. Plano com a remontagem certa à mão (`2*100 + 7*10 + 4*1`), mas não passou dos números para as variáveis. Chama variável de "constante". Comentário do pacote apareceu. |
| 006-if-else | `if` / `else if` / `else`, retorno antecipado, ordem das condições (+ `%`) | 2026-10-14 | 2026-10-14 | `if` sem dificuldade: `Max` idiomático de primeira, ordem do `Classify` escolhida no plano. Achou sozinho o `=` no lugar de `==`. `RoundToTen` pediu dica 3 (desmontou 47 = 40 + 7, mas não viu os dois candidatos). Nenhum comentário. |
| 007-switch | `switch` com e sem expressão, ajudante não exportado | 2026-10-15 | 2026-10-15 | `DaysInMonth` sem nenhuma dica. No `Grade`, tentou `90 >= score <= 100` (comparação encadeada do Python) e vírgula como "e"; explicou o próprio bug antes das dicas 1 e 2. Diário de depuração nas "Dificuldades". Comentários de volta, mas com "reports" para tudo. Usou `switch {` com `month == ...` onde cabia `switch month`. |
| 008-review-week-2 | Revisão de sexta: `/` e `%`, candidatos, `if`/`switch`, constantes, função usando função | 2026-10-16 | — | — |

## Pontos de atenção

Fraquezas que aparecem nos reviews, para revisitar nas sextas.

- **Do exemplo para o código (lógica → sintaxe):** resolvido à mão, perguntar de onde veio cada número e trocar pelo nome da variável. No 005 o plano tinha a resposta com números. (drill 005)
- **Decompor com `/` e `%` — prioridade da revisão de sexta.** Cada valor da saída responde uma pergunta ("quantos inteiros cabem?" → `/`, "o que sobra?" → `%`). O pedaço do meio (minutos, dezenas) e o arredondamento pediram dica 2–3 em três drills seguidos. Conferir o intervalo do resultado. (drills 004, 005 e 006)
- **Escolher entre candidatos:** listar as respostas possíveis, achar a conta de cada uma, e só então o `if`. (drill 006, revisão no 008)
- **Hábitos do Python:** comparação encadeada (`90 <= x <= 100`) não existe em Go; cada comparação é separada e ligada com `&&`. (drill 007)
- **Escolher o tipo de `switch`:** um valor contra opções → `switch x`; faixas → `switch {`. Num `switch {`, preferir `||` à vírgula. (drill 007)
- **Calcular primeiro, formatar depois:** o `Sprintf` não calcula, só coloca valores prontos nos slots, na ordem. (drill 004)
- **Cálculo à mão no plano:** fazer o exemplo do README no papel antes de codar, a desmontagem também, não só a remontagem. No 004 teria pegado o `&&` errado do bissexto; no 007, o 101 do `Grade`. (drills 004, 005 e 007)
- **Vocabulário: variável vs constante.** Chamou `tens := ...` de "constante". Constante = valor conhecido antes de rodar. (drill 005)
- **Comentários:** ~~começar pelo nome~~ (feito no 004), ~~comentário do pacote~~ (feito no 005). Verbo na 3ª pessoa (*returns*, não *return*/*shows*); função `bool` usa "reports whether". No 006 não escreveu nenhum; no 007 voltaram, mas com "reports" em tudo. (drills 003 a 007)
- **Inglês nos comentários e nomes:** ortografia, falsos cognatos (*resume* ≠ resumir → *summarize*), *person* vs *people*, *a*/*an*, *into*. Constante com unidade no nome (`secondsPerHour`, não `oneHour`). (drills 002, 003 e 004)
- **Converter onde os tipos se encontram,** não "na saída". (drill 003)
- **Usar o parâmetro:** testar a função na cabeça com duas entradas diferentes antes de rodar. (drill 001)
- **`=` vs `==` em condição:** achou sozinho no 006 (a mensagem fala em *assignment*). (drill 006)
- ~~**Constante:** `const x = valor`, nunca `:=`.~~ Corrigido no 003.
- ~~**Montar texto:** um caminho só.~~ Corrigido no 003.
- ~~**Imprimir vs devolver.**~~ Não se repetiu desde o 001.
