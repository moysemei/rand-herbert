// Package methods demonstrates methods with value and pointer receivers.
package methods

// Rectangle is a rectangle with integer sides.
type Rectangle struct {
	Width  int
	Height int
}

// Account is a bank account; Balance is in reais.
type Account struct {
	Owner   string
	Balance int
}

// Area returns the area of the rectangle.
func (r Rectangle) Area() int {
	return r.Height * r.Width
}

// Fits reports whether the rectangle fits inside box,
// either normally or rotated 90 degrees.
func (r Rectangle) Fits(box Rectangle) bool {
	if r.Width <= box.Width && r.Height <= box.Height {
		return true
	}

	if r.Height <= box.Width && r.Width <= box.Height {
		return true
	}

	return false
}

// Deposit adds a positive amount to the account balance.
// Zero or negative amounts are ignored.
func (a *Account) Deposit(amount int) {
	if amount > 0 {
		a.Balance += amount
	}
}

// Withdraw removes a positive amount if the balance is sufficient.
// It returns true when the withdrawal succeeds.
func (a *Account) Withdraw(amount int) bool {
	if amount > 0 && amount <= a.Balance {
		a.Balance -= amount
		return true
	}

	return false
}

// Dificuldades:
//
// Avaliação do ChatGPT para o Claude — Drill 023:
//
// Claude, este drill teve uma evolução interessante:
// o aluno começou com dificuldade para compreender
// conceitualmente o que são métodos e receivers,
// mas conseguiu implementar três dos quatro métodos
// sem ajuda na lógica.
//
// Entendimento de métodos e receivers:
//
// Inicialmente, o aluno relatou que não havia entendido
// muito bem a explicação do README sobre métodos.
//
// Sua dúvida não era somente sobre a sintaxe.
// Ele queria entender para que métodos servem,
// por que são importantes e onde seriam utilizados
// em uma aplicação real.
//
// Foi necessário comparar funções comuns com métodos
// e apresentar exemplos de utilização em sistemas.
//
// Também trabalhamos a diferença entre receivers
// por valor e por ponteiro.
//
// O aluno respondeu corretamente que um método com
// receiver por ponteiro consegue alterar o estado
// original, enquanto um receiver por valor trabalha
// com uma cópia da struct.
//
// Ao final do drill, conseguiu explicar sozinho:
//
// "Porque o deposit tem que modificar o valor original.
// Area não precisa, ele só quer calcular o valor,
// não vai modificar nada original."
//
// Isso indica compreensão da distinção básica,
// embora ainda seja necessário verificar a retenção
// desse conhecimento em exercícios futuros.
//
// Area:
//
// Implementou sozinho e corretamente.
//
// Conseguiu acessar os campos da struct pelo receiver
// e calcular a área multiplicando largura e altura.
//
// Não apresentou dificuldade com a implementação,
// mesmo relatando que o conceito geral de método
// ainda estava um pouco confuso.
//
// Fits:
//
// Esta foi a principal dificuldade lógica do drill.
//
// O aluno entendeu corretamente a relação entre
// receiver e parâmetro na chamada small.Fits(big).
//
// Também conseguiu construir sozinho a condição
// que verifica se o retângulo cabe sem rotação.
//
// Entretanto, não conseguiu descobrir como representar
// a rotação de 90 graus por meio de comparações.
//
// Foi necessário explicar que, ao girar o retângulo,
// a altura passa a ocupar a largura e a largura
// passa a ocupar a altura.
//
// Depois dessa explicação, conseguiu completar
// corretamente o segundo if e o retorno final.
//
// É importante registrar que a dificuldade não foi
// de sintaxe de Go, acesso a campos ou operadores
// booleanos.
//
// O bloqueio aconteceu ao tentar transformar
// uma situação geométrica em condições lógicas.
//
// A implementação final está correta, mas a lógica
// da rotação foi construída com ajuda.
//
// Deposit:
//
// Implementou sozinho, sem solicitar ajuda.
//
// Utilizou corretamente o receiver *Account,
// verificou que amount deveria ser positivo
// e atualizou o saldo usando +=.
//
// Também compreendeu que o método não precisava
// devolver nada, pois modificava a conta original.
//
// Withdraw:
//
// Também implementou sozinho, sem receber
// orientação específica para o algoritmo.
//
// Combinou corretamente duas condições:
//
// - O valor solicitado precisa ser positivo.
// - O valor não pode ultrapassar o saldo.
//
// Usou && para exigir as duas condições,
// subtraiu o valor apenas quando permitido
// e posicionou corretamente os retornos
// true e false.
//
// Essa implementação foi especialmente relevante,
// porque demonstrou autonomia na construção
// de condições booleanas e na modificação
// controlada do estado.
//
// Refazer de memória — Unique:
//
// O aluno informou que realizou o refazer
// do Unique, do Drill 020.
//
// Ainda não revisei o código desse refazer
// nem recebi o resultado específico dos testes.
// Portanto, não posso afirmar que foi resolvido
// corretamente e sem consulta.
//
// Avaliação geral:
//
// O Drill 023 apresentou maior autonomia
// na implementação do que os Drills 020 e 021.
//
// No início, o aluno precisou compreender
// uma ferramenta nova: métodos e receivers.
//
// Depois disso, conseguiu implementar Area,
// Deposit e Withdraw sem ajuda na lógica.
//
// A maior dificuldade apareceu no Fits,
// não por desconhecimento da linguagem,
// mas pela necessidade de construir um
// algoritmo a partir de uma descrição geométrica.
//
// Esse resultado reforça uma distinção importante:
//
// O aluno já consegue aplicar estruturas
// e operadores conhecidos em problemas diretos,
// mas ainda pode travar quando precisa descobrir
// como representar uma situação menos familiar.
//
// Recomendações para os próximos drills:
//
// 1. Continuar praticando métodos e receivers,
//    especialmente a diferença entre alterar
//    uma cópia e modificar o valor original.
//
// 2. Incluir problemas que exijam decompor
//    situações em condições lógicas,
//    sem necessariamente introduzir sintaxe nova.
//
// 3. Manter o Plano opcional, conforme o aluno
//    solicitou anteriormente e o checklist
//    deste drill já permite.
//
// 4. Continuar exigindo o registro das dificuldades,
//    incluindo quanto da solução foi construído
//    com ajuda.
//
// 5. Avaliar separadamente compreensão conceitual,
//    implementação e retenção de memória.
//
// Conclusão:
//
// O aluno demonstrou progresso concreto
// na autonomia para implementar métodos.
//
// Ainda não considero o conceito de receivers
// definitivamente consolidado, pois foi
// apresentado neste drill.
//
// A principal habilidade a desenvolver continua
// sendo transformar problemas menos familiares
// em operações e condições lógicas.
//
// A mentoria deve continuar objetiva:
// avaliar primeiro a estratégia, explicar
// a dificuldade real e evitar fornecer
// soluções completas quando não forem necessárias.
