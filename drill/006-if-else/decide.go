package decide

/* Plan:
   - Primeira func saiu natural.
   - Segunda func, Classify(-4), o que eu acho pra testar primeiro:
     if n == 0
	  return zero
	  if n < 0
	  return negative
	  if n > 0 && n%2 = 0
	  return positive even
	  return positive odd
	- Para a terceira func, não entendi ainda mas tenho que destrinchar o número? 47 = 4*10 + 7*1 e 43 = 4*10 + 3*1
*/

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Classify(n int) string {
	if n == 0 {
		return "zero"
	}
	if n < 0 {
		return "negative"
	}
	if n%2 == 0 {
		return "positive even"
	}
	return "positive odd"
}

func RoundToTen(n int) int {
	lastDigit := n % 10
	if lastDigit >= 5 {
		return n - lastDigit + 10
	}
	return n - lastDigit
}

/*
Dificuldades:
- me batendo em conseguir colocar a condição de positivo par (Att: era pq eu tava colocando n%2 = 0 e não n%2 == 0)
- Para a última func, tive que ir até o terceiro hint e mesmo assim não sei se entendi direito.
*/
