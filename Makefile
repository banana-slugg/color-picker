build:
	@go build -o ./bin/color-picker ./cmd/...

run: build
	@./bin/color-picker

tests:
	@go test -v ./test/...


