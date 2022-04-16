package main

import (
	"errors"
	"fmt"
)

//import (
//	"errors"
//	"fmt"
//	"reflect"
//)
//
//type error interface {
//	Error() string
//}
////两种创建error值的方式
//var ErrExampleNew = errors.New("hello world error")
//var ErrExampleFmt = fmt.Errorf("hello world %s", "error")
//
//func main() {
//	fmt.Println(reflect.TypeOf(ErrExampleNew), reflect.TypeOf(ErrExampleFmt))
//}

////error.New的底层实现
//func New(text string) error  {
//	return &errorString(text)
//}
//
//type errorString struct {
//	s string
//}
//
//func (e *errorString) Error() string {
//	return e.s
//}
////fmt.Errorf的底层实现
//func Errorf(format string,a ...interface{}) error{
//	return errors.New(fmt.Sprintf(format,a...))
//}
type ErrorMessage struct {
	Err     error
	Code    int
	Message string
}

func (e *ErrorMessage) Error() string {
	return fmt.Sprintf("e.err = %s,e.code = %d,e.message = %s", e.Err, e.Code, e.Message)
}

var (
	ErrNotRoute = ErrorMessage{
		Err:     errors.New("no route"),
		Code:    404,
		Message: "check route",
	}
	ErrParamNotOk = ErrorMessage{
		Err:     errors.New("param not ok"),
		Code:    404,
		Message: "check param",
	}
)

func main() {
	fmt.Println(ErrNotRoute)
	fmt.Println(ErrParamNotOk)
}
