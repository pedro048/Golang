package main

import "fmt"

func main() {
	a := 32
	b := &a
	c := a
	d := b
	fmt.Printf("Before: a=%d, b=%d, c=%d and d=%d.\n", a, *b, c, *d)
	*b = *b + 1
	fmt.Printf("After: a=%d, b=%d, c=%d and d=%d.\n", a, *b, c, *d)
}

/* As variaveis b e d copiaram o endereco da variavel a, ou seja, todas as mudancas
de valor em a serao refletidas em b e d. No caso da variavel c foi copiado apenas o valor
de a no momento.
*/
