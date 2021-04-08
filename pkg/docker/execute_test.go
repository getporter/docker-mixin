package docker

import (
	"bytes"
	"io/ioutil"
	"path"
	"testing"

	"get.porter.sh/porter/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	test.TestMainWithMockedCommandHandlers(m)
}

// TODO: Add test cases for supported actions, we recommend install, update, uninstall and one custom action
func TestMixin_Execute(t *testing.T) {
	testcases := []struct {
		name        string // Test case name
		file        string // Path to th test input yaml
		wantOutput  string // Name of output that you expect to be created
		wantCommand string // Full command that you expect to be called based on the input YAML
	}{
		{"pull", "testdata/pull-input.yaml", "",
			"docker pull getporter/porter-hello:v0.1.0"},
		{"push", "testdata/push-input.yaml", "",
			"docker push getporter/porter-hello:v0.1.0"},
		{"login", "testdata/login-input.yaml", "",
			"docker login -p password -u gmadhok"},
		{"run", "testdata/run-input.yaml", "",
			"docker run -d --env password=password --name practice --privileged --rm getporter/porter-hello"},
		{"remove", "testdata/remove-input.yaml", "",
			"docker rm -f practice"},
		{"build", "testdata/build-input.yaml", "",
			"docker build -f myfile -t practice /Users/myuser/Documents"},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			m := NewTestMixin(t)

			m.Setenv(test.ExpectedCommandEnv, tc.wantCommand)
			mixinInputB, err := ioutil.ReadFile(tc.file)
			require.NoError(t, err)

			m.In = bytes.NewBuffer(mixinInputB)

			err = m.Execute()
			require.NoError(t, err, "execute failed")

			if tc.wantOutput == "" {
				outputs, _ := m.FileSystem.ReadDir("/cnab/app/porter/outputs")
				assert.Empty(t, outputs, "expected no outputs to be created")
			} else {
				wantPath := path.Join("/cnab/app/porter/outputs", tc.wantOutput)
				exists, _ := m.FileSystem.Exists(wantPath)
				assert.True(t, exists, "output file was not created %s", wantPath)
			}
		})
	}
}
