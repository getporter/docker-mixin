//go:generate packr2
package docker

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"

	"get.porter.sh/porter/pkg/context" // We are not using go-yaml because of serialization problems with jsonschema, don't use this library elsewhere
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
)

const defaultDockerVersion = "19.03.8"

type Mixin struct {
	*context.Context
	//add whatever other context/state is needed here
	schema        *packr.Box
	DockerVersion string
}

// New docker mixin client, initialized with useful defaults.
func New() *Mixin {
	return &Mixin{
		Context:       context.New(),
		schema:        packr.New("schema", "./schema"),
		DockerVersion: defaultDockerVersion,
	}
}

func (m *Mixin) getPayloadData() ([]byte, error) {
	reader := bufio.NewReader(m.In)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "could not read the payload from STDIN")
	}
	return data, nil
}

func (m *Mixin) getOutput(outputName string) ([]byte, error) {
	cmd := m.NewCommand("docker", "output", outputName)
	cmd.Stderr = m.Err

	out, err := cmd.Output()
	if err != nil {
		prettyCmd := fmt.Sprintf("%s %s", cmd.Path, strings.Join(cmd.Args, " "))
		return nil, errors.Wrap(err, fmt.Sprintf("couldn't run command %s", prettyCmd))
	}

	return out, nil
}

func (m *Mixin) handleOutputs(outputs []Output) error {
	for _, output := range outputs {
		bytes, err := m.getOutput(output.Name)
		if err != nil {
			return err
		}

		err = m.Context.WriteMixinOutputToFile(output.Name, bytes)
		if err != nil {
			return errors.Wrapf(err, "unable to write output '%s'", output.Name)
		}
	}
	return nil
}
