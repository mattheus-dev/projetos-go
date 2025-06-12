package main

import "fmt"

type Livro struct {
	ID         int
	Titulo     string
	Autor      string
	Ano        int
	Disponivel bool
}

type Biblioteca struct {
	Livros     []Livro
	ProximoID  int
}

func (b *Biblioteca) AdicionarLivro(titulo, autor string, ano int) {
	livro := Livro{
		ID:         b.ProximoID,
		Titulo:     titulo,
		Autor:      autor,
		Ano:        ano,
		Disponivel: true,
	}
	b.Livros = append(b.Livros, livro)
	b.ProximoID++
}

func (b *Biblioteca) ListarLivros() {
	fmt.Println("\nID | Titulo                  | Autor          | Ano | Disponivel")
	for _, l := range b.Livros {
		disp := "Sim"
		if !l.Disponivel {
			disp = "Não"
		}
		fmt.Printf("%-3d| %-23s| %-15s| %-4d| %-10s\n", l.ID, l.Titulo, l.Autor, l.Ano, disp)
	}
}

func (b *Biblioteca) EmprestarLivro(id int) {
	for i := range b.Livros {
		if b.Livros[i].ID == id {
			if !b.Livros[i].Disponivel {
				fmt.Println("Livro já emprestado.")
				return
			}
			b.Livros[i].Disponivel = false
			fmt.Println("Livro emprestado!")
			return
		}
	}
	fmt.Println("Livro não encontrado.")
}

func (b *Biblioteca) DevolverLivro(id int) {
	for i := range b.Livros {
		if b.Livros[i].ID == id {
			if b.Livros[i].Disponivel {
				fmt.Println("Livro já está disponível.")
				return
			}
			b.Livros[i].Disponivel = true
			fmt.Println("Livro devolvido!")
			return
		}
	}
	fmt.Println("Livro não encontrado.")
}

func main() {
	b := Biblioteca{
		Livros: []Livro{
			{ID: 1, Titulo: "Aventuras do Golang", Autor: "Joana Silva", Ano: 2016, Disponivel: true},
			{ID: 2, Titulo: "Aprendendo Simples", Autor: "Carlos Souza", Ano: 2020, Disponivel: true},
			{ID: 3, Titulo: "O Mistério Go", Autor: "Lúcia Medeiros", Ano: 2018, Disponivel: true},
		},
		ProximoID: 4,
	}

	for {
		fmt.Println("\n1-Adicionar livro\n2-Listar livros\n3-Emprestar livro\n4-Devolver livro\n0-Sair")
		var op int
		fmt.Scanln(&op)
		if op == 0 {
			break
		}
		switch op {
		case 1:
			var titulo, autor string
			var ano int
			fmt.Print("Titulo: ")
			fmt.Scanln(&titulo)
			fmt.Print("Autor: ")
			fmt.Scanln(&autor)
			fmt.Print("Ano: ")
			fmt.Scanln(&ano)
			b.AdicionarLivro(titulo, autor, ano)
			fmt.Println("Livro adicionado!")
		case 2:
			b.ListarLivros()
		case 3:
			var id int
			fmt.Print("ID: ")
			fmt.Scanln(&id)
			b.EmprestarLivro(id)
		case 4:
			var id int
			fmt.Print("ID: ")
			fmt.Scanln(&id)
			b.DevolverLivro(id)
		default:
			fmt.Println("Opção inválida.")
		}
	}
}