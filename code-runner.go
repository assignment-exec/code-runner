package main

import (
	"fmt"
	"termination_project/dockerConfig"
)

func main() {
	fmt.Println("In here")
	config, _ := dockerConfig.GetConfig("config.yaml")
	if !dockerConfig.ValidateConfig(config) {
		fmt.Println("Error in configuration. Docker images for given config not found !")
		return
	}


	
}
