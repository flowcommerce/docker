package main

import (
	"github.com/flowcommerce/tools/executor"
)

func main() {
	executor := executor.Create("docker-play")

	executor = executor.Add("dev tag")
	executor = executor.Add("./build-play `sem-info tag latest` 13")
	executor = executor.Add("./build-play-builder `sem-info tag latest` 13")

	executor.Run()
}
