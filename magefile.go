//go:build mage

package main

import (
	"get.porter.sh/magefiles/mixins"
)

const (
	mixinName    = "docker"
	mixinPackage = "get.porter.sh/mixin/docker"
	mixinBin     = "bin/mixins/" + mixinName
)

var magefile = mixins.NewMagefile(mixinPackage, mixinName, mixinBin)

func ConfigureAgent() {
	magefile.ConfigureAgent()
}

// Build the mixin
func Build() {
	magefile.Build()
}

// Cross-compile the mixin before a release
func XBuildAll() {
	magefile.XBuildAll()
}

// Run unit tests
func TestUnit() {
	magefile.TestUnit()
}

func Test() {
	magefile.Test()
}

// Publish the mixin to github
func Publish() {
	magefile.Publish()
}

// Test the publish logic against your github fork
func TestPublish(username string) {
	magefile.TestPublish(username)
}

// Install the mixin
func Install() {
	magefile.Install()
}

// Remove generated build files
func Clean() {
	magefile.Clean()
}
