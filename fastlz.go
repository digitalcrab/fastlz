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
	length := len(input)
	if length == 0 {
		return nil, errors.New("fastlz: empty input")
	}

	// Output buffer
	output := make([]byte, int(float64(length)*1.4))

	// Run
	num, err := C.fastlz_compress(unsafe.Pointer(&input[0]), C.int(length), unsafe.Pointer(&output[0]))
	if err != nil {
		return nil, fmt.Errorf("fastlz: %v", err)
	}

	// Empty compression result
	if num == 0 {
		return nil, errors.New("fastlz: compression error, empty result")
	}

	return output[:num], nil
}

// Decompress applies fastlz_decompress function to input
func Decompress(input []byte, maxOut uint) ([]byte, error) {
	length := len(input)
	if length == 0 {
		return nil, errors.New("fastlz: empty input")
	}

	// Output buffer
	output := make([]byte, int(float64(length)*10))

	// Run
	num, err := C.fastlz_decompress(unsafe.Pointer(&input[0]), C.int(length), unsafe.Pointer(&output[0]), C.int(maxOut))
	if err != nil {
		return nil, fmt.Errorf("fastlz: %v", err)
	}

	// Empty decompression result
	if num == 0 {
		return nil, errors.New("fastlz: decompression error, empty result")
	}

	return output[:num], nil
}
