.DEFAULT_GOAL := help

.PHONY: ns
ns:
	@go install
	@jira ns

.PHONY: auth
auth:
	@go install
	@jira auth

.PHONY: status
status:
	@go install
	@jira status

.PHONY: issue.list
issue.list:
	@go install
	@jira issue list --summary="rdap-"

.PHONY: issue.delete
issue.delete:
	@go install
	@jira issue delete

.PHONY: project.ns
project.ns:
	@go install
	@jira project ns

.PHONY: project.list
project.list:
	@go install
	@jira project list

.PHONY: project.show
project.show:
	@go install
	@jira project show

.PHONY: issue.delete.i
issue.delete.i:
	@go install
	@jira issue delete -i

.PHONY: issue.epic
issue.epic:
	@go install
	@jira issue epic

.PHONY: issue.transition
issue.transition:
	@go install
	@jira issue transition

.PHONY: issue.assign
issue.assign:
	@go install
	@jira issue assign

.PHONY: issue.apply
issue.apply:
	@go install
	@jira issue apply -f examples/issue_create.yaml

.PHONY: user.list
user.list:
	@go install
	@jira user list

.PHONY: user.show
user.show:
	@go install
	@jira user show
