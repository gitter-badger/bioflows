package main

import (
	"fmt"
	"github.com/dop251/goja"
)

func main(){
	vm := goja.New()
	//var loop_var []interface{} = make([]interface{},1)
	//vm.Set("LOOP_VAR",loop_var)
	const mycode string = `
	var LOOP_VAR = [1,2,3,4,5];
	LOOP_VAR.push(1);
`
  _ , err := vm.RunString(mycode)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	val := vm.Get("LOOP_VAR")
	fmt.Println(val.Export())
}
