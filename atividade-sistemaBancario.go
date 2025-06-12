package main

import (
	"fmt"
)

type ContaBancaria struct {
	Numero      string
	Titular     string
	Saldo       float64
	AnoAbertura int
}

func (c *ContaBancaria) Depositar(valor float64) {
	c.Saldo = c.Saldo + valor
}

func (c *ContaBancaria) Sacar(valor float64) bool {
	if c.Saldo >= valor {
		c.Saldo = c.Saldo - valor
		return true
	}
	return false
}

func main() {
	var contas []ContaBancaria

	for {
		fmt.Println("1 - Cadastrar nova conta")
		fmt.Println("2 - Ver contas cadastradas")
		fmt.Println("3 - Depositar em uma conta")
		fmt.Println("4 - Sacar de uma conta")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")
		var opcao int
		fmt.Scanln(&opcao)

		if opcao == 1 {
			// Cadastro de conta
			var numero, titular string
			var saldo float64
			var ano int

			fmt.Print("Número da conta: ")
			fmt.Scanln(&numero)
			fmt.Print("Titular: ")
			fmt.Scanln(&titular)

			ok := false
			for !ok {
				fmt.Print("Ano de abertura (>1900): ")
				fmt.Scanln(&ano)
				if ano > 1900 {
					ok = true
				} else {
					fmt.Println("Ano inválido.")
				}
			}

			ok = false
			for !ok {
				fmt.Print("Saldo inicial (>=0): ")
				fmt.Scanln(&saldo)
				if saldo >= 0 {
					ok = true
				} else {
					fmt.Println("Saldo inválido.")
				}
			}

			conta := ContaBancaria{Numero: numero, Titular: titular, Saldo: saldo, AnoAbertura: ano}
			contas = append(contas, conta)
			fmt.Println("Conta criada com sucesso!")

		} else if opcao == 2 {
			// Listar contas
			if len(contas) == 0 {
				fmt.Println("Nenhuma conta cadastrada.")
			} else {
				for i, c := range contas {
					fmt.Printf("%d. Número: %s, Titular: %s, Ano: %d, Saldo: %.2f\n",
						i+1, c.Numero, c.Titular, c.AnoAbertura, c.Saldo)
				}
			}

		} else if opcao == 3 {
			// Depositar
			var numero string
			fmt.Print("Digite o número da conta: ")
			fmt.Scanln(&numero)
			encontrou := false
			for i := range contas {
				if contas[i].Numero == numero {
					var valor float64
					for {
						fmt.Print("Valor do depósito (>0): ")
						fmt.Scanln(&valor)
						if valor > 0 {
							contas[i].Depositar(valor)
							fmt.Println("Depósito realizado!")
							encontrou = true
							break
						} else {
							fmt.Println("Valor inválido.")
						}
					}
					break
				}
			}
			if !encontrou {
				fmt.Println("Conta não encontrada.")
			}

		} else if opcao == 4 {
			// Sacar
			var numero string
			fmt.Print("Digite o número da conta: ")
			fmt.Scanln(&numero)
			encontrou := false
			for i := range contas {
				if contas[i].Numero == numero {
					var valor float64
					for {
						fmt.Print("Valor do saque (>0): ")
						fmt.Scanln(&valor)
						if valor > 0 {
							if contas[i].Sacar(valor) {
								fmt.Println("Saque realizado!")
							} else {
								fmt.Println("Saldo insuficiente.")
							}
							encontrou = true
							break
						} else {
							fmt.Println("Valor inválido.")
						}
					}
					break
				}
			}
			if !encontrou {
				fmt.Println("Conta não encontrada.")
			}
		} else if opcao == 0 {
			fmt.Println("Saindo do sistema. Até logo!")
			break
		} else {
			fmt.Println("Opção inválida!")
		}
	}
}