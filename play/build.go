package main

import (
	"github.com/flowcommerce/tools/executor"
)

func main() {
	executor := executor.Create("docker-play")

	executor = executor.Add("dev tag")
	executor = executor.Add("./build-play-base `sem-info tag latest`")
	executor = executor.Add("./build-play `sem-info tag latest` `sem-info tag latest`")
	executor = executor.Add("./build-play-crypto `sem-info tag latest` `sem-info tag latest`")
	executor = executor.Add("./build-update-readme `sem-info tag latest`")

	executor.Run()
}
