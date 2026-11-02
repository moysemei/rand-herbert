package slices2

// Plano:
//
// Evens:
// Exemplo: [1, 2, 3, 4] -> [2, 4].
// Padrão: Montar, com filtro.
// Criar um novo slice vazio, percorrer os números
// e verificar quais são pares.
// Adicionar apenas os pares ao resultado.
// Devolver o novo slice sem modificar o original.
//
// FirstN:
// Exemplo: [1, 2, 3, 4], n=2 -> [1, 2].
// Padrão: Montar usando fatiamento e cópia.
// Verificar se n ultrapassa o tamanho do slice.
// Se ultrapassar, ajustar n para len(nums).
// Selecionar os primeiros n elementos com nums[:n].
// Usar slices.Clone para devolver uma cópia independente.
//
// RemoveDuplicates:
// Exemplo: [1, 2, 1, 3, 2] -> [1, 2, 3].
// Padrões: Buscar + Montar.
// Criar um slice vazio para guardar o resultado.
// Percorrer cada número do slice original.
// Verificar com contains se o número já existe
// no resultado.
// Se não existir, adicionar com append.
// Se já existir, não adicionar e continuar.
// No final, devolver o resultado.
//
// Tabela:
// Número | Já está no resultado? | Resultado
//    1   | Não                   | [1]
//    2   | Não                   | [1, 2]
//    1   | Sim                   | [1, 2]
//    3   | Não                   | [1, 2, 3]
//    2   | Sim                   | [1, 2, 3]
//
// contains:
// Padrão: Buscar.
// Percorrer os elementos e comparar com target.
// Se encontrar, retornar true imediatamente.
// Se terminar sem encontrar, retornar false.
import "slices"

func Evens(nums []int) []int {
	result := []int{}
	for _, numb := range nums {
		if numb%2 == 0 {
			result = append(result, numb)
		}
	}

	return result
}

func FirstN(nums []int, n int) []int {
	if n > len(nums) {
		n = len(nums)
	}

	return slices.Clone(nums[:n])
}

func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}

	return false
}

func RemoveDuplicates(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		if contains(result, num) == !true {
			result = append(result, num)
		}
	}

	return result
}

// Dificuldades:
//
// Evens:
// Consegui identificar o padrão Montar e fazer
// o for com range e a condição dos números pares.
// Minha dificuldade foi entender se deveria criar
// um slice novo ou modificar o original.
// Eu até tinha pensado em criar um novo, mas fiquei
// em dúvida porque interpretei errado o enunciado.
// Depois de esclarecer isso, consegui implementar
// o append e retornar o resultado.
//
// FirstN:
// Inicialmente pensei no padrão Buscar e imaginei
// percorrer o slice para pegar os primeiros índices.
// Não percebi que poderia resolver com fatiamento.
// Precisei de ajuda para entender essa abordagem
// e acabei consultando as dicas 2 e 3 do hints.md.
// A dica 3 praticamente mostrava a implementação,
// então não posso dizer que resolvi sozinho.
// Depois consegui entender por que precisamos
// limitar n e utilizar slices.Clone.
// Também pensei em uma alternativa: devolver
// antecipadamente uma cópia de todo o slice
// quando n ultrapassar seu tamanho.
//
// contains:
// Consegui refazer de memória a função que já
// tinha aprendido no Drill 018.
// Não tive dificuldade relevante nessa parte.
//
// RemoveDuplicates:
// Consegui montar sozinho a tabela no Excel,
// identificando quais números deveriam entrar
// no resultado e quais precisavam ser ignorados.
//
// Na hora de transformar isso em Go, travei.
// Primeiro usei contains no slice original,
// quando deveria verificar o slice result.
// Também tentei return false para ignorar
// repetidos, mas a função precisava devolver []int,
// e eu não queria encerrar a função.
//
// Pensei em usar if/else porque queria ignorar
// o número repetido e adicionar os outros.
// Foi quando aprendi sobre continue e entendi
// que poderia colocar o append apenas dentro
// da condição necessária.
//
// Também não conhecia bem o operador !.
// Escrevi contains(result, num) == !true,
// que funciona, mas aprendi que posso escrever
// apenas !contains(result, num).
// Essa última dificuldade foi mais de sintaxe
// do que de entendimento do algoritmo.
//
// Geral:
// Nesse drill, consegui refazer SumNumbers e
// contains de memória, o que mostra que algumas
// coisas dos drills anteriores estão ficando
// mais naturais.
//
// Ainda tenho dificuldade para identificar qual
// ferramenta usar em problemas novos.
// No FirstN, por exemplo, pensei em usar um loop
// quando o fatiamento já resolvia a seleção.
//
// Também percebi que às vezes entendo o resultado
// esperado, mas tenho dificuldade para expressar
// uma decisão em Go, como aconteceu ao tentar
// ignorar números repetidos no RemoveDuplicates.
//
// Preciso continuar praticando como transformar
// o plano em português em operações de Go,
// principalmente sem depender das dicas.
