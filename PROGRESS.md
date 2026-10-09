# Progresso

O tutor atualiza este arquivo. O campo `mode` você pode trocar à mão.

```
mode: normal              # normal | maintenance (semana de prova: só revisão de ~10 min)
dose: 20min               # fase de hábito; subir quando eu pedir
phase: 1 - Fundamentos
current_drill: 003-review-week-1
status: pending           # pending | done
```

## Histórico

| Drill | Tema | Criado | Entregue | Nota do tutor |
|---|---|---|---|---|
| 001-hello-world | Anatomia de um arquivo Go | 2026-10-06 | 2026-10-07 | Confundiu imprimir com devolver e fixou um valor no lugar do parâmetro; resolveu com perguntas guiadas. `Shout` reaproveitando `Greet` certo de primeira. |
| 002-variables-types | Variáveis, tipos, conversão, constantes | 2026-10-07 | 2026-10-08 | Sem dicas. Achou sozinho a armadilha da divisão inteira. Não usou `const` (tentou com valor calculado e `:=`); misturou `Sprintf` com `Itoa`. |
| 003-review-week-1 | Revisão de sexta: funções que usam funções, int vs float, const, Sprintf | 2026-10-09 | — | — |

## Pontos de atenção

Fraquezas que aparecem nos reviews, para revisitar nas sextas.

- **Constante:** `const x = valor`, nunca `:=`; só para valores fixos, que não dependem da entrada. (drill 002)
- **Montar texto:** um caminho só, `Sprintf` com o verbo certo (`%d`, `%s`) **ou** `+` com `strconv.Itoa`. (drill 002)
- **Inglês nos comentários e nomes:** revisar a ortografia; nomes que digam o que o valor é. (drill 002)
- **Usar o parâmetro:** testar a função na cabeça com duas entradas diferentes antes de rodar. (drill 001)
- **Imprimir vs devolver:** função devolve com `return`; `fmt.Println` é só para mostrar a um humano. (drill 001, não se repetiu no 002)
- **Convenções:** comentário em toda função exportada; arquivo `.go` com nome minúsculo e sem hífen. (drill 001, aplicado no 002)
