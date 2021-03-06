package main

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
)

func main(){
	r := "/home/snouto/projects/bioflows/scripts/looppip.yaml"
	u , _ := url.Parse(r)
	//dir := filepath.Dir(u.Path)
	baseURL := fmt.Sprintf("%s://%s",u.Scheme,u.Host)
	fmt.Println(baseURL)
	fmt.Println(strings.Replace(r,filepath.Base(r),"",1))
	fmt.Println(filepath.Base(r))
}

