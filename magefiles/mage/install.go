package mage

import (
	"fmt"
	"github.com/fizyk/magex/file"
	"github.com/fizyk/magex/github"
	"github.com/fizyk/magex/http"
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
	binPath              string = "/usr/local/bin"
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
	if err != nil {
		return err
	}
	return install(latestVersion, downloadURL)
}

// InstallVersion installs given mage version
func InstallVersion(version string) error {
	downloadURL, latestVersion, err := github.Version(
		"magefile",
		"mage",
		fmt.Sprintf("v%s", version),
		filter,
	)
	if err != nil {
		return err
	}
	return install(latestVersion, downloadURL)
}

func install(latestVersion *version.Version, downloadURL string) error {
	if file.Exists(fmt.Sprintf("%s/%s", binPath, exec)) {
		currentVersion, err := mageVersion()
		if err != nil {
			return err
		}
		if latestVersion.LessThanOrEqual(currentVersion) {
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
	if err := sh.Run("sudo", "mv", "-f", exec, binPath); err != nil {
		return err
	}
	return os.Remove(tarFile)
}
