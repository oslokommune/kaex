

run:
	@go run cmd/kaex/*.go

build:
	@go build -o build/kaex cmd/kaex/*.go

install:
	mkdir -p ~/.local/bin
	cp ./build/kaex ~/.local/bin/kaex

clean:
	@rm -r ./build
