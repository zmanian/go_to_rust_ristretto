package go_to_rust_ristretto

/*
#cgo darwin LDFLAGS: -L./lib -lhello_ristretto
#include "./lib/hello_ristretto.h"
*/
import "C"

func main() {
	GenerateRistrettoPoint()
}
