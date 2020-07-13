package docker

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

func TestMixin_UnmarshalStep(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/step-input.yaml")
	require.NoError(t, err)

	var action Action
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)
	require.Len(t, action.Steps, 1)

	step := action.Steps[0]
	assert.Equal(t, "Prefetch the things", step.Description)

	/*
		assert.NotEmpty(t, step.Outputs)
		assert.Equal(t, Output{Name: "VICTORY", JsonPath: "$Id"}, step.Outputs[0])
	*/

	args := step.GetArguments()
	// docker pull getporter/porter-hello:v0.1.0
	require.Len(t, args, 2)
	assert.Equal(t, "pull", args[0])
	assert.Equal(t, "getporter/porter-hello:v0.1.0", args[1])

	flags := step.GetFlags()
	assert.Empty(t, flags)
}
