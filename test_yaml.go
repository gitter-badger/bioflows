package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"bioflows/models"
)

func main(){
	mytool := `
name: "My tool"
Description: "my description"
command: "hello 'from my tool'"
boolbool: On
ignore: ~
# this is the scripts logic
scripts:
    - type: "js"
      order: 1
      before: on
      after : on
      code: >
          This is a multiline comment
          this is also another line for the code
          hahahahaha
`
	tool := &models.Tool{}
	err := yaml.Unmarshal([]byte(mytool),tool)
	if err != nil {
		//fmt.Println("There was a problem unmarshaling the current tool")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(tool.ToJson())
	fmt.Println(fmt.Sprintf(`
	Before : %v,
	After : %v
`,tool.Scripts[0].IsBefore(),tool.Scripts[0].IsAfter()))
}
