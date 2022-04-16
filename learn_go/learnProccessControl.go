package main

import "fmt"

func HelloGolang(language string) {
	if language == "Go" {
		fmt.Println("Hello " + language)
	} else if language == "Python" {
		fmt.Println("Hello ? " + language)
	} else {
		fmt.Println("No")
	}
}
func useSwitch(str string) {
	switch str {
	case "A":
		fmt.Println("Score >=", 90)
	case "B":
		fmt.Println("80<= Score <90")
	case "C":
		fmt.Println("70<= Score <80")
	case "D":
		fmt.Println("60<= Score <70")
	default:
		fmt.Println("GG")
	}
}
func SayGo(number int) {
	for i := 0; i < number; i++ {
		fmt.Println("Hello Golang!")
	}
}

func useForRange(names []string) {
	for index, name := range names {
		fmt.Println(index, name)
	}
}
func main() {
	HelloGolang("Go")
	HelloGolang("Python")
	HelloGolang("Java")

	useSwitch("A")
	useSwitch("B")
	useSwitch("E")

	SayGo(10)

	useForRange([]string{"String", "HKP", "lkr", "abc", "bilibili"})
}
