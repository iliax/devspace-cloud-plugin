package main

import (
	"os"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/devspace-cloud/devspace-cloud-plugin/cmd"
	"github.com/devspace-cloud/devspace-cloud-plugin/pkg/upgrade"
)

var version string

func main() {
	upgrade.SetVersion(version)

	cmd.Execute()
	os.Exit(0)
}
