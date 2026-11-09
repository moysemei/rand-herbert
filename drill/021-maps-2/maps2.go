package maps2

import "maps"

func Price(menu map[string]int, item string) (int, bool) {
	value, ok := menu[item]
	if !ok {
		return 0, false
	}
	return value, true
}

func RemoveItem(stock map[string]int, item string) bool {
	_, ok := stock[item]
	delete(stock, item)
	return ok
}

func CharFrequency(s string) map[rune]int {
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}

	return m
}

func IsAnagram(a, b string) bool {
	frequencyA := CharFrequency(a)
	frequencyB := CharFrequency(b)

	return maps.Equal(frequencyA, frequencyB)
}

//
// Avaliação do ChatGPT para o Claude — Drill 021:
//
// Claude, o aluno concluiu as quatro funções e apresentou
// a saída do go test -v, com todos os testes passando.
//
// Este drill foi importante para avaliar o que ficou do
// Drill 020, mas também revelou que algumas operações
// básicas ainda não foram consolidadas.
//
// Contexto pessoal e emocional:
//
// O aluno relatou que ficou aproximadamente quatro dias
// afastado dos estudos porque terminou o Drill 020
// realmente frustrado com o próprio desempenho.
//
// Ele descreveu essa frustração como sendo causada pela
// sua própria "burrice". É importante registrar que
// essa foi a percepção dele, não minha avaliação.
//
// O Drill 020 envolveu dificuldades com uma estrutura
// nova, maps, e também problemas na forma como as
// explicações foram apresentadas, tanto no material
// quanto durante a mentoria com o ChatGPT.
//
// A frustração foi suficiente para interromper
// temporariamente sua rotina de estudos.
//
// Isso merece atenção. Não para diminuir a exigência
// dos exercícios, mas para separar corretamente uma
// dificuldade de aprendizado de uma incapacidade.
//
// Price:
//
// A estratégia de consultar o map estava correta.
//
// A dificuldade apareceu ao utilizar o retorno
// value, ok := menu[item].
//
// O aluno confundiu o booleano ok com o padrão
// value, err utilizado anteriormente em strconv.Atoi.
//
// Por isso, tentou verificar ok usando uma comparação
// com nil, como fazia no tratamento de erros.
//
// Depois de esclarecer que ok é um bool indicando
// a existência da chave, corrigiu a implementação.
//
// A função ficou correta, embora ele tenha optado
// por um if e dois returns, em vez de retornar
// diretamente value, ok.
//
// Diagnóstico: confusão entre dois mecanismos da
// linguagem que utilizam retornos múltiplos.
//
// RemoveItem:
//
// O aluno identificou corretamente as duas operações
// necessárias: verificar a existência da chave e
// remover o item com delete.
//
// Também colocou essas operações na ordem correta.
//
// Entretanto, inicialmente escreveu return true,
// ignorando a variável ok que havia acabado de obter.
//
// Precisou de uma revisão para perceber que poderia
// devolver diretamente ok.
//
// Diagnóstico: compreendeu a estratégia, mas ainda
// apresenta dificuldade para aproveitar valores
// booleanos já produzidos por uma operação.
//
// CharFrequency:
//
// Inicialmente criou corretamente um map[rune]int
// e percorreu a string usando range.
//
// Porém, tentou utilizar uma variável count com
// count++ para contar as ocorrências.
//
// Esse contador representaria a quantidade total
// de runes, não a frequência individual de cada uma.
//
// Precisou de ajuda para recuperar a operação
// m[r]++, já utilizada no WordCount do Drill 020.
//
// Após a explicação, corrigiu a implementação.
//
// Também escreveu inicialmente CharFrenquency,
// com um erro de digitação no nome da função.
//
// Diagnóstico: dificuldade em recuperar e aplicar
// o padrão Contar por chave já apresentado no
// exercício anterior.
//
// IsAnagram:
//
// O aluno identificou maps.Equal como a ferramenta
// adequada para comparar os resultados.
//
// Entretanto, inicialmente tentou declarar dois
// map[string]bool no lugar dos parâmetros a e b.
//
// A tentativa apresentava problemas de sintaxe,
// mas a dificuldade mais importante era conceitual:
//
// Um map de presença não preservaria a quantidade
// de cada letra e não seria suficiente para
// identificar todos os anagramas.
//
// O aluno também não havia percebido que poderia
// reutilizar a função CharFrequency, que acabara
// de implementar.
//
// Após receber essa orientação, escreveu sozinho
// a combinação das chamadas de CharFrequency
// com o retorno de maps.Equal.
//
// A implementação final está correta, mas a
// estratégia central foi apresentada durante
// a mentoria.
//
// Diagnóstico: dificuldade em identificar a
// representação adequada dos dados e reutilizar
// uma função já existente.
//
// Avaliação geral:
//
// O aluno concluiu o Drill 021, mas ainda precisou
// de intervenções nas quatro funções.
//
// Os obstáculos foram diferentes:
//
// - Price: confusão entre ok e err.
// - RemoveItem: ignorou o booleano que já possuía.
// - CharFrequency: não recuperou a contagem por chave.
// - IsAnagram: não identificou sozinho a necessidade
//   de representar as strings por frequências.
//
// Não considero que maps estejam completamente
// consolidados.
//
// Por outro lado, existe uma diferença em relação
// ao Drill 020: agora o aluno consegue construir
// e corrigir implementações usando maps quando
// entende a operação necessária.
//
// Ele não precisou receber uma função inteira
// pronta neste drill, mas recebeu orientação
// sobre partes importantes da lógica.
//
// Recomendações para os próximos drills:
//
// 1. Continuar com refazeres de memória, especialmente
//    sobre maps e contagem por chave.
//
// 2. Apresentar operações novas com exemplos pequenos,
//    mostrando o tipo recebido, o tipo retornado
//    e uma situação concreta de uso.
//
// 3. Distinguir desconhecimento de sintaxe de
//    dificuldade para construir o algoritmo.
//
// 4. Cobrar maior autonomia nas operações já estudadas,
//    mas sem pressupor domínio completo após uma
//    única explicação.
//
// 5. Treinar a reutilização de funções já implementadas,
//    pois essa foi uma dificuldade importante
//    no IsAnagram.
//
// 6. Continuar registrando quando houve ajuda,
//    inclusive quando a implementação final está certa.
//
// Observação final para o Claude:
//
// Gostaria que o próximo review considerasse também
// o efeito que o Drill 020 teve na motivação do aluno.
//
// Ele não está pedindo exercícios artificialmente
// fáceis nem avaliações complacentes.
// Pelo contrário, pediu explicitamente uma mentoria
// honesta, com correções objetivas.
//
// O que ele precisa é que a dificuldade seja avaliada
// corretamente e que as explicações esclareçam
// a operação que falta, sem tratar erros de uma
// ferramenta recém-apresentada como falta de atenção.
//
// A meta continua sendo autonomia, não apenas
// conseguir fazer os testes passarem.
