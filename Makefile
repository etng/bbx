VERSION=0.0.4
NAME=bbx
CMD_NAME=bbx
DESC="bbx is abbr for 百宝箱(bai bao xiang), it just means a toolbox"
USAGE="bbx -h"
BREW_TAP="etng/taps/"
AUTHOR="etng"
AUTHOR_EMAIL="etng2004@gmail.com"
local:
	goreleaser release --snapshot --skip-publish --rm-dist
run:
	go run main.go
run_local:
	dist/${NAME}_darwin_amd64_v1/${NAME}
init:
	rm -f .goreleaser.yaml || true
	rm -fr .git || true
	rm -fr go.mod || true
	touch .gitignore
	@echo "">.gitignore
	goreleaser init
	@echo 'project_name: ' ${NAME} >> .goreleaser.yaml
	mkdir -p .github/workflows
	cp ~/Documents/github_action_goreleaser.yml .github/workflows/goreleaser.yml
	git init .
	git remote add origin git@github.com:etng/${NAME}.git
	touch README.md
	@echo '' > README.md
	@echo '# ' ${NAME} >> README.md
	@echo ${DESC} >> README.md
	@echo '## usage' >> README.md
	@echo '```shell' >> README.md
	@echo brew install ${BREW_TAP}${CMD_NAME} >> README.md
	@echo ${USAGE} >> README.md
	@echo '```' >> README.md

	@echo ".idea">>.gitignore
	@echo ".commit_comment">>.gitignore
	@echo 'welcome to the reality!'>.commit_comment
	git config user.email ${AUTHOR}
	git config user.name ${AUTHOR_EMAIL}
	git config pull.rebase true
	go mod init github.com/etng/${CMD_NAME}
tap:
	@echo 'brews:'>>.goreleaser.yaml
	@echo '- name: ' ${NAME} >>.goreleaser.yaml
	@echo '  homepage: https://github.com/etng/homebrew-taps'>>.goreleaser.yaml
	@echo '  tap:'>>.goreleaser.yaml
	@echo '    owner: etng'>>.goreleaser.yaml
	@echo '    name: homebrew-taps'>>.goreleaser.yaml
	@echo "tap patched"
tag:
	git tag v${VERSION} -m "v${VERSION}" -f
	git push origin master -u -f --tags
commit:
	git add .
	git commit -am "$(shell cat .commit_comment)"
	git push origin master -u -f