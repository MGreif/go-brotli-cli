package internal

import (
	"bytes"
	"crypto/md5"
	"io"
	"os"
	"path"
	"testing"
)

func CompareFiles(t *testing.T, a *os.File, b *os.File) bool {
	t.Helper()

	bytesA, err := io.ReadAll(a)
	if err != nil {
		t.Errorf("Could not read A %s\n", err)
		return false
	}

	md5A := md5.Sum(bytesA)

	bytesB, err := io.ReadAll(b)
	if err != nil {
		t.Errorf("Could not read B %s\n", err)
		return false

	}

	md5B := md5.Sum(bytesB)

	sameHash := bytes.Equal(md5A[:], md5B[:])
	return sameHash
}

func CreateTestFile(t *testing.T) (*os.File, func()) {
	t.Helper()
	fileOut, err := os.CreateTemp(os.TempDir(), "")
	if err != nil {
		t.Errorf("Could not create temp file... %s\n", err)
	}

	return fileOut, func() {
		fileOut.Close()
		os.Remove(fileOut.Name())
	}

}

func TestCompress(t *testing.T) {
	var rawFilePath = path.Join(".", "testdata", "testfile.html")
	var compressedFilePath = path.Join(".", "testdata", "testfile.br")

	fileIn, err := os.Open(rawFilePath)
	defer fileIn.Close()
	if err != nil {
		t.Errorf("Could not open file... %s\n", err)
	}

	fileOut, cleanup := CreateTestFile(t)
	defer cleanup()

	err = HandleCompress(fileIn, fileOut, &CompressionConfig{
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

	compareCompressedFile, err := os.Open(compressedFilePath)
	if err != nil {
		t.Errorf("Could not open compare file... %s\n", err)
	}

	if !CompareFiles(t, compareCompressedFile, fileOut) {
		t.Errorf("File are not equal\n")
	}

}
