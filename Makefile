install:
	go get -u github.com/gobuffalo/packr/v2/packr2

build:
	${HOME}/go/bin/packr2 build

release:
	-${HOME}/go/bin/goreleaser build
	${HOME}/go/bin/goreleaser release

clean:
	${HOME}/go/bin/packr2 clean

build-linux:
	GOOS=linux GOARCH=amd64 ${HOME}/go/bin/packr2 build

build-win:
	GOOS=windows GOARCH=amd64 ${HOME}/go/bin/packr2 build