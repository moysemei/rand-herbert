# Roadmap

A ordem das fases é fixa. A velocidade depende da dose (hoje: 20 min/dia, fase de hábito).
Os drills são numerados na ordem em que são criados, contando revisões de sexta e drills de reforço.
Cada tópico pode virar um ou mais drills, conforme o desempenho.

## Fase 1 — Fundamentos (com lógica em tudo)

**Objetivo:** escrever programas pequenos sozinho, do arquivo vazio ao teste verde.

| Tópico | Conceitos | Ideias de problema |
|---|---|---|
| 1 | Anatomia de um arquivo Go: `package`, `import`, `func`, `return`, strings, nomes exportados, `go test` | Hello, Greet, Shout |
| 2 | Variáveis e tipos: `var`, `:=`, zero values, `int`, `float64`, `string`, `bool`, conversão de tipos, constantes | Celsius → Fahrenheit, minutos → horas e minutos |
| 3 | Operadores: aritméticos, divisão inteira, `%`, comparação, lógicos (`&&`, `\|\|`, `!`) | IsEven, LastDigit, SecondsToClock |
| 4 | `if` / `else if` / `else`, `if` com inicialização | Max, IsLeapYear, Classify |
| 5 | `switch` com e sem expressão | DayType, nota → conceito |
| 6 | `for` clássico | SumTo(n), Factorial, CountMultiples |
| 7 | `for` como while, `break`, `continue` | CountDigits, ReverseNumber, SumDigits |
| 8 | Condições combinadas | FizzBuzz, IsPrime |
| 9 | Funções: múltiplos retornos, retornos nomeados, função como valor (intro) | MinMax, Divide com `ok bool` |
| 10 | Strings 1: bytes vs runes, `len`, indexação, `for range` | CountVowels, Reverse (com acentos) |
| 11 | Strings 2: pacotes `strings` e `strconv` | IsPalindrome ignorando espaços e maiúsculas, conversão segura de texto → número |
| 12 | Arrays e slices: criar, `append`, `len`/`cap`, fatiar | Sum, MaxOf, ReverseSlice |
| 13 | Slices 2: filtrar, remover, cópia vs referência | FilterEven, RemoveDuplicates, SecondLargest |
| 14 | Maps: criar, ler com `ok`, `delete`, percorrer | WordCount, CharFrequency |
| 15 | Maps como ferramenta de lógica | TwoSum, IsAnagram, FirstUniqueChar |
| 16 | Structs | Rectangle, Person, slice de structs |
| 17 | Métodos: receiver por valor e por ponteiro | BankAccount com Deposit/Withdraw |
| 18 | Ponteiros: o que são, quando usar, `nil` | Swap, IncrementAll |
| 19 | Mini-projeto de vários dias | gradebook ou inventário em memória (structs + maps + slices) |

## Fase 2 — Go do dia a dia

- Erros: `error`, `errors.New`, `fmt.Errorf` com `%w`, `errors.Is`/`errors.As`. Por que Go trata erro como valor e não tem exceções.
- **Testes table-driven** — a partir daqui ele escreve parte dos testes; depois, todos.
- Interfaces: implementação implícita, `fmt.Stringer`, interfaces pequenas, `any`.
- Pacotes e módulos: vários pacotes no mesmo módulo, `internal/`, visibilidade.
- `defer`, `panic`/`recover` (e por que quase nunca usar `panic`).
- Generics básico: funções genéricas e constraints.
- Biblioteca padrão útil: `slices`, `maps`, `sort`, `time`, `os`, `bufio`, `encoding/json`.
- Ferramentas: `go vet`, `gofmt`, `go doc`, benchmarks simples.
- **Mini-projeto:** CLI de tarefas que lê e grava JSON em arquivo.

## Fase 3 — Estruturas de dados e algoritmos (o lado entrevista)

- Big-O na prática, contando operações nos próprios drills.
- Two pointers, sliding window, prefix sum.
- Pilha e fila com slices (parênteses balanceados).
- Lista ligada (reverter, detectar ciclo).
- Recursão e memoização (fatorial, fibonacci).
- Busca binária.
- Ordenação: insertion e merge sort para entender; `slices.Sort` no dia a dia.
- Padrões com hash map (frequência, agrupamento).
- Árvores (BST, BFS/DFS) e grafos introdutórios.
- **Formato entrevista:** enunciado em inglês, explicar a solução antes de codar, complexidade no final.
- Depois da fase: lista NeetCode 150 resolvida em Go.

## Fase 4 — Backend

- Como a web funciona: cliente/servidor, TCP, HTTP/1.1 (métodos, status, headers).
- Servidor com `net/http`: handlers, `ServeMux` com padrões (`GET /tasks/{id}`), JSON de entrada e saída.
- API REST: CRUD em memória, rotas, status codes corretos, validação.
- Middleware: logging, recovery, autenticação simples.
- `context`: timeouts e cancelamento.
- Banco: SQL, `database/sql`, Postgres (mais pedido em vagas de fora), migrations, transações, SQL injection.
- Configuração por variável de ambiente, logs estruturados com `log/slog`.
- Testes de handler com `httptest`.
- Autenticação com JWT, segurança (OWASP).
- Docker: Dockerfile multi-stage, `docker compose` com Postgres.
- Graceful shutdown.
- Proxy, reverse proxy e load balancer.
- gRPC (muito comum em vagas Go).

## Fase 5 — Concorrência

- Processo vs thread vs goroutine.
- Goroutines e `sync.WaitGroup`.
- Channels (com e sem buffer), `select`, fechar channel.
- `sync.Mutex`, race conditions, `go test -race`.
- Padrões: worker pool, fan-in/fan-out, pipeline, pub/sub.
- `context` para cancelar goroutines.
- Servidor TCP cru: `net.Listen`, loop de accept, uma goroutine por conexão.
- Server-Sent Events e WebSockets em Go.

## Fase 6 — Projeto de portfólio e entrevistas

- `task-manager-api` refeito do jeito certo: API REST em Go, Postgres, testes, Docker, CI no GitHub Actions, deploy e README em inglês.
- System design básico e entrevistas simuladas em inglês.

## Curso complementar: Fundamentals of Backend Engineering (Hussein Nasser)

Começa quando a dose subir. Recomeçar do início da seção 2 (a aula 6 já foi vista).

| Seção do curso | Aulas | Quando entra |
|---|---|---|
| 2. Backend Communication Design Patterns | 6–16 | Fase 4: 6–11 (request/response, push, sync/async, polling, long polling) e 15 (stateful vs stateless) ao desenhar a API. Fase 5: 12 (SSE), 13 (pub/sub), 14 (multiplexing). 16 (sidecar) opcional. |
| 3. Protocols | 17–30 | Início da Fase 4: 17–25 (OSI, IP, UDP, TCP, TLS, HTTP/1.1, HTTPS). Durante a Fase 4: 26 (WebSockets), 27–28 (HTTP/2 e HTTP/3). Fim da Fase 4: 29 (gRPC). 30 (WebRTC) opcional. |
| 4. Many ways to HTTPS | 31–37 | Junto com TLS e HTTPS (aulas 23 e 25). |
| 5. Backend Execution Patterns | 38–50 | Fase 5: processos, threads, como o backend aceita conexões, servidor TCP cru. 48 (idempotência) na API da Fase 4. |
| 6. Proxying and Load Balancing | 51–53 | Fase 4, junto com Docker e deploy. |
| 7. Extras | 54–59 | 56–57 (jornada de uma requisição) no início da Fase 4; 58 (JWT) na autenticação; 59 (`SELECT COUNT`) na parte de banco; 54–55 opcionais. |
| 8. Bonus Content | 60–64 | 60–62 na Fase 5; 63 (OWASP) na autenticação; 64 (graceful shutdown) no fim da Fase 4. |
