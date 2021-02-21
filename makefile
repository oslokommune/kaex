BUILD_DIR ?=./build
INSTALL_DIR ?= ~/.local/bin

run:
	@go run cmd/kaex/*.go

test:
	@go test -v ./...

test-update:
	@go test -update ./...

${BUILD_DIR}/kaex:
	mkdir -p build/
	go build -o build/kaex cmd/kaex/*.go

install: ${BUILD_DIR}/kaex
	mkdir -p ${INSTALL_DIR}
	cp ${BUILD_DIR}/kaex ${INSTALL_DIR}/kaex

uninstall:
	rm ${INSTALL_DIR}/kaex

clean:
	@rm -r ./build
