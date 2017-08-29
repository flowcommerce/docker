package main

import (
	"github.com/flowcommerce/tools/executor"
)

func main() {
	executor := executor.Create("docker-rails")

	executor = executor.Add("dev tag")
	executor = executor.Add("./build-rails `sem-info tag latest`")

	executor.Run()
}
