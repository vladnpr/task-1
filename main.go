package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	strC := task1()
	task2()
	task3()
	task4()
	task5()
	task6()
	task7()
	task8()
	task9(strC)
}

func task1() string {
	fmt.Print("\n\rTask 1 \n")
	strA := "textA"
	strB := "textB"

	strC := strA + strB

	fmt.Print(strC)

	return strC
}

func task2() {
	fmt.Print("\n\rTask 2 \n\r")
	name := "John Doe"
	age := 30

	fmt.Printf("Customer's name is %s \r Age: %d", name, age)
}

func task3() {
	fmt.Print("\n\rTask 3 \n")
	fmt.Println("2+3=" + strconv.Itoa(2+3))
	fmt.Println("2*2=" + strconv.Itoa(2*2))
	fmt.Println("2/2=" + strconv.Itoa(2/2))
	fmt.Println("2-2=" + strconv.Itoa(2-2))
}

func task4() {
	fmt.Print("\n\rTask 4 \n")
	a := rand.Float32() * 100
	b := rand.Float32() * 100

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	if a > b {
		fmt.Println("a > b")
	}

	if a < b {
		fmt.Println("a < b")
	}

	if a == b {
		fmt.Println("a == b")
	}

	if a <= 10 && a >= 0 {
		fmt.Println("a <= 10 && a >= 0 ")
	}

	if a <= 10 || b <= 10 {
		fmt.Println("a <= 10 && a >= 0 ")
	}
}

func task5() {
	fmt.Print("\n\rTask 5 \n")

	var fruit string = "apple"

	// Використання switch case для визначення значення змінної fruit
	switch fruit {
	case "apple":
		fmt.Println("This is an apple.")
		break
	case "banana":
		fmt.Println("This is a banana.")
		break
	case "orange":
		fmt.Println("This is an orange.")
		break
	default:
		fmt.Println("Unknown fruit.")
	}
}

func task6() {
	fmt.Print("\n\rTask 6 \n")

	number := 7

	// не зрозумів, що таке "конструкція switch-case без switch"
	switch {
	case number > 5, number <= 10:
		fmt.Println("Number is between 5 and 10.")
	case number > 10:
		fmt.Println("Number is greater than 10.")
	default:
		fmt.Println("Number is not within the specified ranges.")
	}
}

func task7() {
	fmt.Print("\n\rTask 7 \n")

	language := "Spanish"

	switch language {
	case "English":
		fmt.Println("Hello!")
	case "French":
		fmt.Println("Bonjour!")
	case "German":
		fmt.Println("Hallo!")
	default:
		fmt.Println("Unknown language. Please specify a valid language.")
	}
}

func task8() {
	fmt.Print("\n\rTask 8 \n")

	var arr = [5]int{1, 2, 3, 4, 5}

	arr[3]++
	arr[4]--

	fmt.Println("arr[3] = " + strconv.Itoa(arr[3]))
	fmt.Println("arr[4] = " + strconv.Itoa(arr[4]))
}

func task9(str string) {
	fmt.Print("\n\rTask 9 \n")

	var slice1 []string

	slice2 := make([]string, 3)

	slice1 = append(slice1, str)

	fmt.Println("Елементи першого слайсу:")
	for _, value := range slice1 {
		fmt.Println(value)
	}

	fmt.Println("Кількість елементів другого слайсу:", len(slice2))
}
