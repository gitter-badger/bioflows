package main

import (
	"bioflows/helpers/profiling"
	"fmt"
)

func main(){
	fmt.Println(profiling.GetCPUProfile())
}
