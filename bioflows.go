package main

import (
	"fmt"
	"bioflows/tools"
)

func main(){
	tool := tools.NewTool()
	tool.Name = "TopHat"
	tool.Description = "Tophat Description"
	tool.Version = "0.0.1"
	tool.Website = "http://www.bioflows.com"
	tool.Command = "tophat --GTF {{gtf}}"
	fmt.Println(tool.ToJson())


}

