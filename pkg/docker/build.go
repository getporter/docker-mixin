package docker

import (
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
	"github.com/ghodss/yaml"
)

// This is an example. Replace the following with whatever steps are needed to
// install required components into
const dockerfileLines = `FROM debian:stretch

ARG BUNDLE_DIR
RUN apt-get update && apt-get install -y curl ca-certificates

ARG DOCKER_VERSION=%s
RUN curl -o docker.tgz https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz && \
    tar -xvf docker.tgz && \
    mv docker/docker /usr/bin/docker && \
    chmod +x /usr/bin/docker && \
    rm docker.tgz

COPY . $BUNDLE_DIR
`

// BuildInput represents stdin passed to the mixin for the build command.
type BuildInput struct {
	Config MixinConfig
}

// MixinConfig represents configuration that can be set on the docker mixin in porter.yaml
// mixins:
// - docker:
//	  clientVersion: 19.03.8
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
