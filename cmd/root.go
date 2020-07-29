package cmd

import (
	"github.com/devspace-cloud/devspace-cloud-plugin/cmd/add"
	"github.com/devspace-cloud/devspace-cloud-plugin/cmd/connect"
	"github.com/devspace-cloud/devspace-cloud-plugin/cmd/create"
	"github.com/devspace-cloud/devspace-cloud-plugin/cmd/flags"
	"github.com/devspace-cloud/devspace-cloud-plugin/cmd/list"
	"github.com/devspace-cloud/devspace-cloud-plugin/cmd/remove"
	"github.com/devspace-cloud/devspace-cloud-plugin/cmd/reset"
	"github.com/devspace-cloud/devspace-cloud-plugin/cmd/set"
	"github.com/devspace-cloud/devspace-cloud-plugin/cmd/use"
	"github.com/devspace-cloud/devspace-cloud-plugin/pkg/factory"
	"github.com/devspace-cloud/devspace-cloud-plugin/pkg/upgrade"
	"github.com/devspace-cloud/devspace/pkg/util/exit"
	flagspkg "github.com/devspace-cloud/devspace/pkg/util/flags"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// NewRootCmd returns a new root command
func NewRootCmd(f factory.Factory) *cobra.Command {
	return &cobra.Command{
		Use:           "devspace-cloud-plugin",
		SilenceUsage:  true,
		SilenceErrors: true,
		Short:         "Welcome to the DevSpace Cloud!",
		PersistentPreRun: func(cobraCmd *cobra.Command, args []string) {
			log := f.GetLog()
			if globalFlags.Silent {
				log.SetLevel(logrus.FatalLevel)
			}

			// apply extra flags
			if cobraCmd.DisableFlagParsing == false {
				extraFlags, err := flagspkg.ApplyExtraFlags(cobraCmd, os.Args, false)
				if err != nil {
					log.Warnf("Error applying extra flags: %v", err)
				} else if len(extraFlags) > 0 {
					log.Infof("Applying extra flags from environment: %s", strings.Join(extraFlags, " "))
				}
			}

			// Get version of current binary
			latestVersion := upgrade.NewerVersionAvailable()
			if latestVersion != "" {
				log.Warnf("There is a newer version of the DevSpace Cloud Plugin: v%s. Run `devspace update plugin devspace-cloud` to upgrade to the newest version.\n", latestVersion)
			}
		},
		Long: `DevSpace accelerates developing, deploying and debugging applications with Docker and Kubernetes. Get started by running the init command in one of your projects:
	
		devspace init`,
	}
}

var globalFlags *flags.GlobalFlags

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	f := factory.DefaultFactory()

	// build the root command
	rootCmd := BuildRoot(f)

	// set version for --version flag
	rootCmd.Version = upgrade.GetVersion()

	// execute command
	err := rootCmd.Execute()
	if err != nil {
		// Check if return code error
		retCode, ok := errors.Cause(err).(*exit.ReturnCodeError)
		if ok {
			os.Exit(retCode.ExitCode)
		}

		if globalFlags.Debug {
			f.GetLog().Fatalf("%+v", err)
		} else {
			f.GetLog().Fatal(err)
		}
	}
}

// BuildRoot creates a new root command from the
func BuildRoot(f factory.Factory) *cobra.Command {
	rootCmd := NewRootCmd(f)
	persistentFlags := rootCmd.PersistentFlags()
	globalFlags = flags.SetGlobalFlags(persistentFlags)

	// Add sub commands
	rootCmd.AddCommand(add.NewAddCmd(f))
	rootCmd.AddCommand(create.NewCreateCmd(f))
	rootCmd.AddCommand(use.NewUseCmd(f))
	rootCmd.AddCommand(connect.NewConnectCmd(f))
	rootCmd.AddCommand(list.NewListCmd(f))
	rootCmd.AddCommand(remove.NewRemoveCmd(f))
	rootCmd.AddCommand(reset.NewResetCmd(f))
	rootCmd.AddCommand(set.NewSetCmd(f))

	// Add base commands
	rootCmd.AddCommand(NewLoginCmd(f))
	return rootCmd
}
