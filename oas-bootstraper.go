package main

import (
	"embed"
	"flag"
	"fmt"
	"oas-project-bootstraper/oas"
	"os"
	"runtime"
)

var (
	ApiName string
	Info 	string
	Workspace string
)

//go:embed templates
var Templates embed.FS

func init() {
	helpFlag := flag.Bool("h", false, "Print this message and exit.")
	versionFlag := flag.Bool("v", false, "Print version and build information")
	flag.StringVar(&Workspace, "w", ".", "Directory were the project will be bootstrapped (default: ./)")
	flag.StringVar(&ApiName, "n", "", "API name")
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
				Title: "Test API",
				Description: "Une super API à utiliser par toute la famille",
				Version: "v1",
				Contact: oas.ContactObject {
					Name: "Victor Rouquette",
					Email: "victor@rouquette.me",
				},
				License: oas.LicenseObject {
					Name: "MIT",
				},
			},
			Paths: []oas.Path {
				{
					Name: "/tests",
					Verbs: []oas.HttpVerb {
						oas.GET,
						oas.POST,
					},
				},
				{
					Name: "/tests/{test_ref}",
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

	bootstrapper.Bootstrap()
}