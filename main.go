package main

/*
#cgo darwin LDFLAGS: -L./lib -lhello_ristretto
#include "./lib/hello_ristretto.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	generateRistrettoPoint()
}

func generateRistrettoPoint() {
	buf := make([]byte, 32, 32)
	ptr := (*C.uchar)(unsafe.Pointer(&buf[0]))
	len := C.size_t(len(buf))
	C.generate_ristretto_random(ptr, len)
	fmt.Printf("%v", buf)
}

func GenerateBulletProofs(values []int64, randomness [][32]byte) {

	valuesLen := C.size_t(len(values))
	valuePtr := (*C.ulonglong)(unsafe.Pointer(&values[0]))

	blindingPtr := (*C.uchar)(unsafe.Pointer(&randomness[0]))
	blindingLen := C.size_t(len(randomness))

	proofBuf := make([]byte, 1000, 1000)
	proofBufPtr := (*C.uchar)(unsafe.Pointer(&proofBuf[0]))
	proofBufLen := C.size_t(len(proofBuf))

	valueCommitmentsBuf := make([][32]byte, valuesLen)
	valueCommitPtr := (*C.uchar)(unsafe.Pointer(&valueCommitmentsBuf[0]))
	C.generate_ristretto_range_proof(
		valuePtr,
		valuesLen,
		blindingPtr,
		blindingLen,
		proofBufPtr,
		proofBufLen,
		valueCommitPtr,
		valuesLen,
	)

}
