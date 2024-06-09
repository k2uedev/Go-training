package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	tasks := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n### Sistema de gerenciamento de tarefas ###")
		fmt.Println("1º Adicione sua tarefa aqui!")
		fmt.Println("2º Visualize ela aqui!")
		fmt.Println("3. Sair")
		fmt.Print("Escolha uma opção: ")

		if !scanner.Scan() {
			fmt.Println("Erro, sua entrada esta invalida..")
			return
		}

		option, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Opção invalida.")
			continue
		}

		switch option {
		case 1:
			fmt.Print("Digite a tarefa: ")
			if !scanner.Scan() {
				fmt.Println("Erro, sua entrada esta invalida...")
				return
			}
			task := scanner.Text()
			tasks = append(tasks, task)
			fmt.Println("Boa, agora sim, tarefa adicionada com sucesso!")
		case 2:
			fmt.Println("\n### Lista de Tarefas ###")
			for i, task := range tasks {
				fmt.Printf("[%d] %s\n", i+1, task)
			}
		case 3:
			fmt.Println("Saindo do sistema de gerenciamento de tarefas...")
			return
		default:
			fmt.Println("Opção invalida, escolha uma opção valida, ou ficará preso nesta tela..")
		}
	}
}