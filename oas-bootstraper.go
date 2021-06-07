package main

import (
	"embed"
	"flag"
	"fmt"
	"gitlab.beyond-undefined.fr/terag/oas-project-bootstraper/oas-project-bootstraper/oas"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"runtime"
)

var (
	Workspace string
	BaseFilePath string
)

//go:embed templates
var Templates embed.FS

func init() {
	helpFlag := flag.Bool("h", false, "Print this message and exit.")
	versionFlag := flag.Bool("v", false, "Print version and build information")
	flag.StringVar(&Workspace, "w", ".", "Directory were the project will be bootstrapped (default: ./)")
	flag.StringVar(&BaseFilePath, "b", "", "Base file that can be used to bootstrap the OAS. use -h for more information")
	flag.String("help", "help", "Print help information on how to use the CLI")
	flag.Parse()

	switch {
	case *helpFlag:
		flag.Usage()
		os.Exit(0)
	case *versionFlag:
		fmt.Println(runtime.Version())
		os.Exit(0)
	}
}

func main() {

	err := os.Chdir(Workspace)
	if err != nil {
		fmt.Println("Selected Working Directory does not exist: " + Workspace)
		return
	}

	bootstrapper := oas.Bootstrapper {
		Base: oas.OpenApiObject {
			Openapi: "3.0.3",
			Info: oas.InfoObject {
				Title: "Sample",
				Description: "Your amazing API !",
				Version: "v1",
				Contact: oas.ContactObject {
					Name: "John Doe",
					Email: "john.doe@example.com",
				},
				License: oas.LicenseObject {
					Name: "MIT",
				},
			},
			Paths: map[string]oas.Path {
				"/tests": {
					Verbs: []oas.HttpVerb {
						oas.GET,
						oas.POST,
					},
				},
				"/tests/{test_ref}": {
					Verbs: []oas.HttpVerb{
						oas.GET,
						oas.PUT,
						oas.DELETE,
					},
				},
			},
		},
		Templates: Templates,
	}

	if BaseFilePath != "" {
		BaseFile, err := ioutil.ReadFile(BaseFilePath)
		if err != nil {
			fmt.Println("Error reading base file for bootstrapping: ", err)
			os.Exit(-1)
		}
		err = yaml.Unmarshal(BaseFile, &bootstrapper.Base)
	}

	bootstrapper.Bootstrap()
}