# for updating the version of the project and pushing the tag to the repository
VERSION = 2.0.0

updatev:
		git tag v${VERSION} && git push origin v${VERSION}

test:
	go clean -testcache
	go test -count=1 -v ./...



