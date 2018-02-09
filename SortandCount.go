package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, _ := os.Open("/Users/tejpatel/Documents/Untitled.txt")
	scanner := bufio.NewScanner(f)
	var input [8]string
	scanner.Split(bufio.ScanWords)
	i := 0
	for scanner.Scan() {

		input[i] = scanner.Text()

		fmt.Println(input[i])
		i++
		if i == 8 {
			break
		}
	}
	var temp string
	for s := 0; s <= 7; s++ {
		for t := 0; t <= 7; t++ {
			if input[s] < input[t] {

				temp = input[s]
				input[s] = input[t]
				input[t] = temp
			}

		}

	}
	fmt.Println("After sorting the text looks like")
	for d := 0; d <= 7; d++ {
		fmt.Println(input[d])
	}

	for x := 0; x <= 7; x++ {

		word := input[x]
		counter := 0

		for y := 0; y <= 7; y++ {
			if input[y] == word {

				counter++

			}

		}
		fmt.Println(input[x], "is present", counter, "time/times in the text")

	}

}
