.PHONY: test
test:
	@go test -coverprofile=coverprof.out -cover -v ./...

.PHONY: cover
cover: test
	@go tool cover -html=coverprof.out

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: gen
gen:
	@go generate ./...
