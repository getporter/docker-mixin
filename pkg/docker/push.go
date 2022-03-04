package docker

import (
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
)

var _ DockerCommand = PushCommand{}

type PushCommand struct {
	Name      string        `yaml:"name"`
	Tag       string        `yaml:"tag,omitempty"`
	Arguments []string      `yaml:"arguments,omitempty"`
	Flags     builder.Flags `yaml:"flags,omitempty"`
	Outputs   []Output      `yaml:"outputs,omitempty"`
}

func (c PushCommand) GetSuffixArguments() []string {
	return nil
}

func (c PushCommand) GetCommand() string {
	return "docker"
}

func (c PushCommand) GetWorkingDir() string {
	return "."
}

func (c PushCommand) GetArguments() []string {
	// Final Command: docker push carolynvs/zombies:v1.0 ARGUMENTS --FLAGS

	// Arguments we need to return:
	// push
	// carolynvs/zombies:v1.0
	// ARGUMENTS

	args := []string{
		"push",
		fmt.Sprint(c.Name, ":", c.Tag),
	}
	args = append(args, c.Arguments...)

	return args
}

func (c PushCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c PushCommand) SuppressesOutput() bool {
	return false
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
