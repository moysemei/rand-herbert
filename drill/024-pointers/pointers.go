package pointers

// Account is a bank account; Balance is in reais.
type Account struct {
	Owner   string
	Balance int
}

func Double(n *int) {
	*n = *n * 2
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func AddBonus(accounts []Account, bonus int) {
	for i := range accounts {
		accounts[i].Balance += bonus
	}
}

func FindByOwner(accounts []Account, owner string) *Account {
	for i := range accounts {
		if accounts[i].Owner == owner {
			return &accounts[i]
		}
	}

	return nil
}

// Dificuldades:
// Dificuldade em no Double, para multiplicar e devolver o valor
/*
func Double(n *int) {
	p := &n * 2
	return p
}

// ele dá unmatched type, mas &n não continua sendo um int?
*/
// Uma dúvida que parece boba a essa altura: Double não "promete" retornar nada
// então eu não preciso ou preciso obrigatoriamente usar 'return'?
// Tipo, eu entendi, a questão do endereço & ou do valor do ponteiro * mas me
// confundo bastante ainda. Tá distante ainda as coisas, sabe?
