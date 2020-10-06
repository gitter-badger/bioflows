package helpers

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/url"
)

func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func DownloadBioFlowFile(b interface{}, path string) error {
	resp , err := http.Get(path)
	if err != nil {
		return err
	}
	data , err := ioutil.ReadAll(resp.Body)
	return yaml.Unmarshal(data,b)
}
