package main

import (
	"bioflows/virtualization"
	"fmt"
)

func main(){


	docker_manager := virtualization.GetDefaultVirtualizationManager()
	/*pullbuffer, err := docker_manager.PullImage("docker.io/library/alpine")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(pullbuffer.String())
	stdout,stderr , err := docker_manager.StartContainer("","alpine",[]string{"echo","Hello Mohamed Fawzy"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(stdout.String())
	fmt.Println(stderr.String())*/
	containers := docker_manager.ListContainers()
	if len(containers) > 0 {
		for _ , container := range containers{
			docker_manager.StopContainer(container.ID)
		}
	}

	fmt.Println("Done Stopping all Containers")

}

