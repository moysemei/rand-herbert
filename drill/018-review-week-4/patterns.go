// Package patterns shows how to run through slices, counting, if contains and which one is best.
package patterns

import (
	"strconv"
	"strings"
)

// Plano:
//
// CountChar:
// Exemplo: "banana", 'a' -> 3.
// Padrão: Contar.
// Percorrer as runes da string, comparar com target
// e aumentar o contador quando encontrar a mesma rune.
//
// Contains:
// Exemplo: [4, 8, 15], 8 -> true.
// Padrão: Buscar.
// Percorrer o slice, comparar cada número com target
// e retornar true se encontrar. Caso percorra tudo
// sem encontrar, retornar false.
//
// LongestWord:
// Exemplo: "the quick brown fox" -> "quick".
// Padrão: Melhor até agora.
// Separar a frase em palavras com strings.Fields.
// Guardar a maior palavra encontrada até agora.
// Comparar os tamanhos com len e atualizar apenas
// quando encontrar uma palavra maior.
// Em caso de empate, manter a primeira.
//
// ParseAll:
// Exemplo: ["1", "2", "3"] -> [1, 2, 3], true.
// Padrões: Montar + Buscar.
// Criar um slice vazio, percorrer as strings
// e converter cada uma usando strconv.Atoi.
// Se encontrar um erro, retornar nil, false.
// Caso contrário, adicionar o número ao slice.
// No final, retornar o slice e true.

// CountChar returns how many times target appears in s.
func CountChar(s string, target rune) int {
	count := 0

	for _, r := range s {
		if r == target {
			count++
		}
	}

	return count
}

// Contains reports whether target exists in nums.
func Contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}

	return false
}

// LongestWord returns the first longest word in s.
// It returns an empty string when there are no words.
func LongestWord(s string) string {
	biggestWord := ""
	phraseToWords := strings.Fields(s)

	for _, word := range phraseToWords {
		if len(word) > len(biggestWord) {
			biggestWord = word
		}
	}

	return biggestWord
}

// ParseAll converts a slice of strings into integers.
// It returns nil and false if any conversion fails.
func ParseAll(words []string) ([]int, bool) {
	newSliceWithOnlyIntegers := []int{}

	for _, word := range words {
		number, err := strconv.Atoi(word)
		if err != nil {
			return nil, false
		}

		newSliceWithOnlyIntegers = append(newSliceWithOnlyIntegers, number)
	}

	return newSliceWithOnlyIntegers, true
}

// Dificuldades:
//
// CountChar:
// Nessa função não tive muita dificuldade.
// Consegui identificar o padrão Contar, montar os
// passos em português e depois transformar em Go.
// Foi uma das funções que consegui resolver sozinho.
//
// Contains:
// Também consegui identificar o padrão Buscar.
// No planejamento, fiquei um pouco em dúvida sobre
// precisar usar len(nums) e sobre onde colocar o
// return false.
// Depois entendi que o false só deve vir depois
// do loop, porque preciso verificar todos os números.
// Consegui implementar a função sem receber o código.
//
// LongestWord:
// Aqui minha cabeça deu uma bugada legal kkk.
// Eu entendia que precisava encontrar a maior palavra,
// mas não conseguia organizar os passos nem em português.
//
// Primeiro pensei em Buscar, mas o padrão correto era
// Melhor até agora.
// Também pensei que precisava separar a frase em um
// slice, mas não lembrava do strings.Fields,
// que já tinha utilizado no Drill 016.
//
// Depois de receber ajuda para montar o raciocínio,
// consegui escrever a função, mas fiz a comparação
// diretamente entre as strings.
// Precisei corrigir para comparar len(word) com
// len(biggestWord), porque o objetivo era comparar
// o tamanho das palavras, não sua ordem alfabética.
//
// ParseAll:
// Consegui identificar os dois padrões, Buscar e Montar.
// Sabia que precisava converter as strings em inteiros
// e guardar os resultados em um novo slice.
//
// O que me pegou foi não lembrar exatamente como usar
// strconv.Atoi e como verificar se a conversão funcionou.
// Cheguei a pensar em algo como comparar um número com
// o tipo int, mas isso não fazia sentido.
//
// Com ajuda, lembrei que strconv.Atoi devolve o número
// e um erro, e que err != nil indica uma falha.
//
// Depois disso, consegui montar a implementação,
// incluindo o append e o retorno antecipado.
//
// Geral:
// Nesse drill consegui identificar alguns padrões
// e transformar meus planos em código sem receber
// a estrutura pronta, principalmente no CountChar
// e no Contains.
//
// Ainda tenho bastante dificuldade quando o problema
// exige combinar mais operações, como no LongestWord.
// Às vezes eu entendo o que o exercício pede, mas
// travo tentando descobrir os passos necessários.
//
// No ParseAll, consegui identificar os dois padrões,
// mas precisei de ajuda para lembrar uma ferramenta
// que já tinha utilizado anteriormente.
//
// Acho que estou melhorando na parte de reconhecer
// padrões, mas ainda preciso praticar bastante a
// decomposição dos problemas e a conexão entre
// os conceitos que já aprendi.
