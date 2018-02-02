package main

import (
	"github.com/flowcommerce/tools/executor"
)

func main() {
	executor := executor.Create("docker-node")

	executor = executor.Add("dev tag")
	executor = executor.Add("./build-node `sem-info tag latest`")

	executor.Run()
}