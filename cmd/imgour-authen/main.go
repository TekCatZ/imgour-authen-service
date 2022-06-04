package main

import (
	imgourAuthen "github.com/TekCatZ/imgour-authen-service/internal/imgour-authen"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	env := "dev"
	if len(argsWithoutProg) > 1 {
		env = argsWithoutProg[0]
	}
	imgourAuthen.Start(env)
}
