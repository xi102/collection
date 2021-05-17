BINARY="xi102"

GITTAG=`git describe --long --tags --dirty --always`
BUILD_TIME=`date +%FT%T%z`
LDFLAGS=-ldflags "-X main.GitTag=${GITTAG} -X main.BuildTime=${BUILD_TIME}"

PACKAGES=`go list ./... | grep -v /vendor/`
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

build: 
	go build ${LDFLAGS} -o ${BINARY} main.go

buildfront:
	cd frontend
	yarn build
	cd ../..

run:
	go run main.go

runfront:
	cd frontend
	yarn serve

clean:
	@echo "  >  Cleaning build cache"
	# go clean -cache -modcache -i -r
	go clean -cache