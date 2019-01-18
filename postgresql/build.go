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
	image := fmt.Sprintf("479720515435.dkr.ecr.us-east-1.amazonaws.com/postgresql:%s", latestTag())

	executor = executor.Add("eval $(aws ecr get-login --no-include-email)")
	executor = executor.Add(fmt.Sprintf("docker build -t %s .", image))
	executor = executor.Add(fmt.Sprintf("docker push %s", image))

	executor.Run()
}

func latestTag() string {
	tag, err := exec.Command("sem-info", "tag", "latest").Output()
	util.ExitIfError(err, fmt.Sprintf("Error running sem-info tag latest"))
	return strings.TrimSpace(string(tag))
}
