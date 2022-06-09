package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Funcao que escreve um texto no arquivo e retorna um erro caso tenha algum problema
func escreverTexto(linhas []string, caminhoDoArquivo string) error {
	// Cria o arquivo de texto
	arquivo, err := os.Create(caminhoDoArquivo)
	// Caso tenha encontrado algum erro retornar ele
	if err != nil {
		return err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um escritor responsavel por escrever cada linha do slice no arquivo de texto
	escritor := bufio.NewWriter(arquivo)
	for _, linha := range linhas {
		fmt.Fprintln(escritor, linha)
	}

	// Caso a funcao flush retorne um erro ele sera retornado aqui tambem
	return escritor.Flush()
}

func main() {
	var conteudo []string
	conteudo = append(conteudo, "A new line")
	var aux int

	for aux < 150 {
		err := escreverTexto(conteudo, "new_file.txt")
		if err != nil {
			log.Fatalf("Erro:", err)
		}
		aux += aux
	}
	for indice, linha := range conteudo {
		fmt.Println(indice, linha)
	}
}
