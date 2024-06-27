package main

import (
	"bufio"
	"fmt"
	"github.com/masilvasql/write-in-image/pkg"
	"log"
	"os"
	"sync"
	"time"
)

func main() {

	if len(os.Args) < 5 {
		fmt.Println("Uso: go run main.go <caminho-do-arquivo-com-nomes-a-serem-escritos> <caminho-de-saida> <caminho-do-template> <font-size>")
		return
	}

	filePath := os.Args[1]
	outputPath := os.Args[2]
	templatePath := os.Args[3]
	fontSize := os.Args[4]

	startTime := time.Now()

	arquivo, err := os.Open(filePath)
	defer arquivo.Close()
	if err != nil {
		log.Fatalf("Falha ao abrir arquivo: %v", err)
	}

	scanner := bufio.NewScanner(arquivo)
	var wg sync.WaitGroup
	for scanner.Scan() {
		wg.Add(1)
		input := pkg.NewWriteInImageInput(scanner.Text(), templatePath, outputPath, fontSize, &wg)
		go pkg.WriteInImage(input)
	}
	wg.Wait()

	elapsedTime := time.Since(startTime)
	fmt.Printf("Tempo total de execução: %v\n", elapsedTime)
}
