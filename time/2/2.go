package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

// strings.Fields() splits a string around one or more
// consecutive whitespaces (https://golang.org/pkg/unicode/#IsSpace)
// The number of fields is the actual number of words
func CountWords(s string) int {
	return len(strings.Fields(s))
}

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
	start := time.Now()
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

	file, err := ioutil.ReadFile("new_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)

	//str := "Hello, 世🖖🖖界"
	fmt.Println("Count_characters =", utf8.RuneCountInString(text))

	//wc := CountWords("  How  to count words? \n")
	fmt.Println("Count_words =", text)
	// Output:
	// 4
	//var input string
	//input = "Contrary \n to popular \n belief, Lorem Ipsum is not simply \n random text."
	scanner := bufio.NewScanner(strings.NewReader(text))

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanLines)

	// Count the lines.
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Printf("Count_lines = %d \n", count)
	duration := time.Since(start)
	fmt.Println(duration.Milliseconds())
}
