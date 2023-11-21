package main

import (
	"fmt"
)

func main() {
	var nGray, mWhite, foodIntakeDelay, grayLeft, whiteLeft int

	fmt.Print("Write amount of gray mices: ")
	fmt.Scanln(&nGray)

	fmt.Print("Write amount of white mices: ")
	fmt.Scanln(&mWhite)

	fmt.Print("Write the cat's food intake delay: ")
	fmt.Scanln(&foodIntakeDelay)

	fmt.Print("Write amount of gray mice left: ")
	fmt.Scanln(&grayLeft)

	fmt.Print("Write amount of white mice left: ")
	fmt.Scanln(&whiteLeft)

	if nGray < grayLeft || mWhite < whiteLeft {
		fmt.Println("Amount of left mices can't be greater than initial amount")
		return
	}

	mices := make([]int, nGray+mWhite)
	eatenMices := make([]int, 0)

	for i := 0; i < len(mices); i++ {
		if i < nGray {
			mices[i] = i + 1
		} else {
			mices[i] = i + 1
		}
	}
	mices[len(mices)-1] = 0
	fmt.Println(mices)

	indCurr := 0 //номер мыши, с которой начинается счет
	var indPrev int
	for i := 0; i < nGray-grayLeft+mWhite-whiteLeft; i++ {

		fmt.Println()
		for j := 0; j < foodIntakeDelay; j++ { // Отсчитываем S мышей, начиная с indCurr
			indPrev = indCurr        // в indPrev сохраняем номер текущего человека в круге
			indCurr = mices[indCurr] // и вычисляем номер следующего за ним
		}
		fmt.Println(indPrev+1, indCurr+1)
		eatenMices = append(eatenMices, mices[indPrev])
		mices[indPrev] = mices[indCurr]
		fmt.Println(mices)
		indCurr = mices[indCurr] // Новый номер начальной мышки
	}
	fmt.Println("Eaten ", eatenMices)
	fmt.Println("The last ", indCurr+1)

	colored := make([]int, nGray+mWhite)

	for _, eatenInd := range eatenMices {
		colored[eatenInd] = 1
	}

	if colored[0] == 0 && grayLeft < 1 {
		fmt.Println("Arrangement is not possible")
		return
	}

	fmt.Println(colored)
}
