package pipelines

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

func ReadPipeline(file io.Reader) (*BioPipeline,error) {

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil , err
	}
	p := &BioPipeline{}
	err = yaml.Unmarshal([]byte(contents),p)
	if err != nil {
		return nil , err
	}
	return p , nil
}
