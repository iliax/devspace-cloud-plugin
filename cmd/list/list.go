package list

import (
	"github.com/devspace-cloud/devspace-cloud-plugin/pkg/factory"
	"github.com/spf13/cobra"
)

// NewListCmd creates a new cobra command
func NewListCmd(f factory.Factory) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Lists configuration",
		Long: `
#######################################################
#################### devspace list ####################
#######################################################
	`,
		Args: cobra.NoArgs,
	}

	listCmd.AddCommand(newSpacesCmd(f))
	listCmd.AddCommand(newClustersCmd(f))
	listCmd.AddCommand(newProvidersCmd(f))
	return listCmd
}
