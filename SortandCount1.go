package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter text ")
	address := ""

	fmt.Scanln(address)
	fmt.Println(address)

	f, _ := os.Open(address)
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)

	countWords :=0

	for scanner.Scan() {

		garbage := scanner.Text()
		fmt.Println(garbage)
         countWords++
	}

var input [countWords]string
    i := 0
    for scanner.Scan() {

        input[i] = scanner.Text()

        fmt.Println(input[i])
        i++
        if i == countWords {
            break
        }
    }
    var temp string
    for s := 0; s <= countWords-1; s++ {
        for t := 0; t <= countWords-1; t++ {
            if input[s] < input[t] {

                temp = input[s]
                input[s] = input[t]
                input[t] = temp
            }

        }

    }
    fmt.Println("After sorting the text looks like")
    for d := 0; d <= countWords-1; d++ {
        fmt.Println(input[d])
    }

    for x := 0; x <= countWords-1; x++ {

        word := input[x]
        counter := 0

        for y := 0; y <= countWords-1; y++ {
            if input[y] == word {

                counter++

            }

        }
        fmt.Println(input[x], "is present", counter, "time/times in the text")

    }



}
