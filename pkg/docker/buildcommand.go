package docker

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

var _ DockerCommand = BuildCommand{}

type BuildCommand struct {
	Tag       string        `yaml:"tag"`
	File      string        `yaml:"file,omitempty"`
	Path      string        `yaml:"path,omitempty"`
	Arguments []string      `yaml:"arguments,omitempty"`
	Flags     builder.Flags `yaml:"flags,omitempty"`
	Outputs   []Output      `yaml:"outputs,omitempty"`
}

func (c BuildCommand) GetWorkingDir() string {
	return "."
}

func (c BuildCommand) GetSuffixArguments() []string {
	var path = ""
	if c.Path != "" {
		path = c.Path
	} else {
		path = "."
	}
	args := []string{
		path,
	}
	return args
}

func (c BuildCommand) GetCommand() string {
	return "docker"
}

func (c BuildCommand) GetArguments() []string {
	// Final Command: docker build -t tag -f file path ARGUMENTS --FLAGS

	args := []string{
		"build",
	}
	args = append(args, c.Arguments...)

	return args
}

func (c BuildCommand) GetFlags() builder.Flags {
	var flags builder.Flags
	flags = append(flags, builder.NewFlag("t", c.Tag))
	if c.File != "" {
		flags = append(flags, builder.NewFlag("f", c.File))
	}
	flags = append(flags, c.Flags...)
	return flags
}

func (c BuildCommand) SuppressesOutput() bool {
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
