module github.com/devspace-cloud/devspace-cloud-plugin

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/daviddengcn/go-colortext v1.0.0 // indirect
	github.com/devspace-cloud/devspace v1.1.1-0.20200724074930-ec77a1851818
	github.com/joho/godotenv v1.3.0
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b
	github.com/pkg/errors v0.9.1
	github.com/rhysd/go-github-selfupdate v0.0.0-20180520142321-41c1bbb0804a
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	k8s.io/apimachinery v0.18.6 // indirect
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.0+incompatible
	github.com/agl/ed25519 => github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412
	golang.org/x/sys => golang.org/x/sys v0.0.0-20191128015809-6d18c012aee9
)

go 1.13
