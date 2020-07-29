package vars

import (
	"fmt"
	cloudconfig "github.com/devspace-cloud/devspace-cloud-plugin/pkg/cloud/config"
	"github.com/devspace-cloud/devspace-cloud-plugin/pkg/factory"
	"github.com/mgutz/ansi"
	"os"

	"github.com/spf13/cobra"
)

type spaceNamespaceCmd struct{}

func newSpaceNamespaceCmd(f factory.Factory) *cobra.Command {
	cmd := &spaceNamespaceCmd{}

	return &cobra.Command{
		Use:   "space-namespace",
		Short: "Prints the current spaces namespace",
		Args: cobra.NoArgs,
		RunE: func(cobraCmd *cobra.Command, args []string) error {
			return cmd.Run(f, cobraCmd, args)
		},
	}
}

// Run executes the command logic
func (*spaceNamespaceCmd) Run(f factory.Factory, cobraCmd *cobra.Command, args []string) error {
	retErr := fmt.Errorf("Current context is not a space, but predefined var DEVSPACE_SPACE_NAMESPACE is used.\n\nPlease run: \n- `%s` to create a new space\n- `%s` to use an existing space\n- `%s` to list existing spaces", ansi.Color("devspace create space [NAME]", "white+b"), ansi.Color("devspace use space [NAME]", "white+b"), ansi.Color("devspace list spaces", "white+b"))
	kubeLoader := f.NewKubeConfigLoader()
	kubeContext, err := kubeLoader.GetCurrentContext()
	if err != nil {
		return err
	}
	if os.Getenv("DEVSPACE_PLUGIN_KUBE_CONTEXT_FLAG") != "" {
		kubeContext = os.Getenv("DEVSPACE_PLUGIN_KUBE_CONTEXT_FLAG")
	}

	isSpace, err := kubeLoader.IsCloudSpace(kubeContext)
	if err != nil || !isSpace {
		return retErr
	}

	spaceID, providerName, err := kubeLoader.GetSpaceID(kubeContext)
	if err != nil {
		return err
	}

	loader := cloudconfig.NewLoader()
	cloudConfigData, err := loader.Load()
	if err != nil {
		return retErr
	}

	provider := cloudconfig.GetProvider(cloudConfigData, providerName)
	if provider == nil {
		return retErr
	}
	if provider.Spaces == nil {
		return retErr
	}
	if provider.Spaces[spaceID] == nil {
		return retErr
	}

	_, err = os.Stdout.Write([]byte(provider.Spaces[spaceID].ServiceAccount.Namespace))
	return err
}