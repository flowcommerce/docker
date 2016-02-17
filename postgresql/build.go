package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/flowcommerce/tools/util"
)

func main() {
	image := fmt.Sprintf("flowcommerce/postgresql:%s", latestTag())
	fmt.Printf("Building docker image: %s\n", image)

	runDocker(fmt.Sprintf("docker build -t %s .", image))
	fmt.Printf("Built docker image: %s\n", image)

	runDocker(fmt.Sprintf("docker push %s", image))
	fmt.Printf("Pushed docker image: %s\n", image)
}

func latestTag() string {
	tag, err := exec.Command("sem-info", "tag", "latest").Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(tag))
}

func runDocker(cmdStr string) string {
	fmt.Printf("%s\n", cmdStr)
	return string(util.RunCmd(exec.Command("/bin/sh", "-c", cmdStr), false))
}
