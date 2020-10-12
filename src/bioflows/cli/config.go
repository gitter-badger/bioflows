package cli

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func ReadParamsConfig (paramsConfig string) (map[string] interface{},error) {
	params := make(map[string]interface{})
	file , err := os.Open(paramsConfig)
	if err != nil {
		return nil , err
	}
	contents , err := ioutil.ReadAll(file)
	if err != nil {
		return nil , err
	}
	err = yaml.Unmarshal(contents,params)
	if err != nil {
		return nil , err
	}

	return params , nil
}
