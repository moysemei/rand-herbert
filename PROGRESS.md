# Progresso

O tutor atualiza este arquivo. O campo `mode` você pode trocar à mão.

```
mode: normal              # normal | maintenance (semana de prova: só revisão de ~10 min)
dose: 20min               # fase de hábito; subir quando eu pedir
phase: 1 - Fundamentos
current_drill: 004-operators
status: pending           # pending | done
```

## Histórico

| Drill | Tema | Criado | Entregue | Nota do tutor |
|---|---|---|---|---|
| 001-hello-world | Anatomia de um arquivo Go | 2026-10-06 | 2026-10-07 | Confundiu imprimir com devolver e fixou um valor no lugar do parâmetro; resolveu com perguntas guiadas. `Shout` reaproveitando `Greet` certo de primeira. |
| 002-variables-types | Variáveis, tipos, conversão, constantes | 2026-10-07 | 2026-10-08 | Sem dicas. Achou sozinho a armadilha da divisão inteira. Não usou `const` (tentou com valor calculado e `:=`); misturou `Sprintf` com `Itoa`. |
| 003-review-week-1 | Revisão de sexta: funções que usam funções, int vs float, const, Sprintf | 2026-10-09 | 2026-10-09 | Entregue no mesmo dia. Corrigiu `const` e "um caminho só". Plano com raciocínio e ajuste no meio. Comentários fora da convenção e falso cognato ("resume"). |
| 004-operators | `%`, comparação, `&&` `\|\|` `!` | 2026-10-12 | — | — |

## Pontos de atenção

Fraquezas que aparecem nos reviews, para revisitar nas sextas.

- **Comentários:** começam pelo nome da função (`// TotalSlices returns ...`) e o pacote também tem o seu. (drill 003)
- **Inglês nos comentários e nomes:** ortografia, falsos cognatos (*resume* ≠ resumir → *summarize*), *person* vs *people*. (drills 002 e 003)
- **Converter onde os tipos se encontram,** não "na saída". (drill 003)
- **Usar o parâmetro:** testar a função na cabeça com duas entradas diferentes antes de rodar. (drill 001)
- ~~**Constante:** `const x = valor`, nunca `:=`.~~ Corrigido no 003.
- ~~**Montar texto:** um caminho só.~~ Corrigido no 003.
- ~~**Imprimir vs devolver.**~~ Não se repetiu desde o 001.
