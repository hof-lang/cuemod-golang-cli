package {{ .name }}

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

Cli: schema.#Cli & {

	Name:    "{{ .name }}"
	Package: "{{ .repo }}/cmd/{{ .name }}"

	Usage:      "{{ .name }}"
	Short:      "{{ .about }}"
	Long:       Short

	// set to true to print help from root command
	// and always assume subcommands are to be run
	OmitRun: false

	Args: [
		// add any args here
	]

	Commands: [
		// add any subcommands here
	]

	// extras
	VersionCommand: true
	CompletionCommands: true
	{{ if .updates }}Updates: true{{ end }}

	{{ if .telemetry }}
	// set your GA identifier here
	Telemetry: "ua-xxxxxx"
	{{ end }}

	{{ if .releases }}
	// GoReleaser configuration
	// see https://docs.hofstadter.io/... for options
	Releases: {
		Author: ""
		Homepage: ""

		GitHub: {
			Owner: ""
			Repo:  ""
		}

		Docker: {
			Maintainer: ""
			Repo: ""
		}
	}
	{{ end }}
}

