package remove

import (
	"github.com/devspace-cloud/devspace-cloud-plugin/pkg/factory"
	"github.com/spf13/cobra"
)

// NewRemoveCmd creates a new cobra command
func NewRemoveCmd(f factory.Factory) *cobra.Command {
	removeCmd := &cobra.Command{
		Use:   "remove",
		Short: "Changes devspace configuration",
		Long: `
#######################################################
################## devspace remove ####################
#######################################################
	`,
		Args: cobra.NoArgs,
	}

	removeCmd.AddCommand(newClusterCmd(f))
	removeCmd.AddCommand(newProviderCmd(f))
	removeCmd.AddCommand(newSpaceCmd(f))
	return removeCmd
}
