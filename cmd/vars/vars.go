package vars

import (
	"github.com/devspace-cloud/devspace-cloud-plugin/pkg/factory"
	"github.com/spf13/cobra"
)

// NewVarsCmd creates a new cobra command for the sub command
func NewVarsCmd(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vars",
		Short: "Print predefined variables",
		Long: `
#######################################################
################### devspace vars #####################
#######################################################
	`,
		Args: cobra.NoArgs,
	}

	cmd.AddCommand(newUsernameCmd(f))
	cmd.AddCommand(newSpaceNamespaceCmd(f))
	cmd.AddCommand(newSpaceCmd(f))
	return cmd
}
