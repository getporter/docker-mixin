package docker

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

var _ DockerCommand = RunCommand{}

type RunCommand struct {
	Name           string            `yaml:"name,omitempty"`
	Image          string            `yaml:"image"`
	Detach         bool              `yaml:"detach"`
	Ports          []Ports           `yaml:"ports,omitempty"`
	Env            map[string]string `yaml:"env,omitempty"`
	Privileged     bool              `yaml:"privileged,omitempty"`
	Remove         bool              `yaml:"rm,omitempty"`
	Command        string            `yaml:"command,omitempty"`
	Arguments      []string          `yaml:"arguments,omitempty"`
	Flags          builder.Flags     `yaml:"flags,omitempty"`
	SuppressOutput bool              `yaml:"suppress-output,omitempty"`
	Outputs        []Output          `yaml:"outputs,omitempty"`
}

func (c RunCommand) GetSuffixArguments() []string {

	args := []string{
		c.Image,
	}
	if c.Command != "" {
		args = append(args, c.Command)
	}
	args = append(args, c.Arguments...)
	return args
}

type Ports struct {
	Host      string `yaml:"host"`
	Container string `yaml:"container"`
}

func (c RunCommand) GetCommand() string {
	return "docker"
}

func (c RunCommand) GetWorkingDir() string {
	return "."
}

func (c RunCommand) GetArguments() []string {
	// Final Command: docker run --privileged --env VAR1=value1 --env VAR2=value2 --rm -p host:container -d --name name image ARGUMENTS --FLAGS
	args := []string{
		"run",
	}
	return args
}

func (c RunCommand) GetFlags() builder.Flags {
	var flags builder.Flags
	if c.Privileged {
		flags = append(flags, builder.NewFlag("privileged"))
	}
	for key, value := range c.Env {
		flags = append(flags, builder.NewFlag("env", key+"="+value))
	}
	for _, port := range c.Ports {
		flags = append(flags, builder.NewFlag("p", port.Host+":"+port.Container))
	}
	if c.Detach {
		flags = append(flags, builder.NewFlag("d"))
	}
	if c.Remove {
		flags = append(flags, builder.NewFlag("rm"))
	}
	if c.Name != "" {
		flags = append(flags, builder.NewFlag("name", c.Name))
	}
	flags = append(flags, c.Flags...)
	return flags
}

func (c RunCommand) SuppressesOutput() bool {
	return c.SuppressOutput
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
