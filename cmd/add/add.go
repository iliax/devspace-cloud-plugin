package add

import (
	"github.com/devspace-cloud/devspace-cloud-plugin/pkg/factory"
	"github.com/spf13/cobra"
)

// NewAddCmd creates a new cobra command
func NewAddCmd(f factory.Factory) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a provider",
		Long: `
#######################################################
#################### devspace add #####################
#######################################################
Adds config sections to devspace.yaml
	`,
		Args: cobra.NoArgs,
	}

	addCmd.AddCommand(newProviderCmd(f))
	return addCmd
}
