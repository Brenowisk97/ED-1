//Estrutura do Nó
Estrutura Nó:
    Dado: valor
    Próximo: ponteiro para o próximo nó na lista circular

Estrutura Lista:
    Cra: ponteiro para o nó cabeça

//Exibir nó
Função ExibirListaCircular(Lista):
    Se Lista.Cra é nulo
        Escrever "Lista vazia"
    Senão
        NóAtual <- Lista.Cra.Próximo
        Enquanto NóAtual não é igual a Lista.Cra
            Escrever(NóAtual.Dado)
            NóAtual <- NóAtual.Próximo

//inserção Nó
Procedimento InserirNoInicio(Lista, NovoDado):
    NovoNó <- CriarNovoNó(NovoDado)
    NovoNó.Próximo <- Lista.Cra.Próximo
    Lista.Cra.Próximo <- NovoNó

//Excluir Nó
    Procedimento ExcluirPrimeiroNó(Lista):
    Se Lista.Cra.Próximo é nulo
        Escrever "Lista vazia, nada para excluir"
    Senão
        NóRemovido <- Lista.Cra.Próximo
        Lista.Cra.Próximo <- NóRemovido.Próximo
        Liberar Memória(NóRemovido)