package docker

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const buildOutputTemplate = `FROM debian:stretch

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

func TestMixin_Build(t *testing.T) {
	t.Run("build with the default Docker version", func(t *testing.T) {
		m := NewTestMixin(t)
		err := m.Build()
		require.NoError(t, err)

		wantOutput := fmt.Sprintf(buildOutputTemplate, "19.03.8")

		gotOutput := m.TestContext.GetOutput()
		assert.Equal(t, wantOutput, gotOutput)
	})

	t.Run("build with custom Docker version", func(t *testing.T) {
		b, err := ioutil.ReadFile("testdata/build-input-with-version.yaml")
		require.NoError(t, err)

		m := NewTestMixin(t)
		m.In = bytes.NewReader(b)
		err = m.Build()
		require.NoError(t, err)

		wantOutput := fmt.Sprintf(buildOutputTemplate, "19.03.12")

		gotOutput := m.TestContext.GetOutput()
		assert.Equal(t, wantOutput, gotOutput)
	})
}
