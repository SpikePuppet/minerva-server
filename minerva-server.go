package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Package struct {
	Name    string
	Version string
}

type Tactic struct {
	TacticName string
	Packages   []Package
}

func main() {
	router := gin.Default()

	router.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Minerva Server is running.",
		})
	})

	router.GET("/test-yaml", runGoYamlTest)

	router.Run()
}

func runGoYamlTest(context *gin.Context) {
	testYaml := Tactic{}

	tacticFile, err := ioutil.ReadFile("tactic.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(tacticFile), &testYaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	testYaml.Packages = append(testYaml.Packages, Package{Name: "dotnet-hello-world", Version: "7.2.1"})
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

	context.JSON(200, gin.H{
		"message": string(formattedYaml),
	})

}
