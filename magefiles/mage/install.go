package mage

import (
	"fmt"
	"github.com/fizyk/magex/file"
	"github.com/fizyk/magex/github"
	"github.com/fizyk/magex/http"
	"github.com/fizyk/magex/lib/golang"
	"github.com/hashicorp/go-version"
	"github.com/magefile/mage/sh"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var assetFilterMap = map[string]string{
	"linux":   "Linux-64bit.tar.gz",
	"windows": "Windows-64bit.zip",
}

const (
	tarFile              string = "mage.tar.gz"
	exec                 string = "mage"
	versionRegexpPattern string = "([\\d]+\\.[\\d]+(\\.[\\d]+)?)"
)

// filter filters asset name to correctly choose one for given Operating System and Architecture
func filter(downloadName string) bool {
	return strings.Contains(downloadName, assetFilterMap[runtime.GOOS])
}

// mageVersion reads mage's version
func mageVersion() (*version.Version, error) {
	versionOutput, err := sh.Output("mage", "-version")
	if err != nil {
		return nil, err
	}
	versionLineRegexp := regexp.MustCompile(fmt.Sprintf("Mage Build Tool %s", versionRegexpPattern))
	versionLine := versionLineRegexp.FindString(versionOutput)
	versionRegexp := regexp.MustCompile(versionRegexpPattern)
	versionFound := versionRegexp.FindString(versionLine)
	return version.NewVersion(versionFound)
}

// Install installs latest mage version
func Install() error {
	downloadURL, latestVersion, err := github.Latest("magefile", "mage", filter)
	fmt.Printf("Attempting to install version %s\n", latestVersion)
	if err != nil {
		return err
	}
	return install(latestVersion, downloadURL, golang.BinPath())
}

// InstallVersion installs given mage version
func InstallVersion(version string) error {
	downloadURL, givenVersion, err := github.Version(
		"magefile",
		"mage",
		fmt.Sprintf("v%s", version),
		filter,
	)
	if err != nil {
		return err
	}
	return install(givenVersion, downloadURL, golang.BinPath())
}

func install(latestVersion *version.Version, downloadURL, binPath string) error {
	if file.Exists(fmt.Sprintf("%s/%s", binPath, exec)) {
		currentVersion, err := mageVersion()
		if err != nil {
			fmt.Printf("Could not determine mage's version, got %s\n Attempting to do a fresh install.", err.Error())
		} else if latestVersion.LessThanOrEqual(currentVersion) {
			fmt.Printf("Version %s already installed\n", currentVersion.String())
			return nil
		} else {
			fmt.Printf("Updating from %s to %s\n", currentVersion.String(), latestVersion.String())
		}

	} else {
		fmt.Printf("Fresh installation of a version %s\n", latestVersion.String())
	}

	if err := http.DownloadFile(downloadURL, tarFile); err != nil {
		return err
	}
	if err := sh.Run("tar", "zxvf", tarFile, exec); err != nil {
		return err
	}
	if err := sh.Run("mv", "-f", exec, binPath); err != nil {
		return err
	}
	return os.Remove(tarFile)
}
