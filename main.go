package main

import "myapp/bindings"
import (
	gy "github.com/graniticio/granitic-yaml/v2"
)

func main() {
	gy.StartGraniticWithYaml(bindings.Components())
}
