package github

import (
	"context"
	"errors"
	"github.com/google/go-github/v35/github"
	"github.com/hashicorp/go-version"
)

// Latest finds out the latest version of a repository release and download url
func Latest(owner, repo string, assetFilter func(string) bool) (string, *version.Version, error) {
	ghclient := github.NewClient(nil)
	ctx := context.Background()
	release, _, err := ghclient.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return "", nil, err
	}
	releaseVersion, err := version.NewVersion(*release.TagName)
	if err != nil {
		return "", nil, err
	}
	for _, asset := range release.Assets {
		if assetFilter(asset.GetName()) {
			return asset.GetBrowserDownloadURL(), releaseVersion, nil
		}
	}
	return "", nil, errors.New("no matching release found")
}
