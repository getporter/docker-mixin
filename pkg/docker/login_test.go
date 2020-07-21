package docker

import (
	"io/ioutil"
	"os"
	"sort"
	"testing"

	"get.porter.sh/porter/pkg/exec/builder"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestMixin_LoginEnv(t *testing.T) {
	//login test
	os.Setenv("DOCKER_USERNAME", "gmadhok")
	os.Setenv("DOCKER_PASSWORD", "password")
	b, err := ioutil.ReadFile("testdata/login-input.yaml")
	require.NoError(t, err)

	var action Action
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)
	require.Len(t, action.Steps, 1)

	step := action.Steps[0]
	assert.Equal(t, "Login to docker", step.Description)

	args := step.GetArguments()
	assert.Equal(t, "login", args[0])

	flags := step.GetFlags()
	sort.Sort(flags)
	assert.Equal(t, builder.Flags{builder.NewFlag("p", "password"), builder.NewFlag("u", "gmadhok")}, flags)
	os.Unsetenv("DOCKER_USERNAME")
	os.Unsetenv("DOCKER_PASSWORD")
}
