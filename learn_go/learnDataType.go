package main

import "fmt"

func main() {
	var (
		number        int     = 1
		numberFloat   float32 = 1.23
		stringExample string  = "Hello"
		char                  = 'A'
		isChar        bool    = false

		numberList = [3]int{1, 2, 3}
		stringList = [...]string{"hello", "world", "kali"}

		numberSlice = numberList[1:]
		stringSlice = stringList[1:]
	)
	type Info struct {
		Name       string
		Age        int
		University string
		Habit      map[string]string
	}
	habits := make(map[string]string)
	habits["ONE"] = "Go"
	habits["TWO"] = "Python"
	var info = Info{
		Name:       "huangkepu",
		Age:        21,
		University: "NWPU",
		Habit:      habits,
	}

	var helloWorld func()
	helloWorld = func() {
		fmt.Println("Hello World!")
	}
	var noName interface{}
	noName = 1
	noName = "12"
	noName = 1.23
	helloWorld()
	fmt.Println(number, numberFloat, numberSlice, stringExample, stringList, stringSlice, char, isChar, info, noName)
}
