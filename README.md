# study-go

Aprendendo Go e backend do zero, 20 minutos por dia, de segunda a sexta.

## Como funciona

- Todo dia útil, perto das 10h, um tutor (Claude) revisa o drill anterior e coloca o próximo em `drill/`.
- Cada drill tem a teoria (`README.md`), testes prontos e dicas em três níveis (`HINTS.md`). A solução eu escrevo do zero.
- Depois que eu entrego, o tutor deixa um `REVIEW.md` com feedback e uma solução de referência.
- A trilha completa está em [`ROADMAP.md`](ROADMAP.md), e onde estou, em [`PROGRESS.md`](PROGRESS.md).

## Rotina

```powershell
cd ~\dev\study-go
git pull                          # traz o drill do dia e o review de ontem
cd drill\NNN-nome                 # a pasta do drill atual (veja em PROGRESS.md)
go test -v                        # vermelho é o ponto de partida
# escrever o // Plano: e o código até ficar verde
cd ..\..
git add .
git commit -m "drill NNN: nome"
git push
```
