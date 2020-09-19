package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"bioflows/models"
	"io/ioutil"
	"os"
)

func main(){
//	mytool := `
//name: "My tool"
//Description: "my description"
//command: "hello 'from my tool'"
//boolbool: On
//ignore: ~
//# this is the scripts logic
//scripts:
//    - type: "js"
//      order: 1
//      before: on
//      after : on
//      code: >
//          This is a multiline comment
//          this is also another line for the code
//          hahahahaha
//`
	tool := &models.Tool{}
	tool_in, err := os.Open("/home/snouto/projects/bioflows/scripts/tool.bt")
	if err != nil {
		fmt.Printf("There was an error opening the tool file, %v\n",err)
		os.Exit(1)
	}
	mytool_content, err := ioutil.ReadAll(tool_in)
	if err != nil {
		fmt.Printf("Error reading the contents of the tool , %v\n",err)
		os.Exit(1)
	}

	err = yaml.Unmarshal([]byte(mytool_content),tool)
	if err != nil {
		//fmt.Println("There was a problem unmarshaling the current tool")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(tool.ToJson())

//
//	fmt.Println(fmt.Sprintf(`
//	Before : %v,
//	After : %v
//`,tool.Scripts[0].IsBefore(),tool.Scripts[0].IsAfter()))
}
