// Package lists provides exercises for working with slices.
package lists

// Plano:
//
// Sum: percorrer todos os números do slice e ir somando
// em uma variável que começa em 0. Retornar a soma no final.
//
// MaxOf: primeiro verificar se o slice está vazio.
// Se não estiver, começar com o primeiro elemento como maior,
// percorrer os demais e atualizar caso encontre um número maior.
// Retornar o maior e true.
//
// Countdown: criar um slice vazio e utilizar um contador
// que aumenta a cada volta. Adicionar n-i ao slice,
// até chegar ao número 1.
//
// ReverseInPlace: trocar os elementos das extremidades
// do slice, aproximando as posições até chegar ao meio.
// Não criar outro slice e não retornar nada.
//
// Tabela ReverseInPlace([1 2 3 4]):
//
// Volta | i | j | Slice depois da troca
//   1   | 0 | 3 | [4 2 3 1]
//   2   | 1 | 2 | [4 3 2 1]
//
// O loop para quando i chega em len(nums)/2.
// Se continuasse, acabaria desfazendo as trocas.

// Sum returns the sum of all elements in nums.
func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

// MaxOf returns the largest element in nums.
// It returns 0 and false if the slice is empty.
func MaxOf(nums []int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}

	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max, true
}

// Countdown returns a slice containing numbers from n down to 1.
func Countdown(n int) []int {
	result := []int{}

	for i := 0; i < n; i++ {
		result = append(result, n-i)
	}

	return result
}

// ReverseInPlace reverses the elements of nums without creating a new slice.
func ReverseInPlace(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		j := len(nums) - 1 - i

		nums[i], nums[j] = nums[j], nums[i]
	}
}

// Dificuldades:
//
// MaxOf:
// Minha primeira ideia foi começar com max := nums[0] e
// percorrer os números comparando com o maior.
// Até aí beleza, mas eu coloquei return max, true dentro
// do if, em vez de atualizar a variável max.
// Também não tinha pensado no caso de slice vazio.
// O problema é que eu estava retornando antes de terminar
// de verificar os números. Depois entendi que precisava
// atualizar max dentro do loop e retornar só no final.
//
// Countdown:
// Aqui eu me embananei bastante com a lógica.
// Minha primeira ideia foi criar um slice vazio e usar
// for range n, adicionando -1 com append.
// Só que isso simplesmente adicionava vários -1 no slice.
//
// Depois tentei utilizar um for com i aumentando e
// append(result, n-1), pensando que isso iria diminuir
// os números a cada volta.
// O problema é que n não estava mudando, então eu
// continuava adicionando o mesmo número.
//
// Com ajuda, entendi que poderia utilizar n-i,
// já que i aumenta a cada volta e n permanece fixo.
// Também poderia fazer o i começar em n e ir diminuindo.
// Minha dificuldade foi perceber como gerar os valores
// que eu queria adicionar, e não apenas usar o append.
//
// ReverseInPlace:
// Essa foi a função em que mais travei.
// Eu entendi que precisava inverter o slice, mas não
// consegui pensar sozinho em como transformar isso
// em um algoritmo.
//
// Consegui identificar que o primeiro elemento era nums[0]
// e que o último era nums[len(nums)-1].
// Também cheguei perto da lógica para encontrar o índice
// oposto, mas esqueci que os índices começam em zero.
//
// Outra dificuldade foi entender por que o loop precisava
// parar na metade do slice.
// Na minha tentativa, usei i <= len(nums), o que faria
// as trocas serem desfeitas e ainda causaria um panic.
//
// O que mais me pegou foi que eu conseguia responder
// algumas perguntas isoladas, mas não conseguia juntar
// essas respostas e construir a função sozinho.
// Precisei de bastante ajuda com a estrutura.
//
// Geral:
// Minha maior dificuldade nesse drill não foi exatamente
// entender o que cada função deveria fazer.
// O problema foi pegar esse objetivo e descobrir quais
// passos precisava executar para chegar ao resultado.
//
// Consigo entender exemplos e até responder perguntas
// sobre as operações, mas na hora de juntar tudo,
// escolher o for, definir as variáveis e montar a lógica,
// eu acabo travando.
//
// Percebi também que às vezes tento escrever o código
// antes de entender direito o que precisa acontecer.
// Acabo tentando encaixar um for ou um append sem ter
// definido exatamente qual é a função deles no algoritmo.
//
// Outro ponto é que ainda dependo bastante das estruturas
// e das dicas do ChatGPT. Quando recebo a estrutura,
// normalmente consigo acompanhar e completar algumas partes,
// mas construir do zero continua sendo difícil.
//
// Preciso trabalhar melhor a parte de decompor o problema,
// resolver manualmente, organizar os passos e só depois
// transformar isso em código Go.
