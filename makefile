# for updating the version of the project and pushing the tag to the repository
HELPERFUNCS_VERSION = 0.0.2

updatev:
		git tag v${HELPERFUNCS_VERSION} && git push origin v${HELPERFUNCS_VERSION}

test:
	go clean -testcache
	go test -count=1 -v ./...



