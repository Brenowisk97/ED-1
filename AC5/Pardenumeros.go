package main

import ("fmt")

func encontraPar(arr []int, alvo int) (int, int) {
    i, j := 0, len(arr)-1;

    for i < j {
        soma := arr[i] + arr[j]
        if soma == alvo {
            return arr[i], arr[j]
        } else if soma < alvo {
            i++
        } else {
            j--
        }
    }

    return -1, -1;
}
func main() {
    arr := []int{1, 2, 3, 4, 5, 6};
    alvo := 10;

    num1, num2 := encontraPar(arr, alvo)

    if num1 != -1 && num2 != -1 {
        fmt.Printf("Par encontrado: %d + %d = %d\n", num1, num2, alvo)
    } else {
        fmt.Println("Nenhum par encontrado.");
    }
}
