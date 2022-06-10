package main

import "fmt"

func main() {
	a := 32
	b := &a
	c := a
	d := b
	fmt.Printf("Before: a=%d, b=%d, c=%d and d=%d.\n", a, *b, c, *d)
	fmt.Printf("Before: a=%d, b=%d, c=%d and d=%d.\n", a, b, c, d)
}

/*
As variaveis b e d copiaram o endereco da variavel a, por isso quando
usadas com * apresentam o valor armazenado em a, mas o * elas apresentam
um valor de crash, um vez que, nao foi salvo nada no endereco delas
propriamente.
*/
