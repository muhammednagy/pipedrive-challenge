FILES?=main.go
BUILDTIME=`date "+%F %T%Z"`
VERSION=`git describe --tags`

build:
	go build -ldflags="-X 'github.com/muhammednagy/pipedrive-challenge/config.buildTime=$(BUILDTIME)' -X 'github.com/muhammednagy/pipedrive-challenge/config.version=$(VERSION)' -s -w" -o pipedrive $(FILES)

run:
	go run $(FILES)

clean:
	rm -f pipedrive

test:
	go test -p 1 ./...

cover:
	go test ./... -coverprofile cover.out
