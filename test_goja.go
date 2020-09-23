package main

import (
	"fmt"
	"github.com/dop251/goja"
	"io/ioutil"
)

type io struct{
	vm *goja.Runtime
}

func (o *io) Print(call goja.FunctionCall) goja.Value{
	fmt.Println(call.Arguments[0].String())
	return goja.Null()
}
func (o *io) Listdir(call goja.FunctionCall) goja.Value {
	dir := call.Arguments[0].String()
	files := make([]string,0)
	found_files , err := ioutil.ReadDir(dir)
	if err != nil {
		return goja.Null()
	}
	for _ , file := range found_files{
		files = append(files,file.Name())
	}

	return o.vm.ToValue(files)
}

func main(){
	vm := goja.New()
	js := `
	var files = io.Listdir('/home/snouto/Downloads');
	for(var i = 0 ; i < files.length;i++){
	io.Print(files[i]);
}
`
	io := &io{
		vm : vm,
	}
	vm.Set("io",io)
	_ , err := vm.RunString(js)
	if err != nil {
		fmt.Printf("Error : %s",err.Error())
	}
}
