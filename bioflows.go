package main

import (
	"bioflows/virtualization"
	"fmt"
)

func main(){


	docker_manager := &virtualization.DockerVirtualizationManager{}
	pullbuffer, err := docker_manager.PullImage("docker.io/library/alpine")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(pullbuffer.String())
	stdout,stderr , err := docker_manager.StartContainer("","alpine",[]string{"sleep","10"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(stdout.String())
	fmt.Println(stderr.String())


}

