package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/flowcommerce/tools/executor"
	"github.com/flowcommerce/tools/util"
)

func main() {
	executor := executor.Create("docker-postgresql")

	image := fmt.Sprintf("flowdocker/postgresql:%s", latestTag())
	latestImage := "flowdocker/postgresql:latest"
	executor = executor.Add(fmt.Sprintf("docker build -t %s .", image))
	executor = executor.Add(fmt.Sprintf("docker push %s", image))
	executor = executor.Add(fmt.Sprintf("docker tag %s %s", image, latestImage))
	executor = executor.Add(fmt.Sprintf("docker push %s", latestImage))

	image = fmt.Sprintf("flowdocker/postgresql15:%s", latestTag())
	latestImage = "flowdocker/postgresql15:latest"
	executor = executor.Add(fmt.Sprintf("docker build -f Dockerfile-15 -t %s .", image))
	executor = executor.Add(fmt.Sprintf("docker push %s", image))
	executor = executor.Add(fmt.Sprintf("docker tag %s %s", image, latestImage))
	executor = executor.Add(fmt.Sprintf("docker push %s", latestImage))

	executor.Run()
}

func latestTag() string {
	tag, err := exec.Command("sem-info", "tag", "latest").Output()
	util.ExitIfError(err, fmt.Sprintf("Error running sem-info tag latest"))	
	return strings.TrimSpace(string(tag))
}
