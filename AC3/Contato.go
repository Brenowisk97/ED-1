
package main
import (
	"bufio"
	"fmt"
	"os"
)

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}
func main() {
	var contatos [5]Contato
	var nome, email, opcao string
	var indice int

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover, (3) para exibir contatos, (4) para editar email ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			novoContato := Contato{Nome: nome, Email: email}
			contatos = adicionaContato(Contato{Nome: nome, Email: email}, contatos)
		case "2":
			contatos = excluiContato(contatos)
		case "3":
			fmt.Println("Contatos:")
			for i, contato := range contatos {
				if contato != (Contato{}) {
					fmt.Printf("[%d] Nome: %s | Email: %s\n", i, contato.Nome, contato.Email)
				}
			}
		case "4":
			fmt.Print("Informe o índice do contato que deseja editar: ")
			fmt.Scanln(&indice)

			if indice >= 0 && indice < len(contatos) && contatos[indice] != (Contato{}) {
				fmt.Print("Informe o novo email: ")
				fmt.Scanln(&email)
			}

				editaEmail(&contatos, indice, email)
			}else {
				fmt.Println("Índice inválido ou contato não encontrado.")
			}

			return
		}
}


func adicionaContato(c Contato, a [5]Contato) [5]Contato {
	for ind, contato := range a {
		if (contato == Contato{}) {
			a[ind] = c
			break
		}
	}

	return a
}

func excluiContato(a [5]Contato) [5]Contato {
	if (a[0] == Contato{}) {
		fmt.Println("Lista de contatos vazia!")
		return a
	}

	for ind, contato := range a {
		if (contato == Contato{}) {
			a[ind-1] = Contato{}
		}
	}

	return a
}
func EditaEmail(ind int, novoEmail string, a *[5]Contato) {
	if ind >= 0 && ind < len(a) {
		a[ind].AlterarEmail(novoEmail)
	}
}
func MostraContatoComIndice(ind int, c *Contato) {
	fmt.Printf("[%d] Nome: %s, Email: %s\n", ind, c.Nome, c.Email)
}
func ExibeContatos(a *[5]contato.Contato) {
	for ind, cont := range a {
		if cont != (contato.Contato{}) {
			contato.MostraContatoComIndice(ind, &cont)
		}
	}
}
