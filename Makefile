tests:
	go test ./... -v
coverage:
	go test ./...  -race -covermode=atomic -coverprofile=coverage.out
vet:
	go vet ./...
