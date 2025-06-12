package main

import (
	"fmt"
	"strings"
)

// Struct do veículo
type Veiculo struct {
	Placa           string
	Modelo          string
	Ano             int
	Abastecimentos  []float64
}

// Struct da frota
type Frota struct {
	Veiculos []Veiculo
}

// Adiciona veículo após checar se a placa já existe
func (f *Frota) AdicionarVeiculo(veiculo Veiculo) {
	for _, v := range f.Veiculos {
		if strings.EqualFold(v.Placa, veiculo.Placa) {
			fmt.Println("Veículo com essa placa já cadastrado!")
			return
		}
	}
	f.Veiculos = append(f.Veiculos, veiculo)
	fmt.Println("Veículo adicionado com sucesso!")
}

// Registra abastecimento se o veículo existir
func (f *Frota) RegistrarAbastecimento(placa string, valor float64) {
	for i, v := range f.Veiculos {
		if strings.EqualFold(v.Placa, placa) {
			f.Veiculos[i].Abastecimentos = append(f.Veiculos[i].Abastecimentos, valor)
			fmt.Println("Abastecimento registrado!")
			return
		}
	}
	fmt.Println("Veículo não encontrado.")
}

// Lista todos os veículos
func (f *Frota) ListarVeiculos() {
	if len(f.Veiculos) == 0 {
		fmt.Println("Nenhum veículo cadastrado.")
		return
	}
	fmt.Println("\n=== Veículos cadastrados ===")
	for _, v := range f.Veiculos {
		fmt.Printf("Placa: %s | Modelo: %s | Ano: %d | Abastecimentos: %v\n",
			v.Placa, v.Modelo, v.Ano, v.Abastecimentos)
	}
}

// Mostra total gasto por veículo
func (f *Frota) ResumoAbastecimentos() {
	if len(f.Veiculos) == 0 {
		fmt.Println("Nenhum veículo cadastrado.")
		return
	}
	fmt.Println("\n=== Resumo de Abastecimentos ===")
	for _, v := range f.Veiculos {
		total := 0.0
		for _, valor := range v.Abastecimentos {
			total += valor
		}
		fmt.Printf("Placa: %s | Total gasto: R$ %.2f\n", v.Placa, total)
	}
}

func main() {
	var frota Frota

	for {
		fmt.Println("\n--- Menu Frota ---")
		fmt.Println("1. Adicionar veículo")
		fmt.Println("2. Registrar abastecimento")
		fmt.Println("3. Listar veículos")
		fmt.Println("4. Mostrar resumo de abastecimentos")
		fmt.Println("5. Sair")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			// Leitura e validação do modelo
			var modelo string
			for {
				fmt.Print("Modelo: ")
				fmt.Scanln(&modelo)
				modelo = strings.TrimSpace(modelo)
				if modelo == "" {
					fmt.Println("Modelo não pode ser vazio.")
					continue
				}
				break
			}
			// Leitura e validação da placa
			var placa string
			for {
				fmt.Print("Placa: ")
				fmt.Scanln(&placa)
				placa = strings.TrimSpace(placa)
				if placa == "" {
					fmt.Println("Placa não pode ser vazia.")
					continue
				}
				break
			}
			// Leitura e validação do ano
			var ano int
			for {
				fmt.Print("Ano: ")
				fmt.Scanln(&ano)
				if ano < 1900 || ano > 2025 {
					fmt.Println("Ano deve ser entre 1900 e 2025.")
					continue
				}
				break
			}
			novo := Veiculo{Placa: placa, Modelo: modelo, Ano: ano}
			frota.AdicionarVeiculo(novo)

		case 2:
			var placa string
			fmt.Print("Placa do veículo: ")
			fmt.Scanln(&placa)
			var valor float64
			for {
				fmt.Print("Valor do abastecimento: ")
				fmt.Scanln(&valor)
				if valor <= 0 {
					fmt.Println("Valor deve ser maior que zero.")
					continue
				}
				break
			}
			frota.RegistrarAbastecimento(placa, valor)

		case 3:
			frota.ListarVeiculos()
		case 4:
			frota.ResumoAbastecimentos()
		case 5:
			fmt.Println("Saindo. Até mais!")
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}