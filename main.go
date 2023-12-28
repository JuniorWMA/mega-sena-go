package main

import (
	"fmt"
	"math/rand"
)

func validation(num int, list []int) int {
	numS := num
	for _, v := range list {

		if num != v && num > 0 {

			numS = num
		}
	}

	return numS
}

func main() {

	listNumb := []int{}

	for len(listNumb) < 6 {
		sortNumb := rand.Intn(60) + 1

		numValid := validation(sortNumb, listNumb)

		listNumb = append(listNumb, numValid)
	}

	fmt.Println("Prepare-se para ficar milionário\nSua sorte chegou!\nConfira seus números da sorte 🤞")

	for _, obj := range listNumb {
		fmt.Println("Número: ", obj)
	}

}
