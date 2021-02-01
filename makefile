

run:
	@go run cmd/kaex/*.go

test:
	@go test ./...

test-update:
	@go test -update ./...

./build:
	@go build -o build/kaex cmd/kaex/*.go

install:
	cd cmd/kaex && go install

clean:
	@rm -r ./build
