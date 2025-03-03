# for updating the version of the project and pushing the tag to the repository
TESTING_VERSION = 0.0.1

updatev:
		git tag v${TESTING_VERSION} && git push origin v${TESTING_VERSION}

test:
	go clean -testcache
	go test -count=1 -v ./...




PLAYGROUD = ./sandbox/openssl_test
FILE = hello.txt
ENCODE = test.enc
DECODE = test.dec
PASSWORD = password
# openssl dont support the gcm mode for aes-256 
o:	
	rm -rf ${PLAYGROUD}/${ENCODE} ${PLAYGROUD}/${DECODE}
	openssl enc -aes-256-gcm -in  ${PLAYGROUD}/${FILE} -out ${PLAYGROUD}/${ENCODE} -k ${PASSWORD}
	openssl enc -d --aes-256-gcm -in ${PLAYGROUD}/${ENCODE} -out ${PLAYGROUD}/${DECODE} -k ${PASSWORD}

sand:
	go run ./sandbox/aesgcm/main.go


run:
	go run ./sandbox/aesgcm/server/main.go