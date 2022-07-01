package main

import (
	"github.com/flowcommerce/tools/executor"
)

func main() {
	executor := executor.Create("docker-node")

	executor = executor.Add("dev tag")
	executor = executor.Add("./build-node `sem-info tag latest` 12")
	executor = executor.Add("./build-node_builder `sem-info tag latest` 12")
	executor = executor.Add("./build-node `sem-info tag latest` 16")
	executor = executor.Add("./build-node_builder `sem-info tag latest` 16")
	executor = executor.Add("./build-node_selenium_chrome `sem-info tag latest` 16 103.0")

	executor.Run()
}
