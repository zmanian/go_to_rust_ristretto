ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

all: library build

clean:
	rm -rf ./lib/hello_ristretto/target
	rm -f ./lib/hello_ristretto/Cargo.lock ./lib/libhello_ristretto.dylib go-rust-ristretto

library:
	$(MAKE) -C lib/hello_ristretto build

build:
	cp lib/hello_ristretto/target/release/libhello_ristretto.dylib ./lib
	go build -ldflags="-r $(ROOT_DIR)lib" -o go-rust-ristretto


run: build
	./go-rust-ristretto


