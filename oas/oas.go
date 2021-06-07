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
	OverrideFiles bool
}

type OpenApiObject struct {
	Openapi string `default:"3.0.3" yaml:"openapi"`
	Info InfoObject `yaml:"info"`
	Paths map[string]Path `yaml:"paths"`
}

type InfoObject struct {
	Title string `yaml:"title"`
	Description string `yaml:"description"`
	Version string `yaml:"version"`
	Contact ContactObject `yaml:"contact"`
	License LicenseObject `yaml:"license"`
}

type ContactObject struct {
	Name string `yaml:"name"`
	Url string `yaml:"url"`
	Email string `yaml:"email"`
}

type LicenseObject struct {
	Name string `yaml:"name"`
	Identifier string `yaml:"identifier"`
	Url string `yaml:"url"`
}

type ServerObject struct {
	Url string `yaml:"url"`
	Description string `yaml:"description"`
}

type Path struct {
	Verbs []HttpVerb`yaml:"verbs"`
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
	// Create components/parameters/uri
	b.CreateFolder("openapi/components/parameters/uri")
	// Create components/parameters/query
	b.CreateFolder("openapi/components/parameters/query")
	// Create components/parameters/header
	b.CreateFolder("openapi/components/parameters/headers")
	// Create components/responses
	b.CreateFolder("openapi/components/responses")
	// Create components/responses
	b.CreateFolder("openapi/components/requests")

	// Let's add README in the folder. It helps keeping the folders in git even if they are empty
	fmt.Println(" ----- Populating with basic README files ---- ")
	// Add root README
	b.WriteFileFromTemplates("README.md", "templates/README.root.md.tmpl", b.Base)
	// Add openapi README
	b.WriteFile("openapi/README.md", "templates/README.openapi.md")
	// Add paths README
	b.WriteFile("openapi/paths/README.md", "templates/README.paths.md")
	// Add components README
	b.WriteFile("openapi/components/README.md", "templates/README.components.md")
	// Add entities README
	b.WriteFile("openapi/components/entities/README.md", "templates/README.entities.md")
	// Add examples README
	b.WriteFile("openapi/components/examples/README.md", "templates/README.examples.md")
	// Add examples requests README
	b.WriteFile("openapi/components/examples/requests/README.md", "templates/README.examples.requests.md")
	// Add examples responses README
	b.WriteFile("openapi/components/examples/responses/README.md", "templates/README.examples.responses.md")
	// Add parameters README
	b.WriteFile("openapi/components/parameters/README.md", "templates/README.parameters.md")
	// Add parameters headers README
	b.WriteFile("openapi/components/parameters/headers/README.md", "templates/README.parameters.headers.md")
	// Add parameters query README
	b.WriteFile("openapi/components/parameters/query/README.md", "templates/README.parameters.query.md")
	// Add parameters uri README
	b.WriteFile("openapi/components/parameters/uri/README.md", "templates/README.parameters.uri.md")
	// Add requests README
	b.WriteFile("openapi/components/requests/README.md", "templates/README.requests.md")
	// Add responses README
	b.WriteFile("openapi/components/responses/README.md", "templates/README.responses.md")

	fmt.Println(" ----- Bootstrapping an amazing API ! ----- ")
	// Write root openapi.yaml file
	b.WriteFileFromTemplates("openapi/openapi.yaml", "templates/openapi.yaml.tmpl", b.Base)
	// Write basic path files

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

	if !b.CanWriteFile(filePath) {
		return
	}

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

	err = ioutil.WriteFile(filePath, t, os.ModePerm)
	if err != nil {
		fmt.Println("error writing '", filePath, "' : ", err)
		os.Exit(-1)
	}

	err = f.Close()
	if err != nil {
		fmt.Println("error closing '", filePath, "' : ", err)
		os.Exit(-1)
	}

	fmt.Println("'", filePath, "' file generated")
}

func (b *Bootstrapper) WriteFileFromTemplates(filePath string, templatePath string, data interface{}) {

	if !b.CanWriteFile(filePath) {
		return
	}

	splitTemplatePath := strings.Split(templatePath, "/")
	templateFilename := splitTemplatePath[len(splitTemplatePath)-1]
	t, err := template.New(templateFilename).Funcs(funcMap).ParseFS(b.Templates, templatePath)
	if err != nil {
		fmt.Println("Internal error generating '", filePath, "': ", err)
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

	err = f.Close()
	if err != nil {
		fmt.Println("error closing '", filePath, "' : ", err)
		os.Exit(-1)
	}

	fmt.Println("'", filePath, "' file generated")
}

func (b *Bootstrapper) CanWriteFile(filePath string) bool {
	info, err := os.Stat(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error checking if '", filePath, "' already exists: ", err)
			os.Exit(-1)
		} else {
			return true
		}
	}

	if info.IsDir() {
		fmt.Println("Error '", filePath, "' exists and is a folder")
		os.Exit(-1)
	}

	if b.OverrideFiles {
		fmt.Println("'", filePath, "' exists and will be override")
		return true
	}

	fmt.Print("'", filePath, "' already exists, override it ? [Y(yes)/n(no)/a(yes-all)]: ")
	var answer string
	_, err = fmt.Scanln(&answer)
	if err != nil {
		fmt.Println("error reading input: ", err)
		os.Exit(-1)
	}

	switch {
	case strings.ToUpper(answer) == "N" || strings.ToUpper(answer) == "NO":
		fmt.Println("'", filePath, "' exists and will not be override")
		return false
	case strings.ToUpper(answer) == "A" || strings.ToUpper(answer) == "YES-ALL":
		fmt.Println("'", filePath, "' exists and will be override")
		b.OverrideFiles = true
		return true
	}
	return true
}