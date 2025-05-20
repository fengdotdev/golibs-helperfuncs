### DEPRECATED USE github.com/fengdotdev/golibs-funcs

VERSION = 1.0.2

MODULE_NAME = github.com/fengdotdev/golibs-helperfuncs


.PHONY: test tag



# update the version of the project
tag:
		git tag v${VERSION} && git push origin v${VERSION}

# run all tests
test:
	go clean -testcache
	go test -count=1 -v ./...



