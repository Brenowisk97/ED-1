package main

import (
	"fmt"
)
type Node struct {
	Dado     int    // valor
	Anterior *Node  // ponteiro para o nó anterior
	Proximo  *Node  // ponteiro para o próximo nó
}

type ListaDupla struct {
	Cra  *Node // ponteiro para o início
	Rabo *Node // ponteiro para o final
}

func ExibirListaDupla(lista ListaDupla) {
	no := lista.Cra
	for no != nil {
		fmt.Print(no.Dado, " ")
		no = no.Proximo
	}
	fmt.Println()
}

func BuscarNo(lista ListaDupla, valorBusca int) *Node {
	no := lista.Cra
	for no != nil {
		if no.Dado == valorBusca {
			return no
		}
		no = no.Proximo
	}
	return nil
}

func InserirNo(lista *ListaDupla, novoDado int) {
	novoNode := &Node{
		Dado:     novoDado,
		Anterior: nil,
		Proximo:  nil,
	}
	if lista.Rabo == nil {
		lista.Cra = novoNode
		lista.Rabo = novoNode
	} else {
		lista.Rabo.Proximo = novoNode
		novoNode.Anterior = lista.Rabo
		lista.Rabo = novoNode
	}
}

func RemoverNo(lista *ListaDupla, valorRemover int) {
	nodeParaRemover := BuscarNo(*lista, valorRemover)
	if nodeParaRemover != nil {
		if nodeParaRemover.Anterior != nil {
			nodeParaRemover.Anterior.Proximo = nodeParaRemover.Proximo
		} else {
			lista.Cra = nodeParaRemover.Proximo
		}

		if nodeParaRemover.Proximo != nil {
			nodeParaRemover.Proximo.Anterior = nodeParaRemover.Anterior
		} else {
			lista.Rabo = nodeParaRemover.Anterior
		}
	}
}

func main() {
	// Criando uma lista duplamente encadeada vazia
	lista := ListaDupla{Cra: nil, Rabo: nil}

	// Inserindo elementos na lista
	InserirNo(&lista, 1)
	InserirNo(&lista, 2)
	InserirNo(&lista, 3)

	// Exibindo a lista
	fmt.Print("Lista duplamente encadeada: ")
	ExibirListaDupla(lista)

	// Buscando um elemento
	valorBusca := 2
	elementoBusca := BuscarNo(lista, valorBusca)
	if elementoBusca != nil {
		fmt.Printf("Elemento com valor %d encontrado na lista.\n", valorBusca)
	} else {
		fmt.Printf("Elemento com valor %d não encontrado na lista.\n", valorBusca)
	}

	// Removendo um elemento
	valorRemover := 2
	RemoverNo(&lista, valorRemover)

	// Exibindo a lista após a remoção
	fmt.Print("Lista após remoção: ")
	ExibirListaDupla(lista)
}

//achei melhor fazer a dupla encadeada com node em vez de elemento :)