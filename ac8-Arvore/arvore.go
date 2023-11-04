package main

import ("fmt")

type No struct {
	Chave    int
	Esquerda *No
	Direita  *No
}
type ArvoreBinariadeBusca struct {
	Raiz *No
}

func (arvore *ArvoreBinariadeBusca) Inserir(chave int) {
	novoNo := &No{Chave: chave}

	if arvore.Raiz == nil {
		arvore.Raiz = novoNo
	} else {
		inserirNo(arvore.Raiz, novoNo)
	}
}

func inserirNo(Bi, novoNo *No) {
	if novoNo.Chave < Bi.Chave {
		if Bi.Esquerda == nil {
			Bi.Esquerda = novoNo
		} else {
			inserirNo(Bi.Esquerda, novoNo)
		}
	} else if novoNo.Chave > Bi.Chave {
		if Bi.Direita == nil {
			Bi.Direita = novoNo
		} else {
			inserirNo(Bi.Direita, novoNo)
		}
	}
}

func (arvore *ArvoreBinariadeBusca) Buscar(chave int) bool {
	return buscarNo(arvore.Raiz, chave)
}

func buscarNo(Bi *No, chave int) bool {
	if Bi == nil {
		return false
	}

	if chave == Bi.Chave {
		return true
	} else if chave < Bi.Chave {
		return buscarNo(Bi.Esquerda, chave)
	} else {
		return buscarNo(Bi.Direita, chave)
	}
}

func main() {
	arvore := &ArvoreBinariadeBusca{}

	chaves := []int{19, 12, 15, 4, 6, 11, 17}

	for _, chave := range chaves {
		arvore.Inserir(chave)
	}

	fmt.Println("Inserção feita!")

	chaveBusca := 12
	if arvore.Buscar(chaveBusca) {
		fmt.Printf("A chave %d foi encontrada na árvore.\n", chaveBusca)
	} else {
		fmt.Printf("A chave %d não foi encontrada na árvore.\n", chaveBusca)
	}
}