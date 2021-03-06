package gocommand

const Description = "Runs vgo"

var Usage = []string{`jfrog rt go [command options] <vgo arguments> <target repository>`}

const Arguments string = `	target repository
		Target repository in Artifactory. This will Set GOPROXY environment variable to resolve dependencies from this repository.
	vgo commands
		Arguments and options for the vgo command.`
