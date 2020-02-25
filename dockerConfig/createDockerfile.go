package dockerConfig

import (
	"io/ioutil"
	"log"
	"os/exec"
)

// function to validate the config
func ValidateConfig(config *DockerConfig) bool {
	// validate operating system image
	cmdStr := "sudo docker search " + config.OperatingSystem
	out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	//fmt.Println("%d", len(out))
	if len(out) <= 90 {
		return false
	}

	// validate programming language image
	cmdStr = "sudo docker search " + config.ProgrammingLanguage
	outP, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	//fmt.Println("%d", len(outP))
	if len(outP) <= 90 {
		return false
	}

	return true
}

// function to write the validated config to dockerfile
func createDockerfile(config DockerConfig) {

	// prepare the data to be written

	err := ioutil.WriteFile("Dockerfile", []byte("Dumping bytes to a file\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
