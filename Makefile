NAME=clig
BINARY=${NAME}
INSTALL_DIR=/usr/local/bin

VERSION=v1.0.4
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/ltheinrich/clig/internal/app/handlers.Version=${VERSION} -X github.com/ltheinrich/clig/internal/app/handlers.BuildTime=${BUILD_TIME}"
GO_FILES=./...

.PHONY: build
build: clean fmt test buildgo sign

.PHONY: install
install: ${BINARY}
	cp ${BINARY} ${INSTALL_DIR}

.PHONY: fmt
fmt:
	go fmt ${GO_FILES}

.PHONY: test
test:
	go vet -v ${GO_FILES}
	go test -v -race ${GO_FILES}

.PHONY: buildgo
buildgo:
	go build ${LDFLAGS} -o ${BINARY}

.PHONY: sign
sign:
	gpg2 -a --detach-sign ${BINARY}

.PHONY: installgo
installgo:
	go install ${LDFLAGS}

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm -f ${BINARY} ; fi
	if [ -f ${BINARY}.asc ] ; then rm -f ${BINARY}.asc ; fi
