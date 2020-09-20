package main

import (
	"fmt"
	"github.com/hoisie/mustache"
)

func main(){
	var config map[string]interface{} = make(map[string]interface{})
	config["input_dir"] = "/home/snouto"

	data := mustache.Render("ls -ll {{input_dir}}",config)
	fmt.Println(data)

}
