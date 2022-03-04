package docker

import (
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
)

var _ DockerCommand = PullCommand{}

type PullCommand struct {
	Name      string        `yaml:"name"`
	Tag       string        `yaml:"tag,omitempty"`
	Digest    string        `yaml:"digest,omitempty"`
	Arguments []string      `yaml:"arguments,omitempty"`
	Flags     builder.Flags `yaml:"flags,omitempty"`
	Outputs   []Output      `yaml:"outputs,omitempty"`
}

func (c PullCommand) GetSuffixArguments() []string {
	return nil
}

func (c PullCommand) GetCommand() string {
	return "docker"
}

func (c PullCommand) GetWorkingDir() string {
	return "."
}

func (c PullCommand) GetArguments() []string {
	// Final Command: docker pull carolynvs/zombies:v1.0 ARGUMENTS --FLAGS

	// Arguments we need to return:
	// pull
	// carolynvs/zombies @ or : abc123 or v1.0
	// ARGUMENTS

	var seperator string
	var tagOrDigest string
	if c.Digest != "" {
		seperator = "@"
		tagOrDigest = c.Digest
	} else {
		seperator = ":"
		tagOrDigest = c.Tag
	}

	args := []string{
		"pull",
		fmt.Sprint(c.Name, seperator, tagOrDigest),
	}
	args = append(args, c.Arguments...)

	return args
}

func (c PullCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c PullCommand) SuppressesOutput() bool {
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
