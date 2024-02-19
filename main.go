package main

import "fmt"

func main() {
	names := []string{"Vlad", "Nastya", "Stas", "Maksim", "Artem", "Dima"}

	names = append(names, "Vika")
	names = append(names, "Vitaliy")

	fmt.Println(names)

	var (
		parni   []int
		neparni []int
	)

	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			parni = append(parni, i)
		}
		if i%2 != 0 {
			neparni = append(neparni, i)
		}
	}
	fmt.Println(parni)
	fmt.Println(neparni)
}
