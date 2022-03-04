package docker

import (
	"os"

	"get.porter.sh/porter/pkg/exec/builder"
)

var _ DockerCommand = LoginCommand{}

type LoginCommand struct {
	Username  string        `yaml:"username"`
	Password  string        `yaml:"password"`
	Arguments []string      `yaml:"arguments,omitempty"`
	Flags     builder.Flags `yaml:"flags,omitempty"`
	Outputs   []Output      `yaml:"outputs,omitempty"`
}

func (c LoginCommand) GetSuffixArguments() []string {
	return nil
}

func (c LoginCommand) GetCommand() string {
	return "docker"
}

func (c LoginCommand) GetWorkingDir() string {
	return "."
}

func (c LoginCommand) GetArguments() []string {
	// Final Command: docker login -u username -p password ARGUMENTS --FLAGS

	args := []string{
		"login",
	}
	args = append(args, c.Arguments...)

	return args
}

func (c LoginCommand) GetFlags() builder.Flags {
	var flags builder.Flags
	var username string = c.Username
	if username == "" {
		username = os.Getenv("DOCKER_USERNAME")
	}
	flags = append(flags, builder.NewFlag("u", username))
	var password string = c.Password
	if password == "" {
		password = os.Getenv("DOCKER_PASSWORD")
	}
	flags = append(flags, builder.NewFlag("p", password))
	return flags
}

func (c LoginCommand) SuppressesOutput() bool {
	return true
}

/*
func (s PullStep) GetOutputs() []builder.Output {
	// Go doesn't have generics, nothing to see here...
	outputs := make([]builder.Output, len(s.Outputs))
	for i := range s.Outputs {
		outputs[i] = s.Outputs[i]
	}
	return outputs
}
*/
