package main

import (
	"errors"
	"fmt"
)

func foo() {
	panic("panic in function foo()")
}
func bar() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Printf("recovered with message %s\n", msg)
		}
	}()
	foo()
	fmt.Println("never get executed")

}

func main() {
	//fmt.Println("entering main()")
	//bar()
	//fmt.Println("exiting main to normal way")
	_, err := httpGet("")
	if err != nil {
		fmt.Println(err)
	}
}

func httpGet(uri string) ([]byte, error) {
	panicIf(uri == "" )
	return []byte("hello world"), nil
}
func panicIf(cond bool, args ...interface{}) {
	if !cond {
		return
	}
	if len(args) == 0 {
		panic(errors.New("cond faild"))
	}
	format := args[0].(string)
	args = args[1:]
	err := fmt.Errorf(format, args)
	panic(err)

}
