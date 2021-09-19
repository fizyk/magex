//go:build mage
// +build mage

package main

import (
	// mage:import go
	_ "github.com/fizyk/magex/magefiles/golang"
	// mage:import go:check
	_ "github.com/fizyk/magex/magefiles/golang/check"
	// mage:import mage
	_ "github.com/fizyk/magex/magefiles/mage"
)
