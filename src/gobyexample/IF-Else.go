package main

import "fmt"

//if-else
//Go中没有三元组，因此if即使在基本条件下也需要使用完整的语句
//请注意，您不需要在Go中的条件周围加上括号，但必须使用花括号。
func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println("is negative")
	} else if num < 9 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has mutiple digis")
		fmt.Print(num, "has mutiple digis\n")
		fmt.Printf("%s", num, "has mutiple digis")
	}

}
