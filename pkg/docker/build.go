package docker

import (
	"fmt"
)

// This is an example. Replace the following with whatever steps are needed to
// install required components into
const dockerfileLines = `FROM debian:stretch

ARG BUNDLE_DIR
RUN apt-get update && apt-get install -y curl ca-certificates

ARG DOCKER_VERSION=19.03.8
RUN curl -o docker.tgz https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz && \
    tar -xvf docker.tgz && \
    mv docker/docker /usr/bin/docker && \
    chmod +x /usr/bin/docker && \
    rm docker.tgz

COPY . $BUNDLE_DIR
`

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {
	fmt.Fprintf(m.Out, dockerfileLines)
	return nil
}
