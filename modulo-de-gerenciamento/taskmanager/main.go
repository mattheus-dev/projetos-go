package main

import (
	"fmt"
	"taskmanager/services"
)

func main() {
	service := services.TaskService{}

	for {
		var opcao int
		fmt.Println("\n--- MENU ---")
		fmt.Println("1 - Criar tarefa")
		fmt.Println("2 - Listar tarefas")
		fmt.Println("3 - Concluir tarefa")
		fmt.Println("4 - Remover tarefa")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			fmt.Print("Digite o título da tarefa: ")
			var title string
			fmt.Scanln()
			fmt.Scanln(&title)
			service.CreateTask(title)
			fmt.Println("Tarefa criada!")
		case 2:
			tasks := service.ListTasks()
			if len(tasks) == 0 {
				fmt.Println("Nenhuma tarefa cadastrada.")
			}
			for _, t := range tasks {
				status := "Pendente"
				if t.Completed {
					status = "Concluída"
				}
				fmt.Printf("ID: %s | %s [%s]\n", t.ID, t.Title, status)
			}
		case 3:
			fmt.Print("Digite o ID da tarefa a concluir: ")
			var id string
			fmt.Scan(&id)
			if service.CompleteTask(id) {
				fmt.Println("Tarefa concluída!")
			} else {
				fmt.Println("ID não encontrado.")
			}
		case 4:
			fmt.Print("Digite o ID da tarefa a remover: ")
			var id string
			fmt.Scan(&id)
			if service.RemoveTask(id) {
				fmt.Println("Removida!")
			} else {
				fmt.Println("ID não encontrado.")
			}
		case 0:
			fmt.Println("Até logo!")
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}