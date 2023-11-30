package main

import (
	"bufio"
	"fmt"
	"os"
)

func result(target, guess string) (plus, minus int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if guess[i] == target[j] && i == j {
				plus++
			}
			if guess[i] == target[j] && i != j {
				minus++
			}
		}
	}
	return plus, minus
}
func main() {
	const target = "6453"
	var guess string
	var plus, minus int

	for plus != 4 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Bir sayı tahmin edin: ")
		scanner.Scan()
		guess = scanner.Text()

		plus, minus = result(target, guess)

		fmt.Printf("Artı: %d, Eksi: %d\n", plus, minus)
		plus, minus = 0, 0
		if guess == target {
			break
		}
	}
	fmt.Println("Tebrikler! Sayıyı doğru tahmin ettiniz.")
}
