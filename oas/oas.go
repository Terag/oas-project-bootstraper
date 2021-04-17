package oas

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Bootstrapper struct {
	Base OpenApiObject
	Templates embed.FS
}

type OpenApiObject struct {
	Openapi string `default:"3.0.3"`
	Info InfoObject
	Paths []Path
}

type InfoObject struct {
	Title string
	Description string
	Version string
	Contact ContactObject
	License LicenseObject
}

type ContactObject struct {
	Name string
	Url string
	Email string
}

type LicenseObject struct {
	Name string
	Identifier string
	Url string
}

type ServerObject struct {
	Url string
	Description string
}

type Path struct {
	Name string
	Verbs []HttpVerb
}

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"Replace": func(old string, new string, i int, s string) string {
		return strings.Replace(s, old, new, i)
	},
	"Slice": func(i, j int, s string) string {
		return s[i:j]
	},
	"SliceBeginning": func(i int, s string) string{
		return s[i:]
	},
	"SliceEnd": func(i int, s string) string{
		return s[:i]
	},
}

func (b *Bootstrapper) Bootstrap() {

	// We start by creating the full project tree
	fmt.Println(" ----- Creating oas project tree ----- ")
	// Create the root for for OAS files
	b.CreateFolder("openapi")
	// Create paths
	b.CreateFolder("openapi/paths")
	// Create components
	b.CreateFolder("openapi/components")
	// Create entities
	b.CreateFolder("openapi/components/entities")
	// Create examples
	b.CreateFolder("openapi/components/examples")
	// Create examples/requests
	b.CreateFolder("openapi/components/examples/requests")
	// Create examples/responses
	b.CreateFolder("openapi/components/examples/responses")
	// Create components/parameters
	b.CreateFolder("openapi/components/parameters")
	// Create components/parameters/path
	b.CreateFolder("openapi/components/parameters/path")
	// Create components/parameters/query
	b.CreateFolder("openapi/components/parameters/query")
	// Create components/parameters/header
	b.CreateFolder("openapi/components/parameters/headers")
	// Create components/responses
	b.CreateFolder("openapi/components/responses")
	// Create components/responses
	b.CreateFolder("openapi/components/requests")
	// Write root openapi.yaml file
	b.WriteFileFromTemplates("openapi/openapi.yaml", "templates/openapi.yaml.tmpl", b.Base)

}

func (b* Bootstrapper) CreateFolder(folderName string) {
	err := os.Mkdir(folderName, os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			fmt.Println("Error generating '", folderName, "' folder")
			os.Exit(-1)
		}
		fmt.Println("'", folderName, "' folder already exists, using it")
	} else {
		fmt.Println("'", folderName, "' folder created")
	}
}

func (b *Bootstrapper) WriteFile(filePath string, templatePath string) {

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("create file: ", err)
		os.Exit(-1)
	}

	t, err := b.Templates.ReadFile(templatePath)
	if err != nil {
		fmt.Println("execute: ", err)
		os.Exit(-1)
	}
	ioutil.WriteFile(filePath, t, os.ModePerm)
	f.Close()
	fmt.Println("'", filePath, "' file generated")
}

func (b *Bootstrapper) WriteFileFromTemplates(filePath string, templatePath string, data interface{}) {

	splitTemplatePath := strings.Split(templatePath, "/")
	templateFilename := splitTemplatePath[len(splitTemplatePath)-1]
	t, err := template.New(templateFilename).Funcs(funcMap).ParseFS(b.Templates, templatePath)
	if err != nil {
		fmt.Println("Internal error generating 'openapi.yaml': ", err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("create file: ", err)
		os.Exit(-1)
	}

	err = t.Execute(f, data)
	if err != nil {
		fmt.Println("execute: ", err)
		os.Exit(-1)
	}
	f.Close()
	fmt.Println("'", filePath, "' root file generated")
}