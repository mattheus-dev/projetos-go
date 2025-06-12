package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Estudante struct {
	nome  string
	idade int
}

// Função que verifica se um nome já existe no slice de estudantes
func nomeJaCadastrado(estudantes []Estudante, nome string) bool {
	for _, est := range estudantes {
		if strings.EqualFold(est.nome, nome) { // ignora maiúsculas/minúsculas
			return true
		}
	}
	return false
}

func main() {
	var qtd int
	fmt.Print("Quantos estudantes deseja cadastrar? ")
	fmt.Scanln(&qtd)

	estudantes := make([]Estudante, 0, qtd)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < qtd; i++ {
		var nome string
		for {
			fmt.Printf("Digite o nome do estudante %d: ", i+1)
			scanner.Scan()
			nome = strings.TrimSpace(scanner.Text())
			if nome == "" {
				fmt.Println("O nome não pode ser vazio.")
				continue
			}
			if nomeJaCadastrado(estudantes, nome) {
				fmt.Println("Este nome já foi cadastrado. Digite outro nome!")
			} else {
				break
			}
		}

		var idade int
		for {
			fmt.Printf("Digite a idade de %s: ", nome)
			scanner.Scan()
			idadeStr := strings.TrimSpace(scanner.Text())
			id, err := strconv.Atoi(idadeStr)
			if err != nil || id <= 0 {
				fmt.Println("Idade inválida, digite novamente.")
			} else {
				idade = id
				break
			}
		}

		estudantes = append(estudantes, Estudante{nome: nome, idade: idade})
	}

	fmt.Println("\nLista de estudantes cadastrados:")
	for i, est := range estudantes {
		fmt.Printf("%d) Nome: %s | Idade: %d\n", i+1, est.nome, est.idade)
	}
}