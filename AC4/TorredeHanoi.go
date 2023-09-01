package main

import ("fmt")

func hanoi(n int, origem, auxiliar, destino string) int {
	if n == 1 {
		fmt.Printf("Mova o disco 1 de %s para %s\n", origem, destino);
		return 1
	}

	movimentos := 0
	movimentos += hanoi(n-1, origem, destino, auxiliar)
	fmt.Printf("Mova o disco %d de %s para %s\n", n, origem, destino)
	movimentos++
	movimentos += hanoi(n-1, auxiliar, origem, destino)

	return movimentos
}

func main() {
	numDiscos := 3
	movimentos := hanoi(numDiscos, "Torre X", "Torre Y", "Torre Z");
	fmt.Println("Quantidade de movimentos necessários:", movimentos);
}
