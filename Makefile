FILES?=main.go
BUILDTIME=`date "+%F %T%Z"`
VERSION=`git describe --tags`

build:
	go build -ldflags="-X 'github.com/muhammednagy/pipedirve-challenge/config.buildTime=$(BUILDTIME)' -X 'github.com/muhammednagy/pipedirve-challenge/config.version=$(VERSION)' -s -w" -o pipedrive $(FILES)

run:
	go run $(FILES)

clean:
	rm -rf bin

test:
	go test ./...

cover:
	go test ./... -coverprofile cover.out
