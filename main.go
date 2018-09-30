package main

/*
#cgo darwin LDFLAGS: -L./lib -lhello_ristretto
#include "./lib/hello_ristretto.h"
*/
import "C"

func main() {
	GenerateRistrettoPoint()
}
