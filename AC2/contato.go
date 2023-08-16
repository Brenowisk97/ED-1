//Elabore um struct Contato, que deve conter os campos nome e e-mail. O struct deve conter um método para alterar o e-mail.
package main

import "fmt"

type Contato struct {
	Nome	string
}

	type Contato struct{
	Nome string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}
//Elabore uma função adicionaContato, que recebe um contato e um array com até 5 elementos e inclui o contato no primeiro elemento vazio do array.

func main() {
	var c Contato
	c = Contato{
		Nome: "Victor",
	}

	var Contato [5]Contato
	contatos[0] = c
	for indice, contato := range contatos {
		if (contato == Contatos{}) {
			fmt.Println("achou espaço vazio no índice", indice)
		}
	}

	contatos = adiciona(p, contatos)
	fmt.Println(contatos)
}

func adiciona(p Contato, a [5]Contato) [5]Contato {
	a[1] = c
	return a
}

//Elabore uma função excluiContato, que recebe um array de 5 elementos e elimina o último contato válido.

	func excluiContato(arrayContatos [5]Contato) ([5]Contato, error) {
		for c:=(pessoa == Pessoa{}) {
			if contatos[c-1]:= {}
			break
			return fmt.Println("não existe contatos para excluir.")
	}
}

//Elabore uma interface por linha de comando na função main, que cria um array de 5 elementos e permite a inclusão ou exclusão de contatos.

func main(){
	var c Contatos
	c = Contatos
	-------
	contatos := make([]Contato, 0, 5)
}

	for {
		fmt.Println(" Escolha uma opção:")
		fmt.Println("1 - Adicionar contato")
		fmt.Println("2 - Remover contato")

		var choice int
		fmt.Print("Opção: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addContato(&contatos)
		case 2:
			removeContato(&contatos)

}
}