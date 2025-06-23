package internal

import (
	"os"
	"path"
	"testing"
)

func TestDecompress(t *testing.T) {
	var rawFilePath = path.Join(".", "testdata", "testfile.html")
	var compressedFilePath = path.Join(".", "testdata", "testfile.br")

	fileIn, err := os.Open(compressedFilePath)
	defer fileIn.Close()
	if err != nil {
		t.Errorf("Could not open file... %s\n", err)
	}

	fileOut, cleanup := CreateTestFile(t)
	defer cleanup()

	err = HandleDecompress(fileIn, fileOut, &DecompressionConfig{
		DontTrimZeros: false,
		BufferSize:    4096,
		FlushInterval: 10,
	})
	if err != nil {
		t.Errorf("Could not compress file... %s\n", err)
	}

	fileOut, err = os.Open(fileOut.Name())
	if err != nil {
		t.Errorf("Could not open file... %s\n", err)
	}

	compareDecompressedFile, err := os.Open(rawFilePath)
	if err != nil {
		t.Errorf("Could not open compare file... %s\n", err)
	}

	if !CompareFiles(t, compareDecompressedFile, fileOut) {
		t.Errorf("File are not equal\n")
	}

}
