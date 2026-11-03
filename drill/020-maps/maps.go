// Package counting shows some maps functionalities.
package counting

import "strings"

// Plano:
// Para a primeira função: verifica se aquele item já existe no map
// Caso já exista, soma há quantidade. Caso não, cria o novo item no stock e coloca a quantidade iniciada.

// Para a segunda função: Acredito que aplicar um fields em s, e depois percorrer com um for criando o mapa e o valor de cada palavra.

// Para a terceira função: está nas dificuldades meu plano.

// AddStock add a quantity to an key value
func AddStock(stock map[string]int, item string, qty int) {
	stock[item] += qty
}

// WordCount count how many words exists in s.
func WordCount(s string) map[string]int {
	words := strings.Fields(strings.ToLower(s))

	counts := map[string]int{}

	for _, word := range words {
		counts[word]++
	}

	return counts
}

// Unique returns an slice with only one number of each.
func Unique(nums []int) []int {
	result := []int{}
	seen := map[int]bool{}

	for _, num := range nums {
		if !seen[num] {
			result = append(result, num)
			seen[num] = true
		}
	}

	return result
}

// FirstRepeated returns the first repeated rune.
func FirstRepeated(s string) (rune, bool) {
	seen := map[rune]bool{}

	for _, r := range s {
		if seen[r] {
			return r, true
		}

		seen[r] = true
	}

	return 0, false
}

// Dificuldades:
// Operar com maps. Na primeira função, seguindo meu plano, fiz um if e a minha ideia ali (EU SEI QUE Tá ERRADO), mas minha ideia é comparar se aquele item já tem no meu stock, SE Já, então adicone a quantidade a ele. Se não, crie e adicione. O que me complicou é não saber operar, apesar do review.md mostrar as operações, achei que poderia ter sido mais explícito e ter explicado mais como fazer essas operações, achei as explicações desse drill fracas.
// o que eu havia tentando para a primeira func:
/*
func AddStock(stock map[string]int, item string, qty int) {
	if stock[item] == item {
		stock[item] += qty
	}
	stock[item]qty
}

*/

// Segunda função:
// Dificuldade em utilizar o for com maps, tentei algo como:
/*
func WordCount(s string) map[string]int {
	words := strings.ToLower(s)
	words = strings.Fields(s)

	counts := map[string]int{}

	for _, word := range words {
		counts[word]++
	}

	return counts
}
*/

// Terceira função, tentei verificar com o for se aquele numero já existe em nums, se sim, deletaria:
/*
func Unique(nums []int) []int {
	for _, num := range nums {
		if num == nums[num] {
			delete(nums, num)
		}
	}

	return nums
}
*/

// Geral:
// Confundo slice com map, tanto a funcionalidade mas ainda mais a sintaxe, meu cérebro por algum motivo acha parecido, tanto que no começo da função unique eu achei que era um map ali no parametro.
// Esse drill foi irritante literalmente, meu gato atacou meu pc e eu já fiquei nervoso, depois achei que o drill ficou muito simples e podia ser mais robusto na parte básica, pelo menos mais claro. e o GPT me falando que algo está errado, mesmo eu falando que eu sabia que estava errado, anyway.

//
// Avaliação do ChatGPT para o Claude — Drill 020:
//
// Claude, neste drill apareceram dificuldades diferentes
// das observadas nos exercícios anteriores.
//
// ReverseInPlace (redo):
// O aluno tentou reconstruir a função por aproximadamente
// 10 minutos, mas precisou consultar o Drill 017.
// A implementação ficou correta, porém ele não lembrava
// por que o loop termina em len(nums)/2 nem como calcular
// o índice oposto usando len(nums)-1-i.
// Depois da explicação, conseguiu identificar corretamente
// as trocas para um slice de 7 elementos.
// Isso mostra compreensão após a revisão, mas não permite
// considerar o algoritmo consolidado de memória.
//
// AddStock:
// O plano inicial era verificar se a chave existia,
// atualizar a quantidade se existisse ou criar a entrada
// caso contrário.
// Esse raciocínio era válido, embora desnecessário
// para o problema.
// A dificuldade real foi desconhecer como consultar,
// verificar a existência e atualizar entradas de um map.
// Depois de aprender que uma chave ausente em
// map[string]int retorna zero na leitura, compreendeu
// por que uma única operação de soma resolve o exercício.
//
// Houve também uma falha minha como mentor:
// inicialmente foquei nos erros de sintaxe e de tipos,
// embora o aluno já tivesse reconhecido que sua tentativa
// não compilava.
// Deveria ter avaliado primeiro a estratégia proposta,
// explicado por que ela era válida e apresentado a
// simplificação específica dos maps.
// Isso gerou frustração desnecessária.
//
// WordCount:
// Identificou corretamente a necessidade de separar
// palavras, percorrê-las e contar as ocorrências.
// A dificuldade apareceu na combinação de strings.ToLower
// com strings.Fields.
// Inicialmente tentou guardar um []string em uma variável
// declarada como string. Depois, ao percorrer uma string,
// recebeu runes em vez de palavras.
// Após esclarecer os tipos retornados pelas funções,
// corrigiu a implementação.
//
// Unique:
// Inicialmente tentou identificar e remover os números
// repetidos do slice original, utilizando acesso por
// índice e delete.
// Precisou de orientação para entender que o exercício
// deveria construir um slice novo e usar um map como
// registro dos números já encontrados.
//
// A dificuldade seguinte foi transformar "ainda não vi"
// em uma condição booleana e lembrar de marcar o número
// como visto após adicioná-lo.
// Recebeu diretamente a explicação e o trecho central
// com !seen[num] e seen[num] = true.
// Portanto, a solução final está correta, mas não foi
// construída de forma independente.
//
// FirstRepeated:
// Implementou corretamente a função após o Unique.
// Demonstrou entender que a consulta ao map precisa
// acontecer antes de marcar a rune como vista e que
// o retorno deve ocorrer na primeira repetição.
// Entretanto, recebeu previamente uma explicação
// detalhada, incluindo o tipo do map, a estratégia
// de busca e o comportamento esperado do retorno.
// Houve aplicação correta do conceito recém-explicado,
// mas ainda precisamos testar a retenção posteriormente.
//
// Diagnóstico geral:
//
// O aluno consegue compreender e aplicar operações
// depois de receber explicações concretas e exemplos.
// Entretanto, ainda apresenta dificuldade para escolher
// e combinar ferramentas novas sem orientação.
//
// Neste drill, o principal obstáculo foi a familiaridade
// com maps, especialmente a diferença entre chave e valor,
// a consulta de chaves ausentes e o uso de maps para
// registrar elementos já encontrados.
//
// Também houve confusão entre slices e maps, tanto nas
// operações disponíveis quanto na sintaxe.
// Isso merece prática específica antes de aumentar
// significativamente a complexidade dos exercícios.
//
// A experiência do Drill 020 também mostrou que explicar
// uma operação nova de maneira muito resumida pode não
// ser suficiente para alguém que está tentando construir
// algoritmos com autonomia.
//
// Recomendo que o próximo refazer inclua pelo menos
// uma função com maps, sem consulta inicial, para verificar
// o que realmente ficou retido.
//
// Para a mentoria, a orientação é avaliar primeiro
// se a estratégia do aluno faz sentido, depois explicar
// a ferramenta ou sintaxe que falta e, por último,
// apresentar uma alternativa mais simples quando existir.
//
// Não devemos confundir:
// - Um algoritmo incorreto;
// - Uma estratégia válida, mas menos simples;
// - O desconhecimento de uma operação nova;
// - Uma dificuldade de sintaxe.
//
// O objetivo continua sendo desenvolver autonomia,
// sem atribuir ao aluno uma solução que recebeu pronta
// e sem transformar uma dúvida legítima em erro de lógica.
