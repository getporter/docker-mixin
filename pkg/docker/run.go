package docker

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

var _ builder.ExecutableStep = RunCommand{}

type RunCommand struct {
	Name      string        `yaml:"name,omitempty"`
	Image     string        `yaml:"image"`
	Detach    bool          `yaml:"detach"`
	Ports     []Ports       `yaml:"ports,omitempty"`
	Env       map[string]string `yaml:"env,omitempty"`
	Privileged     bool      `yaml:"privileged,omitempty"`
	Rm 	    bool `yaml:"rm,omitempty"`
	Command string `yaml:"command,omitempty"`
	Arguments []string      `yaml:"arguments,omitempty"`
	Flags     builder.Flags `yaml:"flags,omitempty"`
	Outputs   []Output      `yaml:"outputs,omitempty"`
}

func (c RunCommand) GetSuffixArguments() []string {

	args := []string {
		c.Image,
	}
	if c.Command != "" {
		args = append(args, c.Command)
	}
	args = append(args, c.Arguments...)
	return args
}

type Ports struct {
	Host string `yaml: "host"`
	Container string `yaml: "container"`
}

//func (p Ports) GetHost() string {
//	return p.Host
//}
//
//func (p Ports) GetContainer() string {
//	return p.Container
//}

func (c RunCommand) GetCommand() string {
	return "docker"
}

func (c RunCommand) GetArguments() []string {
	// Final Command: docker run --privileged --env VAR1=value1 --env VAR2=value2 --rm -p host:container -d --name name image ARGUMENTS --FLAGS

	// Arguments we need to return:
	// push
	// carolynvs/zombies:v1.0
	// ARGUMENTS
	//var priveleged_ = ""
	//if c.Privileged {
	//	priveleged_ = "--privileged"
	//}
	//var envs = ""
	//for key, value := range c.Env {
	//	envs = envs + "--env " + key + "=" + value
	//	envs += " "
	//}
	//envs = strings.TrimSpace(envs)
	//var ports = ""
	//for i:=0; i < len(c.Ports); i++ {
	//	ports = ports + "-p " + c.Ports[i].Host + ":" + c.Ports[i].Container
	//	if i!=len(c.Ports)-1 {
	//		ports = ports + " "
	//	}
	//}
	//var detached = ""
	//if c.Detach {
	//	detached = "-d"
	//}
	//var rm = ""
	//if c.Rm {
	//	rm = "--rm"
	//}
	//var name = ""
	//if c.Name != "" {
	//	name += "--name "
	//	name += c.Name
	//}
	//var options string
	//if priveleged_ != "" {
	//	options += priveleged_
	//	options += " "
	//}
	//if envs != "" {
	//	options += envs
	//	options += " "
	//}
	//if rm != "" {
	//	options += rm
	//	options += " "
	//}
	//if ports != "" {
	//	options += ports
	//	options += " "
	//}
	//if detached != "" {
	//	options += detached
	//	options += " "
	//}
	//if name != "" {
	//	options += name
	//	options += " "
	//}
	//options = strings.TrimSuffix(options," ")

	args := []string{
		"run",
	}
	//if options == "" {
	//	args = append(args, c.Image)
	//} else {
	//	args = append(args, options, c.Image)
	//}
	//
	//

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
	for i:=0; i < len(c.Ports); i++ {
		flags = append(flags, builder.NewFlag("p", c.Ports[i].Host + ":" + c.Ports[i].Container))
	}
	if c.Detach {
		flags = append(flags, builder.NewFlag("d"))
	}
	if c.Rm {
		flags = append(flags, builder.NewFlag("rm"))
	}
	if c.Name != "" {
		flags = append(flags, builder.NewFlag("name", c.Name))
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