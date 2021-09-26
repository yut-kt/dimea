test:
	go test -v ./...

bench:
	go test -bench . -benchmem -count 5 -run none

cov:
	go test -cover

fmt:
	gofmt -l -s -w .
	goimports -w .
	golangci-lint run ./...