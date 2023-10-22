Estrutura Nó:
    Dado: valor
    Anterior: ponteiro na lista duplamente encadeada para o nó anterior
    Próximo: ponteiro na lista duplamente encadeada para o próximo nó

    Estrutura Lista Dupla:
    Cra: ponteiro para o inicio
        Rabo: ponteiro para o final

//exibir nós na lista duplamente encadeada
Funcao ExibirListaDupla(Lista):
    no = Lista.Cra
    Enquanto no não é Nulo:
        Escrever(no.Dado)
        no = no.Próximo

//Busca de um Nó na lista duplamente Encadeada
Funcao BuscarNo(Lista, valorBusca):
    no = Lista.Cra
    Enquanto no não é Nulo:
        Se no.Dado é igual a valorBusca
            Retornar no
        no = no.Próximo
    Retornar Nulo

//Inserção de um Nó na lista Duplamente Encadeada
Procedimento InserirNo(Lista, NovoDado):
    NovoNó = CriarNovoNó(NovoDado)
    Se Lista.Rabo é Nulo:
        Lista.Cra = NovoNó
        Lista.Rabo = NovoNó
    Senão:
        Lista.Rabo.Próximo = NovoNó
        NovoNó.Anterior = Lista.Rabo
        Lista.Rabo = NovoNó

//Remover Nó da lista Duplamente encadeada
Procedimento RemoverNo(Lista, valorRemover):
    NoParaRemover = BuscarNo(Lista, valorRemover)
    Se NoParaRemover não é Nulo:
        Se NoParaRemover.Anterior é Nulo:
            Lista.Cra = NoParaRemover.Próximo
        Senão:
            NoParaRemover.Anterior.Próximo = NoParaRemover.Próximo

        Se NoParaRemover.Próximo é Nulo:
            Lista.Rabo = NoParaRemover.Anterior
        Senão:
            NoParaRemover.Próximo.Anterior = NoParaRemover.Anterior

        Liberar Memória(NoParaRemover)