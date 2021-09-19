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
	return getRelease(release, assetFilter)
}

// Version finds out the give version of a repository release and download url
func Version(owner, repo, repoVersion string, assetFilter func(string) bool) (string, *version.Version, error) {
	ghclient := github.NewClient(nil)
	ctx := context.Background()
	release, _, err := ghclient.Repositories.GetReleaseByTag(ctx, owner, repo, repoVersion)
	if err != nil {
		return "", nil, err
	}
	return getRelease(release, assetFilter)
}

func getRelease(release *github.RepositoryRelease, assetFilter func(string) bool) (string, *version.Version, error) {
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
