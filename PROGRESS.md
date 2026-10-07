# Progresso

O tutor atualiza este arquivo. O campo `mode` você pode trocar à mão.

```
mode: normal              # normal | maintenance (semana de prova: só revisão de ~10 min)
dose: 20min               # fase de hábito; subir quando eu pedir
phase: 1 - Fundamentos
current_drill: 002-variables-types
status: pending           # pending | done
```

## Histórico

| Drill | Tema | Criado | Entregue | Nota do tutor |
|---|---|---|---|---|
| 001-hello-world | Anatomia de um arquivo Go | 2026-10-06 | 2026-10-07 | Confundiu imprimir com devolver e fixou um valor no lugar do parâmetro; resolveu com perguntas guiadas. `Shout` reaproveitando `Greet` certo de primeira. |
| 002-variables-types | Variáveis, tipos, conversão, constantes | 2026-10-07 | — | — |

## Pontos de atenção

Fraquezas que aparecem nos reviews, para revisitar nas sextas.

- **Imprimir vs devolver:** função devolve com `return`; `fmt.Println` é só para mostrar a um humano. (drill 001)
- **Usar o parâmetro:** testar a função na cabeça com duas entradas diferentes antes de rodar. (drill 001)
- **Plano:** descrever o que a função *devolve*, não o que ela "imprime". (drill 001)
- **Convenções:** comentário em toda função exportada; arquivo `.go` com nome minúsculo e sem hífen. (drill 001)
