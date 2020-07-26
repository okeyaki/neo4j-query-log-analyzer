.PHONY: install
install:
	$(MAKE) install-git-hooks

	git secrets --register-aws

.PHONY: install-git-hooks
install-git-hooks:
	cp etc/git/hook/commit-msg .git/hooks/
	cp etc/git/hook/pre-commit .git/hooks/
	cp etc/git/hook/prepare-commit-msg .git/hooks/

	chmod +x .git/hooks/commit-msg
	chmod +x .git/hooks/pre-commit
	chmod +x .git/hooks/prepare-commit-msg

.PHONY: build
build:
	go build -o var/bin/neo4j-query-log-analyzer bin/main.go
