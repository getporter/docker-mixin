package docker

import (
	"get.porter.sh/porter/pkg/exec/builder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func TestMixin_UnmarshalStep(t *testing.T) {

	//build test
	b, err := ioutil.ReadFile("testdata/build-input.yaml")
	require.NoError(t, err)

	var action Action
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)
	require.Len(t, action.Steps, 1)

	step := action.Steps[0]
	assert.Equal(t, "Build image", step.Description)

	args := step.GetArguments()
	// docker build -t tag -f file path ARGUMENTS --FLAGS
	require.Len(t, args, 1)
	assert.Equal(t, "build", args[0])
	wantFlags := builder.Flags {
		builder.Flag {
			Name: "t",
			Values: []string{"practice"},
		},
		builder.Flag {
			Name: "f",
			Values: []string{"myfile"},
		},
	}
	assert.Equal(t, wantFlags, step.GetFlags())

	suffixArgs := step.GetSuffixArguments()
	require.Len(t, suffixArgs, 1)
	assert.Equal(t, "/Users/myuser/Documents", suffixArgs[0])
}
