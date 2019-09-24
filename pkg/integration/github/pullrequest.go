package github

import (
	"context"

	"github.com/google/go-github/v27/github"
	"github.com/locona/jira/pkg/gitconfig"
)

func PullRequests(state string) ([]*github.PullRequest, error) {
	ctx := context.Background()
	cli := Client()
	gc, err := gitconfig.Config()
	pullrequests, _, err := cli.PullRequests.List(
		ctx,
		gc.RemoteConfig.Organization,
		gc.RemoteConfig.Repository,
		&github.PullRequestListOptions{
			State: state,
		},
	)
	if err != nil {
		return nil, err
	}

	return pullrequests, nil
}

func PullRequestCommits(number int) ([]*github.RepositoryCommit, error) {
	ctx := context.Background()
	cli := Client()
	gc, err := gitconfig.Config()
	commits, _, err := cli.PullRequests.ListCommits(ctx, gc.RemoteConfig.Organization, gc.RemoteConfig.Repository, number, nil)
	if err != nil {
		return nil, err
	}

	return commits, nil
}
