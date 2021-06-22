darwin:
	GOOS=darwin GOARCH=amd64 go build -o gmng-amd64-darwin .

linux:
	GOOS=linux GOARCH=amd64 go build -o gmng-amd64-linux .

install:
	go install .

.PHONY: darwin linux install
