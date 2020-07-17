package docker

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

var _ builder.ExecutableStep = RemoveCommand{}

type RemoveCommand struct {
	Container  string        `yaml:"container"`
	Force      bool       `yaml:"force,omitempty"`
	Arguments []string      `yaml:"arguments,omitempty"`
	Flags     builder.Flags `yaml:"flags,omitempty"`
	Outputs   []Output      `yaml:"outputs,omitempty"`
}

func (c RemoveCommand) GetSuffixArguments() []string {
	return nil
}

func (c RemoveCommand) GetCommand() string {
	return "docker"
}

func (c RemoveCommand) GetArguments() []string {
	// Final Command: docker rm [OPTIONS] CONTAINER ARGUMENTS --FLAGS
	args := []string{
		"rm",
		c.Container,
	}

	args = append(args, c.Arguments...)

	return args
}

func (c RemoveCommand) GetFlags() builder.Flags {

	var flags builder.Flags
	if c.Force {
		flags = append(flags, builder.NewFlag("f"))
	}
	flags = append(flags, c.Flags...)
	return flags
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
