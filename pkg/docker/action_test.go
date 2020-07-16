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
	//pull test
	// b, err := ioutil.ReadFile("testdata/pull-input.yaml")
	// require.NoError(t, err)

	// var action Action
	// err = yaml.Unmarshal(b, &action)
	// require.NoError(t, err)
	// require.Len(t, action.Steps, 1)

	// step := action.Steps[0]
	// assert.Equal(t, "Prefetch the things", step.Description)

	// /*
	// 	assert.NotEmpty(t, step.Outputs)
	// 	assert.Equal(t, Output{Name: "VICTORY", JsonPath: "$Id"}, step.Outputs[0])
	// */

	// args := step.GetArguments()
	// // docker pull getporter/porter-hello:v0.1.0
	// require.Len(t, args, 2)
	// assert.Equal(t, "pull", args[0])
	// assert.Equal(t, "getporter/porter-hello:v0.1.0", args[1])

	// flags := step.GetFlags()
	// assert.Empty(t, flags)

	//push test
	//b, err := ioutil.ReadFile("testdata/push-input.yaml")
	//require.NoError(t, err)
	//
	//var action Action
	//err = yaml.Unmarshal(b, &action)
	//require.NoError(t, err)
	//require.Len(t, action.Steps, 1)
	//
	//step := action.Steps[0]
	//assert.Equal(t, "Push to registry", step.Description)
	//
	///*
	//	assert.NotEmpty(t, step.Outputs)
	//	assert.Equal(t, Output{Name: "VICTORY", JsonPath: "$Id"}, step.Outputs[0])
	//*/
	//
	//args := step.GetArguments()
	//// docker push getporter/porter-hello:v0.1.0
	//require.Len(t, args, 2)
	//assert.Equal(t, "push", args[0])
	//assert.Equal(t, "getporter/porter-hello:v0.1.0", args[1])
	//
	//flags := step.GetFlags()
	//assert.Empty(t, flags)

	//run test
	//b, err := ioutil.ReadFile("testdata/run-input.yaml")
	//require.NoError(t, err)
	//
	//var action Action
	//err = yaml.Unmarshal(b, &action)
	//require.NoError(t, err)
	//require.Len(t, action.Steps, 1)
	//
	//step := action.Steps[0]
	//assert.Equal(t, "Run container", step.Description)
	//
	///*
	//	assert.NotEmpty(t, step.Outputs)
	//	assert.Equal(t, Output{Name: "VICTORY", JsonPath: "$Id"}, step.Outputs[0])
	//*/
	//
	//args := step.GetArguments()
	//// docker run --privileged --env VAR1=value1 --env VAR2=value2 --rm -p host:container -d --name name image ARGUMENTS --FLAGS
	//require.Len(t, args, 3)
	//assert.Equal(t, "run", args[0])
	//assert.Equal(t, "--env password=password -p 8080:80", args[1])
	////assert.Equal(t, "--env password=pleasework", args[2])
	////assert.Equal(t, "", args[3])
	////assert.Equal(t, "-p 8080:80", args[4])
	////assert.Equal(t, "", args[5])
	//assert.Equal(t, "getporter/porter-hello", args[2])
	//
	//flags := step.GetFlags()
	//assert.Empty(t, flags)

	//login test
	//b, err := ioutil.ReadFile("testdata/login-input.yaml")
	//require.NoError(t, err)
	//
	//var action Action
	//err = yaml.Unmarshal(b, &action)
	//require.NoError(t, err)
	//require.Len(t, action.Steps, 1)
	//
	//step := action.Steps[0]
	//assert.Equal(t, "Login to docker", step.Description)
	//
	///*
	//	assert.NotEmpty(t, step.Outputs)
	//	assert.Equal(t, Output{Name: "VICTORY", JsonPath: "$Id"}, step.Outputs[0])
	//*/
	//
	//args := step.GetArguments()
	//// docker push getporter/porter-hello:v0.1.0
	//require.Len(t, args, 5)
	//assert.Equal(t, "login", args[0])
	//assert.Equal(t, "-u", args[1])
	//assert.Equal(t, "gmadhok", args[2])
	//assert.Equal(t, "-p", args[3])
	//assert.Equal(t, "password", args[4])
	//
	//flags := step.GetFlags()
	//assert.Empty(t, flags)

	////remove test
	//b, err := ioutil.ReadFile("testdata/remove-input.yaml")
	//require.NoError(t, err)
	//
	//var action Action
	//err = yaml.Unmarshal(b, &action)
	//require.NoError(t, err)
	//require.Len(t, action.Steps, 1)
	//
	//step := action.Steps[0]
	//assert.Equal(t, "Remove container", step.Description)
	//
	///*
	//	assert.NotEmpty(t, step.Outputs)
	//	assert.Equal(t, Output{Name: "VICTORY", JsonPath: "$Id"}, step.Outputs[0])
	//*/
	//
	//args := step.GetArguments()
	//// docker run --privileged --env VAR1=value1 --env VAR2=value2 --rm -p host:container -d --name name image ARGUMENTS --FLAGS
	//require.Len(t, args, 2)
	//assert.Equal(t, "rm", args[0])
	//assert.Equal(t, "practice", args[1])
	////assert.Equal(t, "rm", args[0])
	////assert.Equal(t, "-f", args[1])
	////assert.Equal(t, "practice", args[2])
	////assert.Equal(t, "--env password=pleasework", args[2])
	////assert.Equal(t, "", args[3])
	////assert.Equal(t, "-p 8080:80", args[4])
	////assert.Equal(t, "", args[5])
	////assert.Equal(t, "getporter/porter-hello", args[2])
	//
	//flags := step.GetFlags()
	//assert.Empty(t, flags)

	//build test
	b, err := ioutil.ReadFile("testdata/build-input.yaml")
	require.NoError(t, err)

	var action Action
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)
	require.Len(t, action.Steps, 1)

	step := action.Steps[0]
	assert.Equal(t, "Build image", step.Description)

	/*
		assert.NotEmpty(t, step.Outputs)
		assert.Equal(t, Output{Name: "VICTORY", JsonPath: "$Id"}, step.Outputs[0])
	*/

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
	//flags := step.GetFlags()
	//assert.Equal(t, "-f myfile", flags[0])
	//assert.Equal(t, "-t practice ", flags[1])
	//assert.Empty(t, flags)

	suffixArgs := step.GetSuffixArguments()
	require.Len(t, suffixArgs, 1)
	assert.Equal(t, "/Users/myuser/Documents", suffixArgs[0])
}
