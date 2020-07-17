package docker

import (
	"get.porter.sh/porter/pkg/exec/builder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sort"
	"testing"
)

func TestMixin_UnmarshalStep(t *testing.T) {
	testcases := []struct {
		name        string // Test case name
		file        string // Path to th test input yaml
		wantDescription string // Description that you expect to be found
		wantArguments []string // Arguments that you expect to be found
		wantFlags builder.Flags // Flags that you expect to be found
		wantSuffixArgs []string // Suffix arguments that you expect to be found

	} {
		{"pull", "testdata/pull-input.yaml", "Prefetch the things",
			[]string{"pull", "getporter/porter-hello:v0.1.0"}, nil, nil},
		{"push", "testdata/push-input.yaml", "Push to registry",
			[]string{"push", "getporter/porter-hello:v0.1.0"}, nil, nil},
		{"build", "testdata/build-input.yaml", "Build image",
			[]string{"build"}, builder.Flags{ builder.NewFlag("f", "myfile"), builder.NewFlag("t", "practice")}, []string{"/Users/myuser/Documents"}},
		{"run", "testdata/run-input.yaml", "Run container",
			[]string{"run"}, builder.Flags{ builder.NewFlag("d"), builder.NewFlag("env", "password=password"), builder.NewFlag("name", "practice"), builder.NewFlag("privileged"), builder.NewFlag("rm")}, []string{"getporter/porter-hello"}},
		{"remove", "testdata/remove-input.yaml", "Remove container",
			[]string{"rm"}, builder.Flags{ builder.NewFlag("f")}, []string{"practice"}},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			//build test
			b, err := ioutil.ReadFile(tc.file)
			require.NoError(t, err)

			var action Action
			err = yaml.Unmarshal(b, &action)
			require.NoError(t, err)
			require.Len(t, action.Steps, 1)

			step := action.Steps[0]
			assert.Equal(t, tc.wantDescription, step.Description)

			args := step.GetArguments()
			// docker build -t tag -f file path ARGUMENTS --FLAGS
			assert.Equal(t, tc.wantArguments, args)

			flags := step.GetFlags()
			sort.Sort(flags)
			assert.Equal(t, tc.wantFlags, flags)


			suffixArgs := step.GetSuffixArguments()
			assert.Equal(t, tc.wantSuffixArgs, suffixArgs)


		})
	}
}
