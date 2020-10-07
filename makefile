

run:
	@go run cmd/kaex/*.go

./build:
	@go build -o build/kaex cmd/kaex/*.go

install:
	cd cmd/kaex && go install

clean:
	@rm -r ./build
