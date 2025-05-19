# for updating the version of the project and pushing the tag to the repository
VERSION = 2.0.1

MODULE_NAME = github.com/fengdotdev/golibs-helperfuncs

#my GOLIBS For all projects 
TRAITS = github.com/fengdotdev/golibs-traits
TESTING = github.com/fengdotdev/golibs-testing
KWOW = github.com/fengdotdev/golibs-kwowledge
NDRIVE = github.com/fengdotdev/golibs-nativedrive
VDRIVE = github.com/fengdotdev/golibs-vdrive
COMMON = github.com/fengdotdev/golibs-commontypes
HELPER = github.com/fengdotdev/golibs-helperfuncs
ONEDRIVE = github.com/fengdotdev/golibs-1driveclient
BRIDGE = github.com/fengdotdev/golibs-bridge
ML = github.com/fengdotdev/golibs-ml
STAT = github.com/fengdotdev/golibs-statistics
LA = github.com/fengdotdev/golibs-linealalgebra
DUMMY = github.com/fengdotdev/golibs-dummy
DC = github.com/fengdotdev/golibs-datacontainer
SPAGE = github.com/fengdotdev/golibs-staticpages




.PHONY: gi gg gv  gt

# create a new go module
gi:
	go mod init $(MODULE_NAME)

# update the go.mod file with the latest dependencies
gg:
	go get $(TRAITS)@latest
	go get $(TESTING)@latest
	go get $(HELPER)@latest

# update the version of the project
gv:
		git tag v${VERSION} && git push origin v${VERSION}

# run all tests
gt:
	go clean -testcache
	go test -count=1 -v ./...



