package resume

import (
	"github.com/devspace-cloud/devspace-cloud-plugin/pkg/factory"
	"github.com/pkg/errors"
	"os"

	"github.com/spf13/cobra"
)

type spaceCmd struct{}

func newSpaceCmd(f factory.Factory) *cobra.Command {
	cmd := &spaceCmd{}

	return &cobra.Command{
		Use:   "space",
		Short: "Prints the current space",
		Args:  cobra.NoArgs,
		RunE: func(cobraCmd *cobra.Command, args []string) error {
			return cmd.Run(f, cobraCmd, args)
		},
	}
}

// Run executes the command logic
func (*spaceCmd) Run(f factory.Factory, cobraCmd *cobra.Command, args []string) error {
	// Get kubectl client
	client, err := f.NewKubeClientFromContext(os.Getenv("DEVSPACE_PLUGIN_KUBE_CONTEXT_FLAG"), os.Getenv("DEVSPACE_PLUGIN_KUBE_NAMESPACE_FLAG"), false)
	if err != nil {
		return errors.Wrap(err, "new kube client")
	}

	// Signal that we are working on the space if there is any
	return f.NewSpaceResumer(client, f.GetLog()).ResumeSpace(true)
}
