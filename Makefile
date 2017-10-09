
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

bin: setup
	GOOS=linux GOARCH=amd64 go build -o ponto-esperto-importer

dogo:
	go get github.com/liudng/dogo
	$(MAKE) test

vet: # reports suspicious constructs
	godep go tool vet `pwd`

linter:
	go get github.com/alecthomas/gometalinter
	gometalinter --install
	gometalinter --config=gometalinter.json ./... | grep -v comment > gometalinter.txt

lint:
	go get github.com/golang/lint/golint
	golint ./... | grep -v comment > golint.txt
