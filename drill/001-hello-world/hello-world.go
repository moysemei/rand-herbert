// Package hello.
package hello

/* Plano:
Preciso de três funções, a primeira simplesmente para devolver um Hello, World.
A segunda para imprimir outros nomes, a função precisa ser genérica, no sentido de não importa qual nome colocar ali, ou seja precisa ser algo como "Hello," + name.
A terceira, precisa somente passar esse Greetings para uppercase, utilizando strings.ToUpper.
*/

import (
	"strings"
)

func Hello() string {
	return "Hello, World!"
}

func Greet(name string) string {
	return "Hello, " + name + "!"
}

func Shout(name string) string {
	return strings.ToUpper(Greet(name))
}

/* dificuldades:
        -na lógica de criar os nomes em Greet pra dai sim utilizar na func Shout.
		-em definir variáveis, em definir return e funções.
		-criação de funções.
		-erro de retorno com muitos valores, não entendi.
*/
