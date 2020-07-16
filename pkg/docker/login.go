package docker

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

var _ builder.ExecutableStep = LoginCommand{}

type LoginCommand struct {
	Username      string        `yaml:"username"`
	Password      string        `yaml:"password"`
	Arguments []string      `yaml:"arguments,omitempty"`
	Flags     builder.Flags `yaml:"flags,omitempty"`
	Outputs   []Output      `yaml:"outputs,omitempty"`
}

func (c LoginCommand) GetCommand() string {
	return "docker"
}

func (c LoginCommand) GetArguments() []string {
	// Final Command: docker login -u username -p password ARGUMENTS --FLAGS

	args := []string{
		"login",
		"-u",
		c.Username,
		"-p",
		c.Password,
	}
	args = append(args, c.Arguments...)

	return args
}

func (c LoginCommand) GetFlags() builder.Flags {
	return c.Flags
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
