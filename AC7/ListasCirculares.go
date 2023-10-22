package main

import ("fmt")

type Elemento struct {
	Dado    int        // valor
	Proximo *Elemento // ponteiro para o próximo elemento na lista circular
}

type ListaCircular struct {
	Cabeca *Elemento // ponteiro para o elemento cabeça
}

func ExibirListaCircular(lista ListaCircular) {
	if lista.Cabeca == nil {
		fmt.Println("Lista vazia")
		return
	}

	elementoAtual := lista.Cabeca.Proximo
	for elementoAtual != lista.Cabeca {
		fmt.Print(elementoAtual.Dado, " ")
		elementoAtual = elementoAtual.Proximo
	}
	fmt.Println()
}

// inserir Elemento no inicio da lista
func InserirNoInicio(lista *ListaCircular, novoDado int) {
	novoElemento := &Elemento{
		Dado:    novoDado,
		Proximo: nil,
	}

	if lista.Cabeca == nil {
		lista.Cabeca = &Elemento{Proximo: novoElemento}
		lista.Cabeca.Proximo.Proximo = lista.Cabeca
	} else {
		novoElemento.Proximo = lista.Cabeca.Proximo
		lista.Cabeca.Proximo = novoElemento
	}
}

// excluir o primeiro elemento da lista
func ExcluirPrimeiroElemento(lista *ListaCircular) {
	if lista.Cabeca == nil || lista.Cabeca.Proximo == lista.Cabeca {
		fmt.Println("Lista vazia, nada para excluir")
		return
	}

	elementoRemovido := lista.Cabeca.Proximo
	lista.Cabeca.Proximo = elementoRemovido.Proximo
	elementoRemovido.Proximo = nil
}

func main() {
	// lista circular vazia
	lista := ListaCircular{Cabeca: nil}

	// Inserindo elementos na lista
	InserirNoInicio(&lista, 1)
	InserirNoInicio(&lista, 2)
	InserirNoInicio(&lista, 3)

	// mostra a lista
	fmt.Print("Lista circular: ")
	ExibirListaCircular(lista)

	// Excluindo um elemento
	ExcluirPrimeiroElemento(&lista)

	// Exibindo a lista após a exclusão
	fmt.Print("Lista após exclusão: ")
	ExibirListaCircular(lista)
}