package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Project struct {
	Dep []Dependency `xml:"dependencies>dependency"`
}

type Dependency struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
}

func (d Dependency) String() string {
	return fmt.Sprintf("[groupId: %s], [artifactId: %s], [version: %s]", d.GroupId, d.ArtifactId, d.Version)
}

func dependencies(content []byte) []Dependency {
	var project Project
	xml.Unmarshal(content, &project)
	return project.Dep
}

func main() {

	args := os.Args
	if len(args) == 0 {
		log.Fatal("Wrong number of arguments ... ")
	}

	file := os.Args[1]
	xmlFile, err := os.Open(file)

	if err != nil {
		log.Fatal("error opening file ... ", err)
	}

	defer xmlFile.Close()

	fileContent, _ := ioutil.ReadAll(xmlFile)
	deps := dependencies(fileContent)

	for _, x := range deps {
		fmt.Println(x)
	}
}
