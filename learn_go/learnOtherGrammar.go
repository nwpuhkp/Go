package main

import "fmt"

//类型转换

//func main() {
//	var price int = 100
//	fmt.Println(strconv.Itoa(price)) //数值类型转字符串类型
//	fmt.Println(reflect.TypeOf(price))
//	fmt.Println(reflect.TypeOf(strconv.Itoa(price)))
//	var isPrice bool = false
//	fmt.Println(strconv.FormatBool(isPrice)) //布尔型转字符串类型
//	fmt.Println(reflect.TypeOf(isPrice))
//	fmt.Println(reflect.TypeOf(strconv.FormatBool(isPrice)))
//
//}

//自定义类型(从基本类型中派生)
//func main() {
//	type huangkepu int
//	type zhihu string
//	type wechat zhihu
//
//	var huangkepuName huangkepu
//	var huangkepuZhihu zhihu
//	var huangkepuWechat wechat
//
//	huangkepuName = 12
//	huangkepuZhihu = "Learn Golang"
//	huangkepuWechat = "Step by Step"
//
//	fmt.Println(huangkepuName, huangkepuZhihu, huangkepuWechat)
//
//}

//函数
//import "fmt"
//
//func PrintHello() {
//	fmt.Println("Hello World")
//}
//
//var PrintName = func() {
//	fmt.Println("Hello Name") //匿名函数
//}
//
//func main() {
//	PrintHello()
//	PrintName()
//}

//函数带参数和返回值
//func SumNumber(NumberOne int, NumberTwo int) int {
//	return NumberOne * NumberTwo
//}
//
//func main() {
//	//var done int = SumNumber(12, 15)
//	//fmt.Println(done)
//	Done := SumNumber(12, 15)
//	fmt.Println(Done)
//}

//带参数和命名返回值
//func NameResult(numberOne, numberTwo int) (result int) { //相同多个变量可以只用声明一次类型
//	result = numberTwo + numberOne
//	return
//}
//func main() {
//	sumResult := NameResult(12, 15)
//	fmt.Println(sumResult)
//}

//多个返回值
//func MultiResult(numberOne, numberTwo int) (sum int, sum2 string) {
//	sum = numberOne + numberTwo
//	sum2 = strconv.Itoa(sum)
//	return
//}
//func main() {
//	sum, stringSum := MultiResult(15, 16)
//	fmt.Println(sum)
//	fmt.Println(reflect.TypeOf(sum))
//	fmt.Println(stringSum)
//	fmt.Println(reflect.TypeOf(stringSum))
//}

//匿名函数
var anonymousFuncTimes = func(numberOne int) int {
	return numberOne * 10
}

func main() {
	fmt.Println(anonymousFuncTimes(10))
}
