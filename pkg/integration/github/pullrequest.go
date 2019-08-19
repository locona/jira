package github

import (
	"context"

	"github.com/3-shake/jira/pkg/gitconfig"
	"github.com/google/go-github/v27/github"
)

func PullRequests() ([]*github.PullRequest, error) {
	ctx := context.Background()
	cli := Client()
	gc, err := gitconfig.Config()
	pullrequests, _, err := cli.PullRequests.List(ctx, gc.RemoteConfig.Organization, gc.RemoteConfig.Repository, nil)
	if err != nil {
		return nil, err
	}

	return pullrequests, nil
}
