package use

import (
	"github.com/devspace-cloud/devspace-cloud-plugin/pkg/factory"
	"github.com/spf13/cobra"
)

// NewUseCmd creates a new cobra command for the use sub command
func NewUseCmd(f factory.Factory) *cobra.Command {
	useCmd := &cobra.Command{
		Use:   "use",
		Short: "Use specific config",
		Long: `
#######################################################
#################### devspace use #####################
#######################################################
	`,
		Args: cobra.NoArgs,
	}

	useCmd.AddCommand(newProviderCmd(f))
	useCmd.AddCommand(newSpaceCmd(f))
	return useCmd
}
