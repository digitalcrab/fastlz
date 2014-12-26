package fastlz

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCompress(t *testing.T) {
	for i := 0; i <= 9; i++ {
		// Read basic data
		input, inputErr := ioutil.ReadFile(fmt.Sprintf("test_data/fastlz.%d.txt", i))
		if inputErr != nil {
			t.Error(inputErr)
		}

		output, outputErr := ioutil.ReadFile(fmt.Sprintf("test_data/fastlz.%d.mem.bin", i))
		if outputErr != nil {
			t.Fatal(outputErr)
		}

		// Read size of data (this part comes from memcached)
		r := bytes.NewReader(output)
		var size uint32
		binary.Read(r, binary.LittleEndian, &size)

		// New output value
		var readErr error
		if output, readErr = ioutil.ReadAll(r); readErr != nil {
			t.Fatal(readErr)
		}

		// Compress
		outputCompressed, errComp := Compress(input)
		if errComp != nil {
			t.Fatal(errComp)
		}

		if !bytes.Equal(output, outputCompressed) {
			t.Logf("expected: %q", output)
			t.Logf("got: %q", outputCompressed)
			t.Fatalf("result #%d: compression result is not the same as expected\n\n", i)
		}

		// Decompress
		outputDecompressed, errDecomp := Decompress(outputCompressed, uint(size))
		if errDecomp != nil {
			t.Fatal(errDecomp)
		}

		if !bytes.Equal(outputDecompressed, input) {
			t.Logf("expected: %q", input)
			t.Logf("got: %q", outputDecompressed)
			t.Fatalf("result #%d: decompression result is not the same as expected\n\n", i)
		}
	}
}
