go-install:
	go install github.com/vektra/mockery/v2@latest
	go get -u -t -v ./... || :
	go work sync

go-gen: go-install
	go generate ./...
	go test ./...
	go mod tidy || :
