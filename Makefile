BINARY="xi102"

GITTAG=`git describe --tags`
BUILD_TIME=`date +%FT%T%z`
LDFLAGS=-ldflags "-X main.GitTag=${GITTAG} -X main.BuildTime=${BUILD_TIME}"

build: 
	go build ${LDFLAGS} -o ${BINARY} main.go

run:
	go run main.go