

run:
	@go run cmd/kaex/*.go

./build:
	@go build -o build/kaex cmd/kaex/*.go

install:
	go install cmd/kaex/*.go

clean:
	@rm -r ./build
