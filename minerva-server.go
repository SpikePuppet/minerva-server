package main

import (
	"fmt"
	"io/ioutil"
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

	tacticFile, err := ioutil.ReadFile("tactic.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(tacticFile), &testYaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	testYaml.Packages = append(testYaml.Packages, Package{Name: "postgres-hello-world", Version: "1.4.1"})
	fmt.Printf("--- T:\n%v\n\n", testYaml)

	formattedYaml, err := yaml.Marshal(&testYaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- testYaml dump:\n%s\n\n", string(formattedYaml))

	err = ioutil.WriteFile("tactic.yaml", []byte(formattedYaml), 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
