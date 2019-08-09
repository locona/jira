.DEFAULT_GOAL := help

.PHONY: auth
auth:
	@go install
	@jira auth
