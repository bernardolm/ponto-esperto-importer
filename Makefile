
setup:
	go get github.com/tools/godep
	godep restore -v ./...

update-deps:
	-rm -rf vendor
	go get -v -t
	go get -v -fix
	godep save -v -t ./...
	if [ -d Godeps ]; then godep update ./...; fi
	rm -rf vendor

test:
	reset
	godep go test ./...

run:
	godep go run main.go --verbose

bin: godep
	GOOS=linux GOARCH=amd64 go build main.go

dogo:
	go get github.com/liudng/dogo
	$(make) test

vet: # reports suspicious constructs
	godep go tool vet `pwd`
