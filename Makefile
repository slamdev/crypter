build:
	go build -o bin/crypter ./cmd/crypter

run:
	./bin/crypter encrypt -k /tmp/public.key -v some-text
	./bin/crypter decrypt -k /tmp/private.key -v "$$(cat /tmp/test.yaml)"

mod:
	go mod verify
	go mod tidy -v

tag:
	git fetch --tags
	git tag $(TAG)
	git push origin $(TAG)

untag:
	git fetch --tags
	git tag -d $(TAG)
	git push origin :refs/tags/$(TAG)
	curl --request DELETE --header "Authorization: token ${GITHUB_TOKEN}" "https://api.github.com/repos/slamdev/crypter/releases/:release_id/$(TAG)"

release:
ifeq ($(shell which /tmp/goreleaser),)
	curl -sL https://git.io/goreleaser | TMPDIR=/tmp bash -s -- -v
endif
	/tmp/goreleaser --rm-dist
