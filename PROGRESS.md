# Progresso

O tutor atualiza este arquivo. O campo `mode` você pode trocar à mão.

```
mode: normal              # normal | maintenance (semana de prova: só revisão de ~10 min)
dose: 45min               # desde 2026-10-26; 1h30 quando eu confirmar (ver CLAUDE.md)
phase: 1 - Fundamentos
current_drill: 015-strings-runes
status: pending           # pending | done
```

## Agenda

- **Prova:** 2026-12-05 (a confirmar). Semana de `maintenance`: 2026-11-30 a 2026-12-04. Volta ao normal em 2026-12-07.

## Dose 45 min: acompanhamento

Critério para propor 1h30: duas semanas seguidas com o drill entregue em pelo menos 4 dos 5 dias úteis.

- Semana 1 (26/10–30/10): seg ✓ (014)

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
| 008-review-week-2 | Revisão de sexta: `/` e `%`, candidatos, `if`/`switch`, constantes, função usando função | 2026-10-16 | 2026-10-18 | Sem ajuda. `Price` saiu direto do plano com números ("as outras 2 horas" → `hours - 1`): a técnica do 005 funcionou. `BilledHours` com o truque `(m + 59) / 60`, sem explicação no plano. Constantes com unidade no nome; uma exportada sem querer. Comentários com "shows"/"calculates" e "based in". |
| 009-for | `for` clássico, acumulador, valor inicial, limites | 2026-10-19 | 2026-10-19 | Travou: dica 1 no `SumTo`, dica 3 no `Factorial`; `SumOfMultiples` resolvido com ajuda de IA (anotado por ele). Raiz: usava o `n` dentro da conta (`total += i + n`, `n * (n-1)`), sem ver o `i` como "o número da vez". Ele mesmo nomeou a confusão. Plano incompleto, sem comentários. |
| 010-for-practice | Reforço de `for`: papel de cada variável, contar vs somar, tabela de rastreio | 2026-10-20 | 2026-10-20 | Grande melhora: `SumEven` sem dica, `CountMultiples` de primeira (contou, não somou), `Power` com dica 2 (tentava escrever a conta inteira numa linha). Pulou as tabelas ("vou tendo ideia e passo pra prática"): plano passa a ser 1 linha por função, tabela quando travar. Sem espaço depois do `//`. |
| 011-for-while | `for` como while, percorrer dígitos, caso do zero, `break` (intro) | 2026-10-21 | 2026-10-21 | Entregue antes das 15h, com `go fmt`. Caiu no loop infinito e saiu sozinho. `ReverseNumber` com `digit :=` bem nomeado, mas remontar o número foi difícil (sem tabela no plano). `CountDigits(0)` resolvido com `count++` antes do loop (certo, mas "ajambrado" nas palavras dele). Acumulador chamado `f`. |
| 012-combined-conditions | Condições combinadas, padrão busca, `break`/`continue`/`return` no loop | 2026-10-22 | 2026-10-22 | Código limpo: ordem certa no `FizzBuzz`, "não achei" depois do loop, `candidate := n + 1`. Usou `d <= n/d` (não estava no material, sem explicação no plano). Sem comentários, sem contas à mão e sem "Dificuldades": pedir a explicação e o registro. |
| 013-review-week-3 | Revisão de sexta: loops, contar vs somar, padrão busca, cópia do original, ajudantes de memória | 2026-10-23 | 2026-10-25 | Sem erros. Resolveu de uma vez vários pontos de atenção: comentários com verbo certo em tudo, nomes claros (`original`, `reversed`), cópia antes do loop, contas à mão no plano. `d <= n/d` de novo sem explicação; sem "Dificuldades" pelo 2º drill seguido. |
| 014-multiple-returns | Múltiplos retornos, retornos nomeados, vírgula ok (+ refazer: `Power`) | 2026-10-26 | 2026-10-26 | Entrega mais completa até aqui: refazer certo, tabela do `MinMaxDigit` certa, explicação do `d <= n/d`, "Dificuldades" detalhadas (queria guardar todos os dígitos; entendeu os dois acumuladores). Usou `(n/10)%10` do review 005. `=` vs `==` de novo (corrigiu sozinho). Variáveis `min`/`max` escondem as funções embutidas. |
| 015-strings-runes | Strings: bytes vs runes, `len`, indexação, `for range` (+ refazer: `ReverseNumber`) | 2026-10-27 | — | — |

## Pontos de atenção

Fraquezas que aparecem nos reviews, para revisitar nas sextas.

- **Do exemplo para o código (lógica → sintaxe):** resolvido à mão, perguntar de onde veio cada número e trocar pelo nome da variável. No 005 o plano tinha a resposta com números; no 008 o `Price` saiu assim, sem dica. Melhorando. (drills 005 e 008)
- **Papel de cada variável no loop — prioridade da revisão de sexta.** `i` = número da vez; limite = só o teto; acumulador = nasce antes, cresce dentro; valor fixo (`base`) entra em toda volta. Melhorou muito no 010; ainda tenta escrever a conta inteira numa linha em vez de "um pedaço por volta". (drills 009 e 010)
- **Remontar número / fórmula com acumulador** (`rev*10 + digit`, `product *= base`): a parte que ainda pede ajuda nos loops. (drills 010 e 011)
- ~~**Nome de variável:** `f`, `result` para contagem.~~ Resolvido no 013. (drills 010 e 011)
- **Plano:** ele pula tabelas longas. Formato combinado a partir do 011: uma linha por função (muda / fixo / inicial / quando para) e tabela só quando travar, antes da dica. (drill 010)
- **Pedir ajuda na ordem certa:** dicas → chat com o tutor → IA só para dica, sem código. No 009 a parte 3 veio pronta. (drill 009)
- ~~**Só usar o que consegue explicar / registrar as "Dificuldades":**~~ Explicou o `d <= n/d` e voltou a escrever as "Dificuldades" no 014. Manter no radar. (drills 008, 012, 013 e 014)
- **Pensar em "guardar tudo" (reflexo do Python):** no 014 quis guardar todos os dígitos para achar min/max. Treinar *single pass* com acumuladores. (drill 014)
- **Nomes que escondem funções embutidas:** `min`, `max` (Go 1.21+), `len`. (drill 014)
- **Decompor com `/` e `%`.** Cada valor da saída responde uma pergunta ("quantos inteiros cabem?" → `/`, "o que sobra?" → `%`). Pediu dica 2–3 nos drills 004–006; no 008 (revisão) saiu sem dica. Manter no radar. (drills 004 a 008)
- **Escolher entre candidatos:** listar as respostas possíveis, achar a conta de cada uma, e só então o `if`. (drill 006, revisão no 008)
- **Hábitos do Python:** comparação encadeada (`90 <= x <= 100`) não existe em Go; cada comparação é separada e ligada com `&&`. (drill 007)
- **Escolher o tipo de `switch`:** um valor contra opções → `switch x`; faixas → `switch {`. Num `switch {`, preferir `||` à vírgula. (drill 007)
- **Calcular primeiro, formatar depois:** o `Sprintf` não calcula, só coloca valores prontos nos slots, na ordem. (drill 004)
- **Cálculo à mão no plano:** fazer o exemplo do README no papel antes de codar, a desmontagem também, não só a remontagem. No 004 teria pegado o `&&` errado do bissexto; no 007, o 101 do `Grade`. (drills 004, 005 e 007)
- **Vocabulário: variável vs constante.** Chamou `tens := ...` de "constante". Constante = valor conhecido antes de rodar. (drill 005)
- **Comentários:** ~~começar pelo nome~~ (feito no 004), ~~comentário do pacote~~ (feito no 005). Verbo na 3ª pessoa (*returns*, não *return*/*shows*); função `bool` usa "reports whether". No 006 não escreveu nenhum; no 007 voltaram, mas com "reports" em tudo; no 008, "shows"/"calculates"; no 009, nenhum de novo; no 010 voltaram, sem espaço depois do `//` e sem verbo; no 011, formatação certa (`go fmt`), mas ainda "reports" para `int` e inglês travado ("explit", "n have"); no 012, nenhum de novo; ~~no 013, todos certos~~ (resolvido no 013, manter no radar). *based on*, não *based in* (2x). (drills 003 a 013)
- **Inglês nos comentários e nomes:** ortografia, falsos cognatos (*resume* ≠ resumir → *summarize*), *person* vs *people*, *a*/*an*, *into*, *based on*. ~~Constante com unidade no nome~~ (feito no 008). (drills 002 a 008)
- **Maiúscula = exportado** vale para constantes também: no 008 uma constante interna ficou exportada. (drill 008)
- **Variável só para devolver** (`x := ...; return x`) quando o nome não acrescenta nada. (drills 003 e 008)
- **Converter onde os tipos se encontram,** não "na saída". (drill 003)
- **Usar o parâmetro:** testar a função na cabeça com duas entradas diferentes antes de rodar. (drill 001)
- **`=` vs `==` em condição:** achou sozinho no 006 e no 014 (a mensagem fala em *assignment*). (drills 006 e 014)
- ~~**Constante:** `const x = valor`, nunca `:=`.~~ Corrigido no 003.
- ~~**Montar texto:** um caminho só.~~ Corrigido no 003.
- ~~**Imprimir vs devolver.**~~ Não se repetiu desde o 001.
