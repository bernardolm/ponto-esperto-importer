
setup:
	go get github.com/tools/godep
	go get github.com/liudng/dogo
	godep restore -v ./...

update-deps:
	godep save -v -t ./...
	rm -rf vendor

test:
	godep go test ./...

run:
	godep go run main.go --verbose

# install:
# 	apt-get update
# 	apt-get install -y pkg-config lxc libxml2 libxml2-dev

bin: godep
	GOOS=linux GOARCH=amd64 go build main.go

dogo:
	godep go test ./... # -run TestImporter

vet: # reports suspicious constructs
	godep go tool vet `pwd`
