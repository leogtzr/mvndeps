package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Project struct {
	Dep []Dependency `xml:"dependencies>dependency"`
}

type Dependency struct {
	GroupID    string `xml:"groupId"`
	ArtifactID string `xml:"artifactId"`
	Version    string `xml:"version"`
}

func (d Dependency) String() string {
	return fmt.Sprintf("[groupId: %s], [artifactId: %s], [version: %s]", d.GroupID, d.ArtifactID, d.Version)
}

func dependencies(content []byte) []Dependency {
	var project Project
	xml.Unmarshal(content, &project)
	return project.Dep
}

func showDependencies(deps []Dependency, fmtType *string) {

	switch strings.ToLower(*fmtType) {
	case "xml":
		for _, dep := range deps {
			if output, err := xml.MarshalIndent(dep, "  ", "    "); err != nil {
				fmt.Printf("error: %v\n", err)
				continue
			} else {
				fmt.Println(string(output))
			}
		}
	case "json":
		for _, dep := range deps {
			if data, err := json.MarshalIndent(dep, "", "    "); err != nil {
				fmt.Printf("error: %v\n", err)
				continue
			} else {
				fmt.Println(string(data))
			}
		}
	case "text":
		for _, dep := range deps {
			fmt.Println(dep)
		}
	default:
		fmt.Fprintln(os.Stderr, "Wrong format\n")
		os.Exit(1)
	}

}

func main() {

	fmtType := flag.String("type", "text", "Format type")
	flag.Parse()

	args := os.Args

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Wrong number of arguments ... ")
	}

	file := flag.Args()
	xmlFile, err := os.Open(file[0])

	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening file ... ", err)
	}

	defer xmlFile.Close()

	if fileContent, err := ioutil.ReadAll(xmlFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		deps := dependencies(fileContent)
		showDependencies(deps, fmtType)
	}

}
