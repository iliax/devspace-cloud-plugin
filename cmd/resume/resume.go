package resume

import (
	"github.com/devspace-cloud/devspace-cloud-plugin/pkg/factory"
	"github.com/spf13/cobra"
)

// NewResumeCmd creates a new cobra command for the sub command
func NewResumeCmd(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resume",
		Short: "Resume spaces",
		Long: `
#######################################################
################## devspace resume ####################
#######################################################
	`,
		Args: cobra.NoArgs,
	}

	cmd.AddCommand(newSpaceCmd(f))
	return cmd
}
