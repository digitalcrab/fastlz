package fastlz

// #include <stdlib.h>
// #include "fastlz.h"
import "C"

// Other imports
import (
	"errors"
	"fmt"
	"unsafe"
)

// Compress applies fastlz_compress function to input
func Compress(input []byte) ([]byte, error) {
	len := len(input)
	if len == 0 {
		return nil, errors.New("fastlz: empty input")
	}

	// Output buffer
	cs := C.CString("")
	defer C.free(unsafe.Pointer(cs))

	// Run
	num, err := C.fastlz_compress(unsafe.Pointer(&input[0]), C.int(len), unsafe.Pointer(cs))
	if err != nil {
		return nil, fmt.Errorf("fastlz: %v", err)
	}

	// Empty compression result
	if num == 0 {
		return nil, errors.New("fastlz: compression error, empty result")
	}

	return C.GoBytes(unsafe.Pointer(cs), num), nil
}

// Decompress applies fastlz_decompress function to input
func Decompress(input []byte, maxOut uint) ([]byte, error) {
	len := len(input)
	if len == 0 {
		return nil, errors.New("fastlz: empty input")
	}

	// Output buffer
	cs := C.CString("")
	defer C.free(unsafe.Pointer(cs))

	// Run
	num, err := C.fastlz_decompress(unsafe.Pointer(&input[0]), C.int(len), unsafe.Pointer(cs), C.int(maxOut))
	if err != nil {
		return nil, fmt.Errorf("fastlz: %v", err)
	}

	// Empty decompression result
	if num == 0 {
		return nil, errors.New("fastlz: decompression error, empty result")
	}

	return C.GoBytes(unsafe.Pointer(cs), num), nil
}
