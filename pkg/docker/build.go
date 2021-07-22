package docker

import (
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
	"github.com/ghodss/yaml"
)

// This is an example. Replace the following with whatever steps are needed to
// install required components into
const dockerfileLines = `ARG DOCKER_VERSION=%s
RUN apt-get update && apt-get install -y curl && \
	curl -o docker.tgz https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz && \
    tar -xvf docker.tgz && \
    mv docker/docker /usr/bin/docker && \
    chmod +x /usr/bin/docker && \
    rm docker.tgz
`

// BuildInput represents stdin passed to the mixin for the build command.
type BuildInput struct {
	Config MixinConfig
}

// MixinConfig represents configuration that can be set on the docker mixin in porter.yaml
// mixins:
// - docker:
//	  clientVersion: 20.10.7
type MixinConfig struct {
	ClientVersion string `yaml:"clientVersion,omitempty"`
}

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {
	var input BuildInput
	err := builder.LoadAction(m.Context, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &input)
		return &input, err
	})
	if err != nil {
		return err
	}

	if input.Config.ClientVersion != "" {
		m.DockerVersion = input.Config.ClientVersion
	}

	fmt.Fprintf(m.Out, dockerfileLines, m.DockerVersion)
	return nil
}
