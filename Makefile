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
