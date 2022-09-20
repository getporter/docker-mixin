package docker

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"get.porter.sh/porter/pkg/runtime"
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
)

const defaultDockerVersion = "20.10.7"

type Mixin struct {
	runtime.RuntimeConfig
	//add whatever other context/state is needed here
	schema        *packr.Box
	DockerVersion string
}

// New docker mixin client, initialized with useful defaults.
func New() *Mixin {
	return &Mixin{
		RuntimeConfig: runtime.NewConfig(),
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

func (m *Mixin) getOutput(ctx context.Context, outputName string) ([]byte, error) {
	cmd := m.NewCommand(ctx, "docker", "output", outputName)
	cmd.Stderr = m.Err

	out, err := cmd.Output()
	if err != nil {
		prettyCmd := fmt.Sprintf("%s %s", cmd.Path, strings.Join(cmd.Args, " "))
		return nil, errors.Wrap(err, fmt.Sprintf("couldn't run command %s", prettyCmd))
	}

	return out, nil
}

func (m *Mixin) handleOutputs(ctx context.Context, outputs []Output) error {
	for _, output := range outputs {
		bytes, err := m.getOutput(ctx, output.Name)
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
