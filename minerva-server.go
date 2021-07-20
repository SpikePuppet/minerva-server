package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var testTactic = `
tacticname: NormalForm
packages:
  - name: pooledenergy-hello-world
    version: 2.0.1
`

type Package struct {
	Name    string
	Version string
}

type Tactic struct {
	TacticName string
	Packages   []Package
}

func main() {
	testYaml := Tactic{}

	err := yaml.Unmarshal([]byte(testTactic), &testYaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- T:\n%v\n\n", testYaml)

	formattedYaml, err := yaml.Marshal(&testYaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- testYaml dump:\n%s\n\n", string(formattedYaml))
}
