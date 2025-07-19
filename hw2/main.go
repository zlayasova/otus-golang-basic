package main

import "fmt"

var col int
var str int

func ChessTableRev1() {

	fmt.Println("Вариант программы v1")

	for i := 1; i <= col; i++ {
		j := 1
		if i%2 != 0 {
			for j = 1; j <= str; j++ {
				if j%2 != 0 {
					fmt.Printf("#")
				} else {
					fmt.Printf(" ")
				}
			}
		} else {
			for j = 1; j <= str; j++ {
				if j%2 != 0 {
					fmt.Printf(" ")
				} else {
					fmt.Printf("#")
				}
			}
		}
		fmt.Printf("\n")
	}
}

func ChessTableRev2() {

	fmt.Println("Вариант программы v2")

	for i := 1; i <= str; i++ {
		j := 1
			for j = 1; j <= col; j++ {
				if (j + i)%2 != 0 {
					fmt.Printf(" ")
				}else {
					fmt.Printf("#")
				}
			}
		fmt.Printf("\n")
	}
}

func main() {
	fmt.Println("Введите количество колонок")
	fmt.Scanln(&col)
	fmt.Println("Введите количество строк")
	fmt.Scanln(&str)
	
	ChessTableRev1()
	ChessTableRev2()
}
